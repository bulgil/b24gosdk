package tasks

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/bulgil/b24gosdk/transport"
)

var ErrNoTasks = errors.New("tasks not found")

type methods struct {
	add, get, update, delete, list string
}

type method string

var (
	methodAdd          method = "tasks.task.add"
	methodUpdate       method = "tasks.task.update"
	methodGet          method = "tasks.task.get"
	methodList         method = "tasks.task.list"
	methodCommentAdd   method = "task.commentitem.add"
	methodTaskComplete method = "tasks.task.complete"
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

func (s *TaskService) Create(fields map[string]any) (int64, error) {
	const op = "TaskService.Create"

	wh := path.Join(s.webhook, string(methodAdd))

	body := struct {
		Fields map[string]any `json:"fields"`
	}{
		Fields: fields,
	}

	var result struct {
		Task struct {
			ID string `json:"id"`
		} `json:"task"`
	}
	if err := s.transport.Call(http.MethodPost, wh, nil, body, &result); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	taskID, err := strconv.Atoi(result.Task.ID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return int64(taskID), nil
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
		Filter map[string]any    `json:"filter"`
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

	if len(result.Tasks) == 0 {
		return nil, ErrNoTasks
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

func (s *TaskService) Update(taskID int64, fields any) error {
	const op = "TaskService.Update"

	wh := path.Join(s.webhook, string(methodUpdate))

	query := url.Values{
		"taskId": []string{strconv.Itoa(int(taskID))},
	}

	var body = struct {
		Fields any `json:"fields"`
	}{
		Fields: fields,
	}

	if err := s.transport.Call(http.MethodPost, wh, query, body, nil); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TaskService) Complete(taskID int64) error {
	const op = "TaskService.Complete"

	wh := path.Join(s.webhook, string(methodTaskComplete))

	var body = struct {
		TaskID int64 `json:"taskId"`
	}{
		TaskID: taskID,
	}

	if err := s.transport.Call(http.MethodPost, wh, nil, body, nil); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
