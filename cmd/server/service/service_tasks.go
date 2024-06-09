package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"task-manager/internal/models"
)

type TaskService struct {
	Model *models.TaskModel
}

func (s *TaskService) Get(id uint) (*bytes.Buffer, error) {
	task, err := s.Model.Get(id)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(task); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TaskService) GetTasks() (*bytes.Buffer, error) {
	tasks, err := s.Model.GetLists()
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(tasks); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TaskService) Insert(data io.Reader) (*bytes.Buffer, error) {
	model := &models.Task{}
	if err := JsonDecode(data, model); err != nil {
		return nil, err
	}

	task, err := s.Model.Insert(model.Title, model.Description, model.UserID)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(task); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TaskService) Update(id uint, data io.Reader) (*bytes.Buffer, error) {
	model, err := s.Model.Get(id)
	if err != nil {
		return nil, err
	}
	
	if err := JsonDecode(data, model); err != nil {
		return nil, err
	}
	task, err := s.Model.Update(id, model.Title, model.Description, model.Status)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(task); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TaskService) Delete(id uint) error {
	err := s.Model.Delete(id)
	
	return err
}