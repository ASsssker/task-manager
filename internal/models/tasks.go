package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID uint
	Title string
	Description string
	Status bool
	Created time.Time
	Updated time.Time
	UserID uint
}

type TaskModel struct {
	DB *sql.DB
}

func GetTaskModel(db *sql.DB) (*TaskModel, error) {
	model := &TaskModel{DB: db}

	if err := model.initTable(); err != nil {
		return nil, err
	}

	return model, nil
}

func (m *TaskModel) initTable() error {
	stmt := "CREATE TABLE IF NOT EXISTS tasks " + 
	"(id SERIAL PRIMARY KEY, title TEXT, description TEXT,"  +
	" status BOOLEAN, created TIMESTAMP, updated TIMESTAMP, user_id INTEGER)"

	if _, err := m.DB.Exec(stmt); err != nil {
		return err
	}

	return nil
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

