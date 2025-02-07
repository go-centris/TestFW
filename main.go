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
	"stncCms/app/web.api/controller/middleware"
	cms "stncCms/app/web/controller/cms_mod"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leonelquinteros/gotext"
	myPongoGinRender "github.com/stnc/myPongoGinRender/v5"
	csrf "github.com/utrack/gin-csrf"

	ctrl "stncCms/app/web/controller"
	//modSacrife "stncCms/app/web.api/controller/modSacrife"
	auth "stncCms/app/web/controller/auth_mod"
	common "stncCms/app/web/controller/common_mod"
	fundraising "stncCms/app/web/controller/fundraising_mod"
	region "stncCms/app/web/controller/region_mod"
	reportSacrife "stncCms/app/web/controller/reportSacrife_mod"
	sacrifice "stncCms/app/web/controller/sacrifice_mod"
)

var cacheControlSelman = false

// https://github.com/stnc-go/gobyexample/blob/master/pongo2render/render.go
func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	/* //bu sunucuda çalışıyor
		    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	        if err != nil {
	            log.Fatal(err)
	        }
	        environmentPath := filepath.Join(dir, ".env")
	        err = godotenv.Load(environmentPath)
	        fatal(err)
	*/

}

func main() {
	//redis details
	// redisService, err := cache.RedisDBInit(redisHost, redisPort, redisPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	debugMode := os.Getenv("MODE")

	//
	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	//@uses ./main -autoRelation
	autoRelation := flag.Bool("autoRelation", false, "db relation ")
	flag.Parse()

	if *autoRelation {
		fmt.Printf("\033[1;34m%s\033[0m", "-----------iliskiler kuruluyor-------------")
		services.AutoRelation()
		fmt.Printf("\033[1;34m%s\033[0m", "-----------iliskiler kuruldu-------------")
		return
	}

	// redisService, err := auth.RedisDBInit(redisHost, redisPort, redisPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// token := auth.NewToken()
	// upload := stncupload.NewFileUpload()
	//usersAPI := apiController.InitUsers(services.User, redisService.Auth, token)
	//postsAPI := apiController.InitPost(services.Post, services.User, redisService.Auth, token)

	indexHandle := sacrifice.InitDashboard(services.Dashboard)
	posts := cms.InitPost(services.Post, services.Cat, services.CatPost, services.Lang, services.User)

	//cat := controller.InitPost(services.Post, services.User)
	kurbanHandle := sacrifice.InitGsacrifice(services.Kurban, services.Kisiler, services.Media)

	odemelerHandle := sacrifice.InitOdemeler(services.Kodemeler, services.Kurban)

	kisilerHandle := sacrifice.InitGKisiler(services.Kisiler)

	GroupslarHandle := sacrifice.InitGruplar(services.HayvanBilgisi, services.Kurban, services.Gruplar, services.Options, services.Kodemeler, services.Kisiler)

	hayvanSatisYerleriHandle := sacrifice.InitHayvanSatisYerleri(services.HayvanSatisYerleri)

	hayvanBilgisiHandle := sacrifice.InitHayvanBilgisi(services.HayvanBilgisi, services.Media)

	loginHandle := auth.InitLogin(services.User)

	userHandle := auth.InitUserControl(services.User, services.Region, services.Role)
	roleHandle := auth.InitRoles(services.Permission, services.Modules, services.Role, services.RolePermission)

	optionsHandle := sacrifice.InitOptions(services.Options)

	branchHandle := region.InitBranch(services.Branch, services.Region)

	fundraisingypeHandle := fundraising.InitFundraisingType(services.FundraisingType, services.Region)

	fundraisingDonorsHandle := fundraising.InitFundraisingDonors(services.FundraisingDonors, services.FundraisingType)

	regionHandle := region.InitRegion(services.Region)

	reportHandle := reportSacrife.InitReport(services.Report)

	modulesHandle := common.InitModules(services.Modules)

	// region := controller.InitRegion(services.Region, services.Branch, services.Role, services.RolePermission)

	//authenticate := apiController.NewAuthenticate(services.User, redisService.Auth, token)

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

	r.Use(middleware.CORSMiddleware()) //For CORS

	//TODO: csrf kontrolu nasıl olacak
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
		c.Redirect(http.StatusMovedPermanently, "/admin/")
	})

	dashboardGroup := r.Group("/admin/")
	{
		dashboardGroup.GET("/", ctrl.Index)

	}

	r.GET("optionsDefault", sacrifice.OptionsDefault)
	r.GET("cacheReset", sacrifice.CacheReset)

	//public url
	r.GET("/kurbanBilgi/:slug", kurbanHandle.KurbanBilgi)
	r.GET("/fundraising", fundraisingypeHandle.IndexFront)
	r.GET("/fundraising/show/:ID", fundraisingypeHandle.ShowFront)

	r.GET("/admin/common/branch/getBranchListForRegion/:regionID", branchHandle.GetBranchListForRegion) //ajax ortak kullanim icin

	r.GET("/admin/:ModulName/dashboard-fundraising", indexHandle.Index)
	r.GET("/admin/:ModulName/dashboard-sacrifece", indexHandle.SacrifeceIndex)

	fundraisingTypeGroup := r.Group("/admin/:ModulName/fundraising-type")
	{
		fundraisingTypeGroup.GET("/", fundraisingypeHandle.Index)
		fundraisingTypeGroup.GET("/index", fundraisingypeHandle.Index)
		fundraisingTypeGroup.GET("/create", fundraisingypeHandle.Create)
		fundraisingTypeGroup.POST("/store", fundraisingypeHandle.Store)
		fundraisingTypeGroup.GET("/edit/:ID", fundraisingypeHandle.Edit)
		fundraisingTypeGroup.GET("/show/:ID", fundraisingypeHandle.Show)
		fundraisingTypeGroup.POST("/update", fundraisingypeHandle.Update)
		fundraisingTypeGroup.GET("/delete/:ID", fundraisingypeHandle.Delete)
	}

	fundraisingDonorsGroup := r.Group("/admin/:ModulName/fundraising-donors")
	{
		fundraisingDonorsGroup.GET("/", fundraisingDonorsHandle.Index)
		fundraisingDonorsGroup.GET("/index", fundraisingDonorsHandle.Index)
		fundraisingDonorsGroup.GET("/create", fundraisingDonorsHandle.Create)
		fundraisingDonorsGroup.POST("/store", fundraisingDonorsHandle.Store)
		fundraisingDonorsGroup.GET("/edit/:ID", fundraisingDonorsHandle.Edit)
		fundraisingDonorsGroup.GET("/quick-add/:ID", fundraisingDonorsHandle.Quick)
		fundraisingDonorsGroup.POST("/quick-store", fundraisingDonorsHandle.Quick)
		fundraisingDonorsGroup.POST("/update", fundraisingDonorsHandle.Update)
		fundraisingDonorsGroup.GET("/delete/:ID", fundraisingDonorsHandle.Delete)
	}
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

	//api routes
	// v1 := r.Group("/api/v1")
	// {
	// 	v1.POST("users", usersAPI.SaveUser)
	// 	v1.GET("users", usersAPI.GetUsers)

	// 	v1.GET("users/:user_id", usersAPI.GetUser)
	// 	v1.GET("postall", postsAPI.GetAllPost)
	// 	v1.POST("post", postsAPI.SavePost)
	// 	v1.PUT("post/:post_id", middleware.AuthMiddleware(), postsAPI.UpdatePost)
	// 	v1.GET("post/:post_id", postsAPI.GetPostAndCreator)
	// 	v1.DELETE("post/:post_id", middleware.AuthMiddleware(), postsAPI.DeletePost)
	// 	// cs.GET("/allcoins", controller.AllCoins())
	// 	// cs.GET("/mycoins/:id", controller.MyCoins())
	// 	// cs.GET("/create", controller.CreateCoin())
	// 	// cs.POST("/store", controller.StoreCoin())
	// 	// cs.GET("/edit/:id", controller.EditCoin())
	// 	// cs.PUT("/update/:id", controller.UpdateCoin())
	// 	// e.GET("/cpr/:slug", controller.CoinPreview())
	// 	// cs.DELETE("/:id", controller.DeleteCoin())
	// 	//authentication routes
	// 	v1.POST("login", authenticate.Login)
	// 	v1.POST("logout", authenticate.Logout)
	// 	v1.POST("refresh", authenticate.Refresh)
	// }

	//GENEL KURBAN<
	kurbanGroup := r.Group("/admin/:ModulName/kurban")
	{
		kurbanGroup.GET("/", kurbanHandle.Index)
		kurbanGroup.GET("/test", kurbanHandle.DBTest)
		kurbanGroup.GET("userList/:userID", kurbanHandle.UserList)
		kurbanGroup.GET("create", kurbanHandle.Create)

		kurbanGroup.POST("store", kurbanHandle.Store)
		kurbanGroup.GET("edit/:kurbanID", kurbanHandle.Edit)
		kurbanGroup.POST("update", kurbanHandle.Update)
		//api
		kurbanGroup.GET("odemeEkleCreateModalBox", kurbanHandle.OdemeEkleCreateModalBox)
		kurbanGroup.GET("grupLideri/:kurbanID", kurbanHandle.GrupLideriAta)
		kurbanGroup.GET("vekaletDurumu/:kurbanID", kurbanHandle.VekaletDurumu)
		kurbanGroup.GET("groupVekaletDurumu/:groupID", kurbanHandle.GroupVekaletDurumu)
		kurbanGroup.POST("upload/:ID", kurbanHandle.Upload)
		kurbanGroup.GET("media-delete/:ID", kurbanHandle.MediaDelete)
		kurbanGroup.GET("listDataTable", kurbanHandle.ListDataTable)
		kurbanGroup.GET("ListDataTableUserID/:userID", kurbanHandle.ListDataTableUserID)

		kurbanGroup.GET("quickAdd/", kurbanHandle.QuickAdd)
		kurbanGroup.POST("/quickAddstore", kurbanHandle.QuickUpdate)
	}

	kisilerGroup := r.Group("/admin/:ModulName/kisiler")
	{
		kisilerGroup.GET("/", kisilerHandle.Index)
		kisilerGroup.GET("index", kisilerHandle.Index)
		kisilerGroup.GET("IndexV1", kisilerHandle.IndexV1)
		kisilerGroup.GET("create", kisilerHandle.Create)
		kisilerGroup.POST("store", kisilerHandle.Store)
		kisilerGroup.GET("edit/:ID", kisilerHandle.Edit)
		kisilerGroup.GET("delete/:ID", kisilerHandle.Delete)
		kisilerGroup.POST("update", kisilerHandle.Update)
		kisilerGroup.POST("kisiAra", kisilerHandle.KisiAraAjax)

		kisilerGroup.GET("listDataTable", kisilerHandle.ListDataTable)
		kisilerGroup.GET("kisiAciklama/:ID", kisilerHandle.PersonComment)
		kisilerGroup.GET("kisiAciklamaEdit/:ID", kisilerHandle.PersonCommentEdit)
		//ajax methods
		kisilerGroup.GET("kisiBilgileriModalBox/:ID", kisilerHandle.KisiGosterModalBox)
		kisilerGroup.GET("referansCreateModalBox/:viewID", kisilerHandle.ReferansCreateModalBox)
		kisilerGroup.POST("kisiEkleAjax", kisilerHandle.KisiEkleAjax)
	}

	odemelerGroup := r.Group("/admin/:ModulName/odemeler")
	{
		odemelerGroup.GET("create", odemelerHandle.Create)
		odemelerGroup.POST("store", odemelerHandle.Store)
		odemelerGroup.POST("odemeEkleAjax", odemelerHandle.OdemeEkleAjax)
		odemelerGroup.GET("edit/:gID", odemelerHandle.Edit)
		odemelerGroup.GET("editSetAllGroupPrice/:groupID", odemelerHandle.EditSetAllGroupPrice)
		odemelerGroup.POST("update", odemelerHandle.Update)
		odemelerGroup.GET("makbuz/:gID", odemelerHandle.Makbuz)
		odemelerGroup.GET("delete/:odemeID", odemelerHandle.OdemeSil)
	}

	groups := r.Group("/admin/:ModulName/groups")
	{
		groups.GET("/", GroupslarHandle.Index)
		groups.GET("index", GroupslarHandle.Index)
		groups.GET("indexV1", GroupslarHandle.IndexV1)
		groups.GET("show/:ID", GroupslarHandle.Group)
		groups.GET("indexAjax", GroupslarHandle.IndexAjax)
		groups.GET("create", GroupslarHandle.Create)
		groups.POST("store", GroupslarHandle.Store)
		groups.GET("createEmpty", GroupslarHandle.CreateEmpty)
		groups.POST("storeEmpty", GroupslarHandle.StoreEmpty)

		groups.GET("edit/:kID", GroupslarHandle.Edit)
		groups.GET("editEmpty/:kID", GroupslarHandle.EditEmpty)
		groups.POST("update", GroupslarHandle.Update)
		// groups.POST("grupDegisim", GroupslarHandle.grupDegisim)
		groups.GET("degistir", GroupslarHandle.Degistir)
		groups.POST("degistir", GroupslarHandle.DegistirStore)
		groups.GET("grupBilgisiAPI/:kID", GroupslarHandle.GroupInfoApi)
		groups.GET("GroupInfoEmptyAnimalApi/:kID", GroupslarHandle.GroupInfoEmptyAnimalApi)
		groups.GET("hayvanata", GroupslarHandle.HayvanAtamasiYap)
		groups.POST("hayvanata", GroupslarHandle.HayvanAtamasiStore)
		groups.GET("groupsListeAjaxAgirlikTuru/:agirlikID", GroupslarHandle.GruplarListeAjaxAgirlikTuru)
		groups.GET("excel", GroupslarHandle.Excel)
		groups.GET("IndexAta", GroupslarHandle.IndexAta) //TODO: index ata ismi değişmeli bir anlam ifade etmiyor
		groups.POST("yerDegistir", GroupslarHandle.YerDegistir)
	}

	hayvanSatisYerleriGroup := r.Group("/admin/:ModulName/hayvanSatisYerleri")
	{
		hayvanSatisYerleriGroup.GET("/", hayvanSatisYerleriHandle.Index)
		hayvanSatisYerleriGroup.GET("index", hayvanSatisYerleriHandle.Index)
		hayvanSatisYerleriGroup.GET("create", hayvanSatisYerleriHandle.Create)
		hayvanSatisYerleriGroup.POST("store", hayvanSatisYerleriHandle.Store)
		hayvanSatisYerleriGroup.GET("edit/:kID", hayvanSatisYerleriHandle.Edit)
		hayvanSatisYerleriGroup.POST("update", hayvanSatisYerleriHandle.Update)
		hayvanSatisYerleriGroup.GET("listDataTable", hayvanSatisYerleriHandle.ListDataTable)
	}

	hayvanBilgisiGroup := r.Group("/admin/:ModulName/hayvanBilgisi")
	{
		hayvanBilgisiGroup.GET("/", hayvanBilgisiHandle.Index)
		hayvanBilgisiGroup.GET("index", hayvanBilgisiHandle.Index)
		hayvanBilgisiGroup.GET("listDataTable", hayvanBilgisiHandle.ListDataTable)
		hayvanBilgisiGroup.GET("create", hayvanBilgisiHandle.Create)
		hayvanBilgisiGroup.POST("store", hayvanBilgisiHandle.Store)
		hayvanBilgisiGroup.GET("edit/:kID", hayvanBilgisiHandle.Edit)
		hayvanBilgisiGroup.POST("update", hayvanBilgisiHandle.Update)
		hayvanBilgisiGroup.GET("hayvanListeAjax/:hID", hayvanBilgisiHandle.HayvanListeAjax)
		hayvanBilgisiGroup.POST("upload/:ID", hayvanBilgisiHandle.Upload)
		hayvanBilgisiGroup.GET("media-delete/:ID", hayvanBilgisiHandle.MediaDelete)
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

	reportGroup := r.Group("/admin/:ModulName/report")
	{
		reportGroup.GET("getAllUsersWhoAddedMostUser", reportHandle.GetAllUsersWhoAddedMostUser)
		reportGroup.GET("getAllUsersWhoAddedMostBranch", reportHandle.GetAllUsersWhoAddedMostBranch)
		reportGroup.GET("getAllCharitableWhoAddedMostSacrife", reportHandle.GetAllCharitableWhoAddedMostSacrife)
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
