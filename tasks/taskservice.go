package tasks

import (
	"fmt"
	"net/http"
	"path"
)

type methods struct {
	add, get, update, delete, list string
}

type method string

var (
	methodAdd    method = "tasks.task.add"
	methodUpdate method = "tasks.task.update"
	methodGet    method = "tasks.task.get"
	methodList   method = "tasks.task.list"
)

type TaskService struct {
	transport Transport
	webhook   string
}

func NewTaskService(transport Transport, webhook string) *TaskService {
	return &TaskService{
		transport: transport,
		webhook:   webhook,
	}
}

func (s *TaskService) Get(id int, sel []string) (Task, error) {
	const op = "TaskService.Get"

	sel = append(sel, "UF_CRM_TASK")

	wh := path.Join(s.webhook, string(methodGet))
	var body = struct {
		TaskID int      `json:"taskId"`
		Select []string `json:"select"`
	}{
		TaskID: id,
		Select: sel,
	}

	var result struct {
		Task Task `json:"task"`
	}
	if err := s.transport.Call(http.MethodPost, wh, nil, body, &result); err != nil {
		return Task{}, fmt.Errorf("%s: %w", op, err)
	}

	return result.Task, nil
}
