package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
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
	if err != nil  && errors.Is(err, sql.ErrNoRows){
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(tasks); err != nil {
		return nil, err
	}

	return buf, nil
}

func (s *TaskService) Insert( title, description string, userID uint) (*bytes.Buffer, error) {
	task, err := s.Model.Insert(title, description, userID)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(task); err != nil {
		return nil, err
	}
	
	return buf, nil
}

func (s *TaskService) Update(id uint, title, descritpion string, status bool) (*bytes.Buffer, error) {
	task, err := s.Model.Update(id, title, descritpion, status)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}

	if err := json.NewEncoder(buf).Encode(task); err != nil {
		return nil, err
	}
	
	return buf, nil
}