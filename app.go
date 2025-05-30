package main

import (
	"context"
	"log"
	"path/filepath"
	"tsdoa/internal/db"
	"tsdoa/internal/services"
	"tsdoa/internal/utils"
)

// App struct
type App struct {
	ctx         context.Context
	taskService *services.TaskService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		taskService: services.NewTaskService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.taskService.SetContext(ctx)

	// Resolve cross platform DB paths
	path, err := utils.GetAppDataPath()
	if err != nil {
		log.Fatal("Failed to get app data path:", err)
	}

	dbPath := filepath.Join(path, "db")

	// Initialize badger DB
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal("Failed to initialise database:", err)
	}

	log.Println("📦 BadgerDB initialized at", dbPath)
}
