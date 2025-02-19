package app

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	repository "stncCms/app/domain/repository/cacheRepository"
)

type App struct {
	
	router      *gin.Engine
	DB *gorm.DB
	
	modules     []module.Module
}

func NewApp() *App {



	//
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}

	services.Automigrate()

	autoRelation := flag.Bool("autoRelation", false, "db relation ")
	flag.Parse()

	if *autoRelation {
		fmt.Printf("\033[1;34m%s\033[0m", "-----------setup relations-------------")
		services.AutoRelation()
		fmt.Printf("\033[1;34m%s\033[0m", "-----------done relations-------------")
		return
	}



	app := &App{

		router:      gin.Default(),
		firestoreDB: firestoreDB,

	}

	// Setup middleware
	//app.setupMiddleware()

	// Setup modules
	//app.setupModules()


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