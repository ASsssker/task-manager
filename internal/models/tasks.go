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

func (m *TaskModel) Insert(title, description string, userID uint) (*Task, error) {
	stmt := "INSERT INTO tasks (title, description, status, created, updated, user_id) " +
			"VALUES ($1, $2, $3, $4, $5, $6) " +
			"RETURNING id"
	
	t := &Task{
		Title: title,
		Description: description,
		Status: false,
		Created: time.Now(),
		Updated: time.Now(),
		UserID: userID,
	}

	if err := m.DB.QueryRow(stmt, t.Title, t.Description, t.Status, t.Created, t.Updated, t.UserID).Scan(&t.ID); err != nil {
		return nil, err
	}

	return t, nil

}

func (m *TaskModel) Get(id uint) (*Task, error) {
	stmt := "SELECT id, title, description, status, created, updated, user_id FROM tasks WHERE id = $1"

	t := &Task{}
	if err := m.DB.QueryRow(stmt, id).Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Created, &t.Updated, &t.UserID); err != nil {
		return nil, err
	}

	return t, nil

}

func (m *TaskModel) GetLists() ([]*Task, error) {
	stmt := "SELECT id, title, description, status, created, updated, user_id FROM tasks"

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*Task{}

	for rows.Next(){
		t := &Task{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.Created, &t.Updated, &t.UserID); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil

	
}

func (m *TaskModel) Update(id uint, title, description string, status bool) (*Task, error) {
	stmt := "UPDATE tasks SET (title = $1, description = $2, updated = $3) WHERE id = $4 " +
			"RETURNING created, user_id"

	t := &Task{
		ID: id,
		Title: title,
		Description: description,
		Status: status,
		Updated: time.Now(),
	}

	if err := m.DB.QueryRow(stmt, t.Title, t.Description, t.Updated, t.ID).Scan(&t.Created, &t.UserID); err != nil {
		return nil, err
	}

	return t, nil
}

