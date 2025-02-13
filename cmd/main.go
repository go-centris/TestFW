// golang gin framework mvc and clean code project
// Licensed under the Apache License 2.0
// @author Selman TUNÇ <selmantunc@gmail.com>
// @link: https://github.com/stnc/go-mvc-blog-clean-code
// @license: Apache License 2.0
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	repository "stncCms/app/domain/repository/cacheRepository"

	cms "stncCms/app/post/controller"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leonelquinteros/gotext"
	myPongoGinRender "github.com/stnc/myPongoGinRender/v5"
	csrf "github.com/utrack/gin-csrf"

	//modSacrife "stncCms/app/web.api/controller/modSacrife"
	auth "stncCms/app/auth/controller"
	common "stncCms/app/web/controller/common_mod"
	region "stncCms/app/web/controller/region_mod"

	sacrifice "stncCms/app/web/controller/sacrifice_mod"
)

var cacheControlSelman = false

// hk@gmail.com  mutluerF9E
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

}

func main() {

	debugMode := os.Getenv("MODE")

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

	// indexHandle := sacrifice.InitDashboard(services.Dashboard)
	posts := cms.InitPost(services.Post, services.Cat, services.CatPost, services.Lang, services.User)

	loginHandle := auth.InitLogin(services.User)

	userHandle := auth.InitUserControl(services.User, services.Region, services.Role)
	roleHandle := auth.InitRoles(services.Permission, services.Modules, services.Role, services.RolePermission)

	optionsHandle := sacrifice.InitOptions(services.Options)

	branchHandle := region.InitBranch(services.Branch, services.Region)

	regionHandle := region.InitRegion(services.Region)

	modulesHandle := common.InitModules(services.Modules)

	switch debugMode {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	store := cookie.NewStore([]byte("rodrigoHunter"))
	////60 dakika olan 1 saat tam olarak ( 60x60) 3600 saniyedir.
	//60 saniye * 60 = 1 saat //60*60
	//3600 (1 saat ) * 5 = 5 saat
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 3600 * 8}) //Also set Secure: true if using SSL, you should though
	// store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: -1}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))

	//TODO: csrf ???
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "rodrigoHunter",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.HTMLRender = myPongoGinRender.TemplatePath("public/view")

	r.MaxMultipartMemory = 1 >> 20 // 8 MiB

	r.NoRoute(func(c *gin.Context) {
		var getText *gotext.Locale
		getText = gotext.NewLocale("public/locales", "tr_TR")
		getText.AddDomain("l404")
		viewData := pongo2.Context{
			"title":  "404",
			"locale": getText,
		}
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusNotFound,
			"/admin/404.html",
			viewData,
		)
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	//r.GET("/", controller.Index)
	//r.GET("admin", controller.Index)
	//r.GET("admin/", controller.Index)

	//default router --- for direct admin dashboard
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "admin/sacrifece/post/")
	})

	// r.GET("optionsDefault", sacrifice.OptionsDefault)
	// r.GET("cacheReset", sacrifice.CacheReset)

	r.GET("/admin/common/branch/getBranchListForRegion/:regionID", branchHandle.GetBranchListForRegion) //for ajax

	// r.GET("/admin/:ModulName/dashboard-fundraising", indexHandle.Index)
	// r.GET("/admin/:ModulName/dashboard-sacrifece", indexHandle.SacrifeceIndex)

	userGroup := r.Group("/admin/:ModulName/user")
	{
		userGroup.GET("/", userHandle.Index)
		userGroup.GET("index", userHandle.Index)
		userGroup.GET("create", userHandle.Create)
		userGroup.POST("store", userHandle.Store)
		userGroup.GET("edit/:UserID", userHandle.Edit)
		userGroup.GET("delete/:ID", userHandle.Delete)
		userGroup.POST("update", userHandle.Update)
		userGroup.GET("NewPasswordModalBox", userHandle.NewPasswordModalBox)
		userGroup.POST("NewPasswordAjax", userHandle.NewPasswordCreateModalBox)
		userGroup.POST("passportchange", userHandle.PassportChange)
	}

	loginGroup := r.Group("/admin/login")
	{
		loginGroup.GET("/", loginHandle.Login)
		//loginGroup.GET("password", login.SifreVer)
		loginGroup.POST("loginpost", loginHandle.LoginPost)
		loginGroup.GET("logout", loginHandle.Logout)
	}

	adminPost := r.Group("/admin/:ModulName/post")
	{
		adminPost.GET("/", posts.Index)
		adminPost.GET("index", posts.Index)
		adminPost.GET("create", posts.Create)
		adminPost.POST("store", posts.Store)
		adminPost.GET("edit/:postID", posts.Edit)
		adminPost.POST("update", posts.Update)
		adminPost.POST("upload", posts.Upload)
	}

	branchGroup := r.Group("/admin/:ModulName/branch")
	{
		branchGroup.GET("/", branchHandle.Index)
		branchGroup.GET("index", branchHandle.Index)

		branchGroup.GET("/create", branchHandle.Create)
		branchGroup.POST("/store", branchHandle.Store)
		branchGroup.GET("/edit/:ID", branchHandle.Edit)
		branchGroup.POST("/update", branchHandle.Update)

	}

	roleGroup := r.Group("/admin/:ModulName/roles")
	{
		roleGroup.GET("/knockout", roleHandle.IndexKnockout)
		roleGroup.GET("/", roleHandle.Index)
		roleGroup.GET("/create", roleHandle.Create)
		roleGroup.POST("/store", roleHandle.Store)
		roleGroup.GET("/edit/:ID", roleHandle.Edit)
		roleGroup.POST("/update", roleHandle.Update)
		roleGroup.GET("delete/:ID", roleHandle.Delete)
	}

	regionGroup := r.Group("/admin/:ModulName/region")
	{
		regionGroup.GET("/", regionHandle.Index)
		regionGroup.GET("index", regionHandle.Index)
		regionGroup.GET("/create", regionHandle.Create)
		regionGroup.POST("/store", regionHandle.Store)
		regionGroup.GET("/edit/:ID", regionHandle.Edit)
		regionGroup.POST("/update", regionHandle.Update)
	}

	modulesGroup := r.Group("/admin/:ModulName/modules")
	{
		modulesGroup.GET("/", modulesHandle.Index)
		modulesGroup.GET("index", modulesHandle.Index)
		modulesGroup.GET("/create", modulesHandle.Create)
		modulesGroup.POST("/store", modulesHandle.Store)
		modulesGroup.GET("/edit/:ID", modulesHandle.Edit)
		modulesGroup.POST("/update", modulesHandle.Update)
	}

	optionsGroup := r.Group("/admin/:ModulName/options")
	{
		optionsGroup.GET("/", optionsHandle.Index)
		optionsGroup.POST("update", optionsHandle.Update)
		optionsGroup.GET("receiptNo", optionsHandle.ReceiptNo)
	}

	appPort := os.Getenv("PORT")
	siteName := os.Getenv("SITENAME")
	if appPort == "" {
		appPort = "9090" //localhost
	}

	switch debugMode {
	case "RELEASE":
		//log.Fatal(autotls.Run(r, os.Getenv("SITENAME")))
		log.Fatal(r.RunTLS(":"+appPort, "/home/admin/conf/web/"+siteName+"/ssl/"+siteName+".crt", "/home/admin/conf/web/"+siteName+"/ssl/"+siteName+".key"))
	case "DEBUG":
		log.Fatal(r.Run(":" + appPort))
	case "TEST":
		log.Fatal(r.Run(":" + appPort))
	default:
		log.Fatal(r.Run(":" + appPort))
	}

}
