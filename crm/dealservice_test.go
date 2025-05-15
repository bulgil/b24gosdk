package crm

import (
	"reflect"
	"testing"

	"github.com/bulgil/b24gosdk"
)

func newMockDealService(client client) *DealService {
	return &DealService{
		CRMService: CRMService[Deal]{
			client:  client,
			webhook: "/webhook",
			methods: methods{
				add:    "crm.deal.add",
				get:    "crm.deal.get",
				update: "crm.deal.update",
				delete: "crm.deal.delete",
				list:   "crm.deal.list",
			},
		},
	}
}

func TestDealService_Add_Success(t *testing.T) {
	mockClient := &MockClient[b24gosdk.B24int]{
		RespponseToReturn: ptr(b24gosdk.B24int(123)),
	}

	service := newMockDealService(mockClient)

	fields := map[string]any{"TITLE": "New Deal"}
	id, err := service.Add(fields, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id != 123 {
		t.Errorf("expected id 123, got %d", id)
	}

	if mockClient.CalledMethod != "POST" {
		t.Errorf("expected POST method, got %s", mockClient.CalledMethod)
	}

	if mockClient.CalledURL != "/webhook/crm.deal.add" {
		t.Errorf("unexpected URL: %s", mockClient.CalledURL)
	}
}

func TestDealService_Add_ErrorOnNilFields(t *testing.T) {
	service := newMockDealService(nil)

	_, err := service.Add(nil, nil)
	if err == nil {
		t.Fatal("expected error for nil fields, got nil")
	}
}

func TestDealService_Get_Success(t *testing.T) {
	expectedDeal := Deal{
		ID:    ptr(b24gosdk.B24int(42)),
		Title: ptr("Test Deal"),
	}

	mockClient := &MockClient[Deal]{
		RespponseToReturn: ptr(expectedDeal),
	}

	service := newMockDealService(mockClient)

	deal, err := service.Get(42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if deal.ID != ptr(b24gosdk.B24int(42)) || deal.Title != ptr("Test Deal") {
		t.Errorf("unexpected deal returned: %+v", deal)
	}

	if mockClient.CalledMethod != "GET" {
		t.Errorf("expected GET method, got %s", mockClient.CalledMethod)
	}

	if mockClient.CalledQuery.Get("id") != "42" {
		t.Errorf("expected query id=42, got %v", mockClient.CalledQuery)
	}
}

func TestDealService_Update_Success(t *testing.T) {
	mockClient := &MockClient[bool]{
		RespponseToReturn: ptr(true),
	}

	service := newMockDealService(mockClient)

	fields := map[string]any{"TITLE": "Updated Deal"}
	result, err := service.Update(42, fields, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result {
		t.Error("expected true result")
	}

	if mockClient.CalledMethod != "POST" {
		t.Errorf("expected POST method, got %s", mockClient.CalledMethod)
	}
}

func TestDealService_Update_ErrorOnNilFields(t *testing.T) {
	service := newMockDealService(nil)

	_, err := service.Update(42, nil, nil)
	if err == nil {
		t.Fatal("expected error for nil fields, got nil")
	}
}

func TestDealService_Delete_Success(t *testing.T) {
	mockClient := &MockClient[bool]{
		RespponseToReturn: ptr(true),
	}

	service := newMockDealService(mockClient)

	result, err := service.Delete(42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !result {
		t.Error("expected true result")
	}

	if mockClient.CalledMethod != "GET" {
		t.Errorf("expected GET method, got %s", mockClient.CalledMethod)
	}

	if mockClient.CalledQuery.Get("id") != "42" {
		t.Errorf("expected query id=42, got %v", mockClient.CalledQuery)
	}
}

func TestDealService_List_Success(t *testing.T) {
	expectedDeals := []*Deal{
		{ID: ptr(b24gosdk.B24int(1)), Title: ptr("Deal 1")},
		{ID: ptr(b24gosdk.B24int(2)), Title: ptr("Deal 2")},
	}

	mockClient := &MockClient[[]*Deal]{
		RespponseToReturn: ptr(expectedDeals),
	}

	service := newMockDealService(mockClient)

	sel := []string{"ID", "TITLE"}
	filter := map[string]any{"STATUS": "OPEN"}
	order := map[string]string{"ID": "DESC"}
	start := 0

	deals, err := service.List(sel, filter, order, start)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(deals, expectedDeals) {
		t.Errorf("expected %v, got %v", expectedDeals, deals)
	}

	if mockClient.CalledMethod != "POST" {
		t.Errorf("expected POST method, got %s", mockClient.CalledMethod)
	}
}
