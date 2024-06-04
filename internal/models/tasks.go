package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID string
	Title string
	Description string
	Status bool
	Created time.Time
	Updated time.Time
	UserID string
}

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Insert(title, description, UserID string) error {
	return nil
}

func (m *TaskModel) Get(id string) (*Task, error) {
	return nil, nil
}

func (m *TaskModel) GetLists(id []string) ([]*Task, error) {
	return nil, nil
}

func (m *TaskModel) Update(id string) error {
	return nil
}

