package crm

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/bulgil/b24gosdk"
)

type MockClient[T any] struct {
	CalledMethod string
	CalledURL    string
	CalledQuery  url.Values
	CalledBody   any

	RespponseToReturn *T
	ErrorToReturn     error
}

func (m *MockClient[T]) Call(method, url string, query url.Values, body, resp any) error {
	m.CalledMethod = method
	m.CalledURL = url
	m.CalledQuery = query
	m.CalledBody = body

	if m.RespponseToReturn != nil {
		if outPtr, ok := resp.(*T); ok {
			*outPtr = *m.RespponseToReturn
		}
	}

	return m.ErrorToReturn
}

type MockListClient[T any] struct {
	Called   bool
	Response []*T
	Err      error
}

func (m *MockListClient[T]) Call(method, url string, query url.Values, body, resp any) error {
	m.Called = true
	if r, ok := resp.(*[]*T); ok {
		*r = m.Response
	}
	return m.Err
}

func Test_CRMServiceGet_Success(t *testing.T) {
	type entity struct {
		ID   int
		Name string
	}

	expected := &entity{ID: 42, Name: "Test"}
	mockClient := &MockClient[entity]{
		RespponseToReturn: expected,
	}

	service := CRMService[entity]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			get: "crm.get",
		},
	}

	result, err := service.Get(42)
	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}

	if result == nil || result.ID != 42 || result.Name != "Test" {
		t.Fatalf("expected result %+v, got %+v", expected, result)
	}

	if mockClient.CalledMethod != http.MethodGet {
		t.Errorf("expected GET method, got %v", mockClient.CalledMethod)
	}

	if mockClient.CalledURL != "/webhook/crm.get" {
		t.Errorf("expected webhook %v, got %v", "/webhook/crm.get", mockClient.CalledURL)
	}

	if mockClient.CalledQuery.Get("id") != "42" {
		t.Errorf("expected query id=42, got %+v", mockClient.CalledQuery.Get("id"))
	}
}

func Test_CRMServiceGet_Error(t *testing.T) {
	mockClient := &MockClient[struct{}]{
		ErrorToReturn: fmt.Errorf("network error"),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			get: "crm.get",
		},
	}

	result, err := service.Get(42)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "network error") {
		t.Errorf("expected error to contain 'network error', got %+v", err)
	}

	if result != nil {
		t.Errorf("expected result to be nil, got %+v", result)
	}
}

func Test_CRMServiceUpdate_Success(t *testing.T) {
	mockClient := &MockClient[bool]{
		RespponseToReturn: ptr[bool](true),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			update: "crm.update",
		},
	}

	result, err := service.Update(42, map[string]any{
		"TITLE":          "name",
		"ASSIGNED_BY_ID": 42,
	}, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !result {
		t.Errorf("expected result true, got %v", result)
	}

	if mockClient.CalledMethod != http.MethodPost {
		t.Errorf("expected method POST, got %v", mockClient.CalledMethod)
	}

	if mockClient.CalledQuery != nil {
		t.Errorf("expected query nil, got %+v", mockClient.CalledQuery)
	}

	if mockClient.CalledURL != "/webhook/crm.update" {
		t.Errorf("expected URL /webhook/crm.update, got %v", mockClient.CalledURL)
	}

	expectedBody := struct {
		ID     int `json:"id"`
		Fields any `json:"fields"`
		Params any `json:"params,omitempty"`
	}{
		ID:     42,
		Fields: map[string]any{"TITLE": "name", "ASSIGNED_BY_ID": 42},
		Params: nil,
	}
	if !reflect.DeepEqual(mockClient.CalledBody, expectedBody) {
		t.Errorf("expected body %+v, got %+v", expectedBody, mockClient.CalledBody)
	}
}

func Test_CRMServiceUpdate_Error(t *testing.T) {
	mockClient := &MockClient[bool]{
		ErrorToReturn: fmt.Errorf("update failed"),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			update: "crm.update",
		},
	}

	result, err := service.Update(42, map[string]any{
		"TITLE":          "name",
		"ASSIGNED_BY_ID": 42,
	}, nil)

	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "update failed") {
		t.Errorf("expected error to contain 'update failed', got %v", err)
	}

	if result {
		t.Errorf("expected result to be false, got true")
	}

	if mockClient.CalledMethod != http.MethodPost {
		t.Errorf("expected method POST, got %v", mockClient.CalledMethod)
	}

	if mockClient.CalledURL != "/webhook/crm.update" {
		t.Errorf("expected URL /webhook/crm.update, got %v", mockClient.CalledURL)
	}

	expectedBody := struct {
		ID     int `json:"id"`
		Fields any `json:"fields"`
		Params any `json:"params,omitempty"`
	}{
		ID:     42,
		Fields: map[string]any{"TITLE": "name", "ASSIGNED_BY_ID": 42},
		Params: nil,
	}

	if !reflect.DeepEqual(mockClient.CalledBody, expectedBody) {
		t.Errorf("expected body %+v, got %+v", expectedBody, mockClient.CalledBody)
	}
}

func Test_CRMServiceDelete_Success(t *testing.T) {
	mockClient := &MockClient[bool]{
		RespponseToReturn: ptr(true),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			delete: "crm.delete",
		},
	}

	result, err := service.Delete(42)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !result {
		t.Errorf("expected result true, got %v", result)
	}

	if mockClient.CalledMethod != http.MethodGet {
		t.Errorf("expected method GET, got %v", mockClient.CalledMethod)
	}

	if mockClient.CalledURL != "/webhook/crm.delete" {
		t.Errorf("expected URL /webhook/crm.delete, got %v", mockClient.CalledURL)
	}

	if mockClient.CalledQuery.Get("id") != "42" {
		t.Errorf("expected query id=42, got %v", mockClient.CalledQuery.Get("id"))
	}
}

func Test_CRMServiceDelete_Error(t *testing.T) {
	mockClient := &MockClient[bool]{
		ErrorToReturn: fmt.Errorf("delete failed"),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			delete: "crm.delete",
		},
	}

	result, err := service.Delete(42)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "delete failed") {
		t.Errorf("expected error to contain 'delete failed', got %v", err)
	}

	if result {
		t.Errorf("expected result to be false, got true")
	}

	if mockClient.CalledMethod != http.MethodGet {
		t.Errorf("expected method GET, got %v", mockClient.CalledMethod)
	}

	if mockClient.CalledURL != "/webhook/crm.delete" {
		t.Errorf("expected URL /webhook/crm.delete, got %v", mockClient.CalledURL)
	}

	if mockClient.CalledQuery.Get("id") != "42" {
		t.Errorf("expected query id=42, got %v", mockClient.CalledQuery.Get("id"))
	}
}

func Test_CRMServiceAdd_Success(t *testing.T) {
	mockClient := &MockClient[b24gosdk.B24int]{
		RespponseToReturn: ptr(b24gosdk.B24int(123)),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			add: "crm.add",
		},
	}

	fields := map[string]any{
		"TITLE": "New Deal",
	}
	params := map[string]any{
		"REGISTER_SONET_EVENT": "Y",
	}

	result, err := service.Add(fields, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != 123 {
		t.Errorf("expected ID 123, got %d", result)
	}

	if mockClient.CalledMethod != http.MethodPost {
		t.Errorf("expected POST, got %s", mockClient.CalledMethod)
	}
	if mockClient.CalledURL != "/webhook/crm.add" {
		t.Errorf("expected URL /webhook/crm.add, got %v", mockClient.CalledURL)
	}
}

func Test_CRMServiceAdd_ErrorFromClient(t *testing.T) {
	mockClient := &MockClient[b24gosdk.B24int]{
		ErrorToReturn: fmt.Errorf("call failed"),
	}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			add: "crm.add",
		},
	}

	fields := map[string]any{"TITLE": "New Deal"}

	result, err := service.Add(fields, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "call failed") {
		t.Errorf("expected error to contain 'call failed', got %v", err)
	}
	if result != 0 {
		t.Errorf("expected result 0, got %d", result)
	}
}

func Test_CRMServiceAdd_NilFields(t *testing.T) {
	mockClient := &MockClient[b24gosdk.B24int]{}

	service := CRMService[struct{}]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{
			add: "crm.add",
		},
	}

	result, err := service.Add(nil, nil)
	if err == nil {
		t.Fatal("expected error for nil fields, got nil")
	}
	if !errors.Is(err, ErrGivenNoFields) {
		t.Errorf("expected ErrGivenNoFields, got %v", err)
	}
	if result != 0 {
		t.Errorf("expected result 0, got %d", result)
	}
}

func Test_CRMServiceList_Success(t *testing.T) {
	type entity struct {
		ID   int
		Name string
	}

	expected := []*entity{
		{ID: 1, Name: "Test"},
		{ID: 2, Name: "More"},
	}

	mockClient := &MockListClient[entity]{
		Response: expected,
		Err:      nil,
	}

	service := CRMService[entity]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{list: "crm.list"},
	}

	res, err := service.List([]string{"ID", "NAME"}, nil, nil, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}
}

func Test_CRMServiceList_ErrorFromClient(t *testing.T) {
	type entity struct {
		ID   int
		Name string
	}

	mockClient := &MockClient[[]*entity]{
		ErrorToReturn: fmt.Errorf("network error"),
	}

	service := CRMService[entity]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{list: "crm.list"},
	}

	res, err := service.List([]string{"ID", "NAME"}, nil, nil, 0)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if !strings.Contains(err.Error(), "network error") {
		t.Errorf("expected error to contain 'network error', got: %v", err)
	}

	if res != nil {
		t.Errorf("expected result to be nil, got: %+v", res)
	}
}

func Test_CRMServiceList_EmptyResponse(t *testing.T) {
	type entity struct {
		ID   int
		Name string
	}

	mockClient := &MockClient[[]*entity]{
		RespponseToReturn: &[]*entity{},
	}

	service := CRMService[entity]{
		client:  mockClient,
		webhook: "/webhook",
		methods: methods{list: "crm.list"},
	}

	res, err := service.List([]string{"ID"}, nil, nil, 0)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if res == nil {
		t.Fatal("expected non-nil result (even if empty slice), got nil")
	}

	if len(res) != 0 {
		t.Errorf("expected empty result, got: %+v", res)
	}
}

func ptr[T any](v T) *T {
	return &v
}
