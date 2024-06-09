package server

import (
	"log"
	"os"
	"task-manager/cmd/server/service"
	"task-manager/internal/models"

	"github.com/fatih/color"
	_ "github.com/jackc/pgx/v5/stdlib"
)



type Applicaton struct {
	*config
	TaskService *service.TaskService
	InfoLog *log.Logger
	ErrorLog *log.Logger
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
	app.TaskService = &service.TaskService{
		Model: taskModel,
	}

	app.InfoLog = getLogger(os.Stdout, "INFO", color.FgGreen, log.Ldate|log.Ltime)
	app.ErrorLog = getLogger(os.Stderr, "ERROR", color.FgRed, log.Ldate|log.Ltime|log.Lshortfile)

	return &app, nil
}

func (app *Applicaton) Run() error{
	
	err := app.RunServer()
	return err
}