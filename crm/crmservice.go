package crm

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/bulgil/b24gosdk"
)

var ErrGivenNoFields = errors.New("no given fields to update")

type methods struct {
	add, get, update, delete, list string
}

type client interface {
	Call(method, webhook string, query url.Values, body, result any) error
}

type CRMService[T any] struct {
	client  client
	webhook string
	methods methods
}

func NewCrmService[T any](webhook string, methods methods) CRMService[T] {
	const op = "NewCrmService"

	u, err := url.Parse(webhook)
	if err != nil {
		panic(fmt.Sprintf("%s: webhook parsing error: %s", err, op))
	}

	return CRMService[T]{
		client:  b24gosdk.NewClient(nil, webhook),
		webhook: u.Path,
		methods: methods,
	}
}

func (s *CRMService[T]) Add(fields, params any) (int, error) {
	const op = "CRMService.Add"

	if fields == nil {
		return 0, fmt.Errorf("%s: %w", op, ErrGivenNoFields)
	}

	wh := path.Join(s.webhook, s.methods.add)
	var body = struct {
		Fields any `json:"fields"`
		Params any `json:"params,omitempty"`
	}{
		Fields: fields,
		Params: params,
	}

	var id b24gosdk.B24int
	err := s.client.Call(http.MethodPost, wh, nil, body, &id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int(id), nil
}

func (s *CRMService[T]) Get(id int) (*T, error) {
	const op = "CRMSerivce.Get"

	wh := path.Join(s.webhook, s.methods.get)
	query := url.Values{
		"id": {strconv.Itoa(id)},
	}

	var entity T
	err := s.client.Call(http.MethodGet, wh, query, nil, &entity)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &entity, nil
}

func (s *CRMService[T]) Update(id int, fields, params any) (bool, error) {
	const op = "CRMService.Update"

	if fields == nil {
		return false, fmt.Errorf("%s: %w", op, ErrGivenNoFields)
	}

	wh := path.Join(s.webhook, s.methods.update)
	body := struct {
		ID     int `json:"id"`
		Fields any `json:"fields"`
		Params any `json:"params,omitempty"`
	}{
		ID:     id,
		Fields: fields,
		Params: params,
	}

	var result bool
	err := s.client.Call(http.MethodPost, wh, nil, body, &result)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

func (s *CRMService[T]) Delete(id int) (bool, error) {
	const op = "CRMService.Delete"

	wh := path.Join(s.webhook, s.methods.delete)
	query := url.Values{
		"id": {strconv.Itoa(id)},
	}

	var result bool
	err := s.client.Call(http.MethodGet, wh, query, nil, &result)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil
}

func (s *CRMService[T]) List(sel []string, filter, order any, start int) ([]*T, error) {
	const op = "CRMService.List"

	wh := path.Join(s.webhook, s.methods.list)

	body := struct {
		Select []string `json:"select,omitempty"`
		Filter any      `json:"filter,omitempty"`
		Order  any      `json:"order,omitempty"`
		Start  int      `json:"start,omitempty"`
	}{
		Select: sel,
		Filter: filter,
		Order:  order,
		Start:  start,
	}

	var entities []*T
	err := s.client.Call(http.MethodPost, wh, nil, body, &entities)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return entities, nil
}
