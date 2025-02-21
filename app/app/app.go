package app

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

)

type App struct {

	router      *gin.Engine
	modules     []module.Module
}

func NewApp(cfg *config.Config) *App {


	app := &App{

		router:      gin.Default(),
	}

	app.setupModules()

	return app
}

func (a *App) Router() *gin.Engine {
	return a.router
}

func (a *App) Start() error {
	return a.router.Run(":" + a.config.Application.Ports[0])
}

func (a *App) Close() error {
	return a.firestoreDB.Close()
}
