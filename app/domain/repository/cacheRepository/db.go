package cacheRepository

import (
	"fmt"
	"os"
	"stncCms/app/domain/entity"
	Iauth "stncCms/app/services/authServices_mod"
	Icms "stncCms/app/services/cmsServices_mod"
	Icommon "stncCms/app/services/commonServices_mod"

	Iregion "stncCms/app/services/regionServices_mod"
	Ireport "stncCms/app/services/reportSacrifeServices_mod"
	Isacrife "stncCms/app/services/sacrifeServices_mod"

	"github.com/hypnoglow/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // here
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

var DB *gorm.DB

// Repositories strcut
type Repositories struct {
	User               Iauth.UserAppInterface
	Role               Iauth.RoleAppInterface
	Permission         Iauth.PermissionAppInterface
	RolePermission     Iauth.RolePermissionAppInterface

	Dashboard          Isacrife.DashboardAppInterface

	Kisiler            Isacrife.KisilerAppInterface

	Region             Iregion.RegionAppInterface
	Branch             Iregion.BranchAppInterface
	Modules            Icommon.ModulesAppInterface

	Post    Icms.PostAppInterface
	Cat     Icms.CatAppInterface
	CatPost Icms.CatPostAppInterface
	Media   Isacrife.MediaAppInterface
	Report  Ireport.ReportAppInterface

	Lang              Icommon.LanguageAppInterface
	Options           Isacrife.OptionsAppInterface


	DB *gorm.DB
}


func DbConnect() *gorm.DB {
	dbdriver := os.Getenv("DB_DRIVER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")
	debug := os.Getenv("MODE")
	//	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword) //bu postresql

	//DBURL := "root:sel123C#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //mysql
	var DBURL string

	if dbdriver == "mysql" {
		DBURL = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	} else if dbdriver == "postgres" {
		DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", dbHost, dbPort, dbUser, dbPassword, dbName) //Build connection string
	}

	// dsn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s sslmode=disable",
	// HOST, PORT, username, password, database)

	db, err := gorm.Open(dbdriver, DBURL)
	db.Set("gorm:table_options", "charset=utf8")
	// }

	// db, err := gorm.Open(dbdriver, DBURL)
	//nunlar gorm 2 versionunda prfexi falan var
	// db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "krbn_", // table name prefix, table for `User` would be `t_users`
	// 		SingularTable: true,    // use singular table name, table for `User` would be `user` with this option enabled
	// 	},
	// 	// Logger: gorm_logrus.New(),
	// })

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if debug == "DEBUG" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
		log := zap.NewExample()
		db.SetLogger(gormzap.New(log, gormzap.WithLevel(zap.DebugLevel)))
	} else if debug == "DEBUG" || debug == "TEST" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
	} else if debug == "RELEASE" {
		fmt.Println(debug)
		db.LogMode(false)
	}
	DB = db

	db.SingularTable(true)

	return db
}

//https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/

// RepositoriesInit initial
func RepositoriesInit(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		User:               UserRepositoryInit(db),
		Permission:         PermissionRepositoryInit(db),
		Role:               RoleRepositoryInit(db),
		RolePermission:     RolePermissionRepositoryInit(db),
		Dashboard:          DashboardRepositoryInit(db),

		Kisiler:            KisilerRepositoryInit(db),

		Modules:            ModulesRepositoryInit(db),

		Region: RegionRepositoryInit(db),
		Branch: BranchRepositoryInit(db),

		Report: ReportRepositoryInit(db),

		Post:    PostRepositoryInit(db),
		Cat:     CatRepositoryInit(db),
		CatPost: CatPostRepositoryInit(db),
		Media:   MediaRepositoryInit(db),

		Lang:              LanguageRepositoryInit(db),
		Options:           OptionRepositoryInit(db),

		DB:                db,
	}, nil
}

// CategoriesBranchJoin: CategoriesBranchJoinRepositoryInit(db),

//Close closes the  database connection
// func (s *Repositories) Close() error {
// 	return s.db.Close()
// }

// AutoRelation This migrate all tables
func (s *Repositories) AutoRelation() error {
	s.DB.AutoMigrate(&entity.Users{}, &entity.Role{}, &entity.Permission{}, &entity.RolePermisson{},
		&entity.Languages{}, &entity.Modules{}, &entity.Notes{}, &entity.Options{}, &entity.Currency{},
		&entity.SacrificeKurbanlar{},
		&entity.SacrificeOdemeler{}, &entity.SacrificeGruplar{}, &entity.Kisiler{}, &entity.Users{},
		&entity.SacrificeHayvanSatisYerleri{},
		&entity.SacrificeHayvanBilgisi{}, &entity.Region{}, &entity.Branches{}, &entity.Notification{}, &entity.NotificationTemplate{},
		&entity.Post{}, &entity.Categories{}, &entity.CategoryPosts{}, &entity.Media{}, &entity.FundraisingDonors{}, &entity.FundraisingType{})

	s.DB.Model(&entity.Permission{}).AddForeignKey("modul_id", "modules(id)", "CASCADE", "CASCADE")     // one to many (one=modules) (many=Permission)
	s.DB.Model(&entity.RolePermisson{}).AddForeignKey("role_id", "rbca_role(id)", "CASCADE", "CASCADE") // one to many (one=rbca_role) (many=RolePermisson)
	s.DB.Model(&entity.Branches{}).AddForeignKey("region_id", "br_region(id)", "CASCADE", "CASCADE")    // one to many (one=region) (many=branches)

	s.DB.Model(&entity.SacrificeHayvanBilgisi{}).AddForeignKey("hayvan_satis_yerleri_id", "sacrifice_hayvan_satis_yerleri(id)", "CASCADE", "CASCADE") // one to one (one=hayvan_satis_yerleri) (one=HayvanBilgisi)
	s.DB.Model(&entity.SacrificeOdemeler{}).AddForeignKey("kurban_id", "sacrifice_kurbanlar(id)", "CASCADE", "CASCADE")                               // one to many (one=kurbanlar) (many=odemeler)
	s.DB.Model(&entity.SacrificeKurbanlar{}).AddForeignKey("grup_id", "sacrifice_gruplar(id)", "CASCADE", "CASCADE")                                  // one to many (one=gruplar) (many=kurbanlar)
	return s.DB.Model(&entity.SacrificeKurbanlar{}).AddForeignKey("kisi_id", "sacrifice_kisiler(id)", "CASCADE", "CASCADE").Error                     // one to many (one=kisiler) (many=kurbanlar)

}

func (s *Repositories) Automigrate() error {
	return s.DB.AutoMigrate(&entity.Users{}, &entity.Role{}, &entity.Permission{}, &entity.RolePermisson{},
		&entity.Languages{}, &entity.Modules{}, &entity.Notes{}, &entity.Options{}, &entity.Currency{},
		&entity.SacrificeKurbanlar{},
		&entity.SacrificeOdemeler{}, &entity.SacrificeGruplar{}, &entity.Kisiler{}, &entity.Users{},
		&entity.SacrificeHayvanSatisYerleri{},
		&entity.SacrificeHayvanBilgisi{}, &entity.Region{}, &entity.Branches{}, &entity.Notification{}, &entity.NotificationTemplate{},
		&entity.Post{}, &entity.Categories{}, &entity.CategoryPosts{}, &entity.Media{}, &entity.FundraisingDonors{}, &entity.FundraisingType{}).Error
}

/*func GetAllStatusFindAndAgirlikTipiGroup(db *gorm.DB, durum int, agirlikTipi int) ([]entity.SacrificeGruplar, error) {
	repo := repository.GruplarRepositoryInit(db)
	datas, _ := repo.GetAllStatusFindAndAgirlikTipi(durum, agirlikTipi)
	return datas, nil
}

//GetAllStatusFind all data -- TODO: buradaki yil olayi degisecek
func (r *GruplarRepo) GetAllStatusFindAndAgirlikTipi(durum int, agirlikTipi int) ([]entity.SacrificeGruplar, error) {
	var data []entity.SacrificeGruplar
	access := repository.OptionRepositoryInit(r.db)
	cacheControl := access.GetOption("cache_open_close")

	if cacheControl == "false" {
		data, _ = GetAllStatusFindAndAgirlikTipiGroup(r.db, durum, agirlikTipi)
	} else {
		redisClient := cache.RedisDBInit()

		key := "GetAllStatusFindAndAgirlikTipiGroup_" + stnccollection.IntToString(durum) + stnccollection.IntToString(agirlikTipi)

		cachedProducts, err := redisClient.GetKey(key)

		if err != nil {
			data, _ = GetAllStatusFindAndAgirlikTipiGroup(r.db, durum, agirlikTipi)
			err = redisClient.SetKey(key, data, time.Minute*7200) //7200 5 gun eder
			fmt.Println("key olustur")
			if err != nil {
				fmt.Println("hata baş")
			}
			return data, nil
		}
		err = json.Unmarshal(cachedProducts, &data)
		if err != nil {
			fmt.Println("hata son")
		}
	}
	return data, nil

}*/
