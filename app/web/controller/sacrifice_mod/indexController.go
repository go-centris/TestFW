package sacrifice_mod

import (
	"net/http"
	"stncCms/pkg/cache"
	"stncCms/app/domain/entity"
	"stncCms/pkg/helpers/lang"
	"stncCms/pkg/helpers/rbac"
	"stncCms/pkg/helpers/stnchelper"
	"stncCms/pkg/helpers/stncsession"
	repository "stncCms/app/domain/repository/cacheRepository"
	Isacrife "stncCms/app/services/sacrifeServices_mod"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const viewPathIndex = "admin/sacrifece/dashboard/"

// Dashboard constructor
type Dashboard struct {
	IDashboard Isacrife.DashboardAppInterface
}

func InitDashboard(iDashboardApp Isacrife.DashboardAppInterface) *Dashboard {
	return &Dashboard{
		IDashboard: iDashboardApp,
	}
}

// Index all list f
func (access *Dashboard) SacrifeceIndex(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	permissionList := rbac.RbacCheckForComponent(c, "dashboard")
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("dashboard")
	res := Dashboard{IDashboard: access.IDashboard}
	grandTotalsComponentValue := res.GrandTotalsComponent()
	shareHoldersComponentValue := res.ShareHoldersComponent()
	charitableWhoAddedMostSacrife, _ := access.IDashboard.CharitableWhoAddedMostSacrife()         //En cok kurban kestiren hayirsever
	usersWhoAddedMostSacrifeAndBranch, _ := access.IDashboard.UsersWhoAddedMostSacrifeAndBranch() //En cok kurban ekleyen  subemiz
	usersWhoAddedMostSacrifeAndUser, _ := access.IDashboard.UsersWhoAddedMostSacrifeAndUser()     //En cok kurban ekleyen  hocamiz
	viewData := pongo2.Context{
		"title":                             locale.Get("Dashboard"),
		"flashMsg":                          flashMsg,
		"csrf":                              csrf.GetToken(c),
		"locale":                            locale,
		"grandTotalsComponentValue":         grandTotalsComponentValue,
		"shareHoldersComponentValue":        shareHoldersComponentValue,
		"charitableWhoAddedMostSacrife":     charitableWhoAddedMostSacrife,
		"usersWhoAddedMostSacrifeAndBranch": usersWhoAddedMostSacrifeAndBranch,
		"usersWhoAddedMostSacrifeAndUser":   usersWhoAddedMostSacrifeAndUser,
		"localeMenus":                       menuLanguage,
		"permissionList":                    permissionList,
		"modulNameUrl":                      "sacrifece",
	}
	c.HTML(
		http.StatusOK,
		viewPathIndex+"dashboard.html",
		viewData,
	)
}
func (access *Dashboard) Index(c *gin.Context) {

}

// grandTotalsComponent structure
type grandTotalsComponent struct {
	TotalSacrife               int64
	SharedSacrifeTotal         int64
	TotalPrice                 float64
	SacrifeSharedPriceTotal    float64
	RemainingDebt              float64
	SharedSacrifeRemainingDebt float64
}

// GrandTotalsComponent
func (access *Dashboard) GrandTotalsComponent() (data grandTotalsComponent) {
	var totalPrice, sacrifeSharedPriceTotal, remainingDebt, sharedSacrifeRemainingDebt float64
	var totalSacrife, sharedSacrifeTotal int64
	access.IDashboard.TotalPrice(&totalPrice)
	access.IDashboard.SacrifeSharedPriceTotal(&sacrifeSharedPriceTotal)
	access.IDashboard.TotalSacrife(&totalSacrife)
	access.IDashboard.RemainingDebt(&remainingDebt)
	access.IDashboard.SharedSacrifeTotal(&sharedSacrifeTotal)
	access.IDashboard.SharedSacrifeRemainingDebt(&sharedSacrifeRemainingDebt)
	res := grandTotalsComponent{
		TotalPrice:                 totalPrice,
		SacrifeSharedPriceTotal:    sacrifeSharedPriceTotal,
		TotalSacrife:               totalSacrife,
		RemainingDebt:              remainingDebt,
		SharedSacrifeRemainingDebt: sharedSacrifeRemainingDebt,
		SharedSacrifeTotal:         sharedSacrifeTotal,
	}
	return res
}

// shareHoldersComponent structure
type shareHoldersComponent struct {
	ShareSacrifeCount2021          int64   // sadece hisseli kurban miktari yil 2021
	ShareSacrifeCount2022          int64   // sadece hisseli kurban miktari yil 2022
	ShareSacrifeCount2023          int64   // sadece hisseli kurban miktari yil 2023
	SharedSacrifeRemainingDebt2021 float64 //sadece hisseli  Kalan toplam Borc 2021
	SharedSacrifeRemainingDebt2022 float64 //sadece hisseli  Kalan toplam Borc 2022
	SharedSacrifeRemainingDebt2023 float64 //sadece hisseli  Kalan toplam Borc 2023
	SacrifeSharedPriceTotal2021    float64 //sadece hisseli kurban paRASI 2021
	SacrifeSharedPriceTotal2022    float64 //sadece hisseli kurban paRASI 2022
	SacrifeSharedPriceTotal2023    float64 //sadece hisseli kurban paRASI 2023

}

// ShareHoldersComponent is the runner
func (access *Dashboard) ShareHoldersComponent() (data shareHoldersComponent) {
	var sharedSacrifeRemainingDebt2021, sharedSacrifeRemainingDebt2022, sharedSacrifeRemainingDebt2023, sacrifeSharedPriceTotal2021, sacrifeSharedPriceTotal2022, sacrifeSharedPriceTotal2023 float64
	var shareSacrifeCount2021, shareSacrifeCount2022, shareSacrifeCount2023 int64
	access.IDashboard.ShareSacrifeCount2021(&shareSacrifeCount2021)
	access.IDashboard.ShareSacrifeCount2022(&shareSacrifeCount2022)
	access.IDashboard.ShareSacrifeCount2023(&shareSacrifeCount2023)
	access.IDashboard.SharedSacrifeRemainingDebt2021(&sharedSacrifeRemainingDebt2021)
	access.IDashboard.SharedSacrifeRemainingDebt2022(&sharedSacrifeRemainingDebt2022)
	access.IDashboard.SharedSacrifeRemainingDebt2023(&sharedSacrifeRemainingDebt2023)
	access.IDashboard.SacrifeSharedPriceTotal2021(&sacrifeSharedPriceTotal2021)
	access.IDashboard.SacrifeSharedPriceTotal2022(&sacrifeSharedPriceTotal2022)
	access.IDashboard.SacrifeSharedPriceTotal2023(&sacrifeSharedPriceTotal2023)
	res := shareHoldersComponent{
		ShareSacrifeCount2021:          shareSacrifeCount2021,
		ShareSacrifeCount2022:          shareSacrifeCount2022,
		ShareSacrifeCount2023:          shareSacrifeCount2023,
		SharedSacrifeRemainingDebt2021: sharedSacrifeRemainingDebt2021,
		SharedSacrifeRemainingDebt2022: sharedSacrifeRemainingDebt2022,
		SharedSacrifeRemainingDebt2023: sharedSacrifeRemainingDebt2023,
		SacrifeSharedPriceTotal2021:    sacrifeSharedPriceTotal2021,
		SacrifeSharedPriceTotal2022:    sacrifeSharedPriceTotal2022,
		SacrifeSharedPriceTotal2023:    sacrifeSharedPriceTotal2023,
	}
	return res
}

// Index all list f
func Index2(c *gin.Context) {
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "Posts",
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathIndex+"index.html",
		viewData,
	)
}

// OptionsDefault all list f
func OptionsDefault(c *gin.Context) {
	// stncsession.IsLoggedInRedirect(c)

	//buraya bir oprion otılacak bunlar giriş yaptıktan sonra veri varmı yok mu bakacak

	db := repository.DB

	option1 := entity.Options{OptionName: "siteUrl", OptionValue: "http://localhost:8888/"}
	db.Debug().Create(&option1)

	option2 := entity.Options{OptionName: "kurban_yili", OptionValue: "2022"}
	db.Debug().Create(&option2)

	option3 := entity.Options{OptionName: "hisse_adeti", OptionValue: "7"}
	db.Debug().Create(&option3)

	option4 := entity.Options{OptionName: "satis_birim_fiyati_1", OptionValue: "20"}
	db.Debug().Create(&option4)

	option5 := entity.Options{OptionName: "satis_birim_fiyati_2", OptionValue: "25"}
	db.Debug().Create(&option5)

	option6 := entity.Options{OptionName: "satis_birim_fiyati_3", OptionValue: "30"}
	db.Debug().Create(&option6)

	option7 := entity.Options{OptionName: "hayvan_dusuk_agirligi", OptionValue: "0-200"}
	db.Debug().Create(&option7)

	option78 := entity.Options{OptionName: "hayvan_orta_agirligi", OptionValue: "200-600"}
	db.Debug().Create(&option78)

	option786 := entity.Options{OptionName: "hayvan_yuksek_agirligi", OptionValue: "600-1500"}
	db.Debug().Create(&option786)

	option8 := entity.Options{OptionName: "alis_birim_fiyati_1", OptionValue: "10"}
	db.Debug().Create(&option8)

	option9 := entity.Options{OptionName: "alis_birim_fiyati_2", OptionValue: "15"}
	db.Debug().Create(&option9)

	option10 := entity.Options{OptionName: "alis_birim_fiyati_3", OptionValue: "20"}
	db.Debug().Create(&option10)

	option11 := entity.Options{OptionName: "otomatik_sira_buyukbas", OptionValue: "1"}
	db.Debug().Create(&option11)

	option12 := entity.Options{OptionName: "otomatik_sira_kuyukbas", OptionValue: "1"}
	db.Debug().Create(&option12)

	option13 := entity.Options{OptionName: "whatsAppMsg", OptionValue: "Merhaba Efendim, Kurbanınız ile ilgili bilgilere bu [link] adresden ulaşabilirsiniz."}
	db.Debug().Create(&option13)

	option14 := entity.Options{OptionName: "whatsAppMsgMap", OptionValue: "Merhaba Efendim bize bu adresden ulaşın "}
	db.Debug().Create(&option14)

	option15 := entity.Options{OptionName: "whatsAppMsgEk1", OptionValue: "ek mesaj "}
	db.Debug().Create(&option15)

	//mutluerF9E
	user := entity.Users{FirstName: "Sel", LastName: "t", Email: "hk@gmail.com", Password: "4544bcb2ce39fe656c64f0860895bdaccff7b8c0"} //mutluerF9E
	db.Debug().Create(&user)


	ModulKurban := entity.Modules{ModulName: "kurban", Status: 1, UserID: 1}
	db.Debug().Create(&ModulKurban)

	Modulodemeler := entity.Modules{ModulName: "odemeler", Status: 1, UserID: 1}
	db.Debug().Create(&Modulodemeler)

	ModulHayvanBilgisi := entity.Modules{ModulName: "hayvanBilgisi", Status: 1, UserID: 1}
	db.Debug().Create(&ModulHayvanBilgisi)

	ModulDashborad := entity.Modules{ModulName: "dashboard", Status: 1, UserID: 1}
	db.Debug().Create(&ModulDashborad)

	ModulAyarlar := entity.Modules{ModulName: "ayarlar", Status: 1, UserID: 1}
	db.Debug().Create(&ModulAyarlar)

	ModulGruplar := entity.Modules{ModulName: "gruplar", Status: 1, UserID: 1}
	db.Debug().Create(&ModulGruplar)

	ModulhayvanSatisYerleri := entity.Modules{ModulName: "hayvanSatisYerleri", Status: 1, UserID: 1}
	db.Debug().Create(&ModulhayvanSatisYerleri)

	ModulKisiler := entity.Modules{ModulName: "kisiler", Status: 1, UserID: 1}
	db.Debug().Create(&ModulKisiler)

	Modulkullanici := entity.Modules{ModulName: "kullanici", Status: 1, UserID: 1}
	db.Debug().Create(&Modulkullanici)



	c.JSON(http.StatusOK, "yapıldı")
}

// OptionsDefault all list f
func CacheReset(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	stncsession.SetFlashMessage("Cache Temizlendi", "success", c)
	redisClient := cache.RedisDBInit()
	redisClient.FlushAll()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/options/")
}
