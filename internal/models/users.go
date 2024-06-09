package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID uint
	Name string
	LastName string
	Username string
	HashedPassword string
	Status bool
	Created time.Time
	LastOnlineTime time.Time
	
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, lastName, username, password string) error {
	return nil
}

func (m *UserModel) Authenticate(username, password string) (bool, error) {
	return false, nil
}

func (m *UserModel) Exists(id string) (bool, error) {
	return false, nil
}