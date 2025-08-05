package tasks

import (
	"fmt"
	"net/http"
	"path"

	"github.com/bulgil/b24gosdk/transport"
)

type methods struct {
	add, get, update, delete, list string
}

type method string

var (
	methodAdd        method = "tasks.task.add"
	methodUpdate     method = "tasks.task.update"
	methodGet        method = "tasks.task.get"
	methodList       method = "tasks.task.list"
	methodCommentAdd method = "task.commentitem.add"
)

type TaskService struct {
	transport *transport.Transport
	webhook   string
}

func NewTaskService(transport *transport.Transport, webhook string) *TaskService {
	return &TaskService{
		transport: transport,
		webhook:   webhook,
	}
}

func (s *TaskService) Get(id int, sel []string) (Task, error) {
	const op = "TaskService.Get"

	wh := path.Join(s.webhook, string(methodGet))
	sel = append(sel, "UF_CRM_TASK")
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

func (s *TaskService) List(order map[string]string, filter map[string]any, sel []string, limit, start int) ([]Task, error) {
	const op = "TaskService.List"

	wh := path.Join(s.webhook, string(methodList))
	var body = struct {
		Order  map[string]string `json:"order"`
		Filter map[string]string `json:"filter"`
		Sel    []string          `json:"select"`
		Limit  int               `json:"limit"`
		Start  int               `json:"start"`
	}{
		Order:  order,
		Filter: filter,
		Sel:    sel,
		Limit:  limit,
		Start:  start,
	}

	var result struct {
		Tasks []Task `json:"tasks"`
	}
	if err := s.transport.Call(http.MethodPost, wh, nil, body, &result); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result.Tasks, nil
}

func (s *TaskService) AddComment(taskID int, message string, authorID int) (int, error) {
	const op = "TaskService.AddComment"

	wh := path.Join(s.webhook, string(methodCommentAdd))
	var body = struct {
		TaskID int `json:"TASK_ID"`
		Fields struct {
			PostMessage string `json:"POST_MESSAGE"`
			AuthorID    int    `json:"AUTHOR_ID"`
		} `json:"FIELDS"`
	}{
		TaskID: taskID,
		Fields: struct {
			PostMessage string `json:"POST_MESSAGE"`
			AuthorID    int    `json:"AUTHOR_ID"`
		}{
			PostMessage: message,
			AuthorID:    authorID,
		},
	}

	var commentID int
	if err := s.transport.Call(http.MethodPost, wh, nil, body, &commentID); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return commentID, nil
}
