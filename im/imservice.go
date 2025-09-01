package im

import (
	"fmt"
	"net/http"
	"path"

	"github.com/bulgil/b24gosdk"
	"github.com/bulgil/b24gosdk/transport"
)

type method string

var MethodImMessageAdd method = "im.message.add"

type IMService struct {
	transport *transport.Transport
	webhook   string
}

func NewIMService(transport *transport.Transport, webhook string) *IMService {
	return &IMService{
		transport: transport,
		webhook:   webhook,
	}
}

func (s *IMService) MessageAdd(dialogID string, message string) (int64, error) {
	const op = "IMService.MessageAdd"

	wh := path.Join(s.webhook, string(MethodImMessageAdd))

	var body = Message{
		DialogID: dialogID,
		Message:  message,
	}

	var messageID b24gosdk.B24int

	if err := s.transport.Call(http.MethodPost, wh, nil, body, &messageID); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int64(messageID), nil
}
