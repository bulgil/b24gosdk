package notify

import (
	"fmt"
	"net/http"
	"path"

	"github.com/bulgil/b24gosdk/transport"
)

type method string

var methodAdd method = "im.notify.personal.add"

type NotifyService struct {
	transport *transport.Transport
	webhook   string
}

func NewNotifyService(transport *transport.Transport, webhook string) *NotifyService {
	return &NotifyService{
		transport: transport,
		webhook:   webhook,
	}
}

func (s *NotifyService) Send(userID int64, message string) (int64, error) {
	const op = "NotifyService.Send"

	wh := path.Join(s.webhook, string(methodAdd))

	body := struct {
		UserID  int64  `json:"USER_ID"`
		Message string `json:"MESSAGE"`
	}{
		UserID:  userID,
		Message: message,
	}

	var id int64
	if err := s.transport.Call(http.MethodPost, wh, nil, body, &id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
