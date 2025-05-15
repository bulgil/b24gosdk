package crm

import "net/url"

type Client interface {
	Call(method, webhook string, query url.Values, body, result any) error
}
