package app

import (

	// "github.com/ynwd/awesome-blog/internal/comments"
	// "github.com/ynwd/awesome-blog/internal/likes"
	// "github.com/ynwd/awesome-blog/internal/posts"
	// "github.com/ynwd/awesome-blog/internal/summary"
	// "github.com/ynwd/awesome-blog/internal/users"
	// "github.com/ynwd/awesome-blog/pkg/middleware"
	// "github.com/ynwd/awesome-blog/pkg/module"
	// "github.com/ynwd/awesome-blog/pkg/utils"

	"flag"
	"fmt"
	database "stncCms/pkg/database"

	cms "stncCms/app/post/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func (a *App) setupModules() {

	db := database.DbConnect()
	services, err := database.RepositoriesInit(db)
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

	// indexHandle := sacrifice.InitDashboard(services.Dashboard)
	posts := cms.InitPost(services.Post, services.Cat, services.CatPost, services.Lang, services.User)

	a.router.Use(gin.Recovery())
	a.router.Use(gin.Logger())

	store := cookie.NewStore([]byte("rodrigoHunter"))
	////60 dakika olan 1 saat tam olarak ( 60x60) 3600 saniyedir.
	//60 saniye * 60 = 1 saat //60*60
	//3600 (1 saat ) * 5 = 5 saat
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 3600 * 8}) //Also set Secure: true if using SSL, you should though
	// store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: -1}) //Also set Secure: true if using SSL, you should though

	a.router.Use(sessions.Sessions("myCRM", store))

	//TODO: csrf ???
	a.router.Use(csrf.Middleware(csrf.Options{
		Secret: "rodrigoHunter",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))



	modules := []module.Module{
		// users.NewModule(client),
		posts.NewModule(client, a.pubsub),
		// comments.NewModule(client, a.pubsub),
		// likes.NewModule(client, a.pubsub),
		// summary.NewModule(client),
	}

	for _, m := range modules {
		m.RegisterRoutes(a.router)
	}

	a.modules = modules
}
