package user

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/bulgil/b24gosdk/transport"
)

type method string

var (
	methodCurrent method = "user.current"
)

type UserService struct {
	transport *transport.Transport
	webhook   string
}

func NewUserService(transport *transport.Transport, webhook string) *UserService {
	return &UserService{
		transport: transport,
		webhook:   webhook,
	}
}

func (s *UserService) Current(authID string) (User, error) {
	const op = "UserService.Current"

	wh := path.Join(s.webhook, string(methodCurrent))

	var u User

	if err := s.transport.Call(http.MethodGet, wh, url.Values{
		"auth": []string{authID},
	}, nil, &u); err != nil {
		return User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}
