package crm

import "net/url"

type Transport interface {
	Call(method, webhook string, query url.Values, body, result any) error
}
