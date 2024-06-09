package server

import (
	"task-manager/internal/models"

	_"github.com/jackc/pgx/v5/stdlib"
)



type Applicaton struct {
	*config
	taskModel *models.TaskModel
	userModel *models.UserModel
}

func GetApp() (*Applicaton, error) {
	var app Applicaton
	
	conf, err := getConfig()
	if err != nil {
		return nil, err
	}
	app.config = conf

	db, err := getDB("pgx", "postgres://task_manager:qwerty123@localhost:5432/task_manager_db")
	if err != nil {
		return nil, err
	}

	taskModel, err := models.GetTaskModel(db)
	if err != nil {
		return nil, err
	}
	app.taskModel = taskModel


	return &app, nil
}

func (app *Applicaton) Run() error{
	err := app.RunServer()
	return err
}