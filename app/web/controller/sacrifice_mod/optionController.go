package sacrifice_mod

import (
	"net/http"
	"stncCms/app/domain/helpers/lang"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stnchelper"
	"stncCms/app/domain/helpers/stncsession"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	//Iregion "stncCms/app/services/regionServices_mod"
	//Iauth "stncCms/app/services/authServices_mod"
	//Icms "stncCms/app/services/cmsServices_mod"
	//Icommon "stncCms/app/services/commonServices_mod"
	Isacrife "stncCms/app/services/sacrifeServices_mod"
) //Iregion "stncCms/app/services/regionServices_mod"

// Options constructor
type Options struct {
	IOption Isacrife.OptionsAppInterface
}

var (
	paginator = &pagination.Paginator{}
)

const viewPathOptions = "admin/sacrifece/options/"

// InitOptions post controller constructor
func InitOptions(iOptionApp Isacrife.OptionsAppInterface) *Options {
	return &Options{
		IOption: iOptionApp,
	}
}

// option list data
func (access *Options) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("options")
	numberOfShares := access.IOption.GetOption("hisse_adeti")
	salesUnitPrice1 := access.IOption.GetOption("satis_birim_fiyati_1")
	salesUnitPrice2 := access.IOption.GetOption("satis_birim_fiyati_2")
	salesUnitPrice3 := access.IOption.GetOption("satis_birim_fiyati_3")
	purchaseUnitPrice1 := access.IOption.GetOption("alis_birim_fiyati_1")
	purchaseUnitPrice2 := access.IOption.GetOption("alis_birim_fiyati_2")
	purchaseUnitPrice3 := access.IOption.GetOption("alis_birim_fiyati_3")
	whatsAppMsg := access.IOption.GetOption("whatsAppMsg")
	whatsAppMsgMap := access.IOption.GetOption("whatsAppMsgMap")
	whatsAppMsgEk1 := access.IOption.GetOption("whatsAppMsgEk1")
	otomatik_sira_buyukbas := access.IOption.GetOption("otomatik_sira_buyukbas")
	receipt_otomatik_sira_no := access.IOption.GetOption("receipt_otomatik_sira_no")
	cache_open_close := access.IOption.GetOption("cache_open_close")
	// dusukagirlik := access.IOption.GetOption("hayvan_dusuk_agirligi")
	// ortaagirlik := access.IOption.GetOption("hayvan_orta_agirligi")
	// yuksekagirlik := access.IOption.GetOption("hayvan_yuksek_agirligi")
	yearSacrifice := access.IOption.GetOption("kurban_yili")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":                    "Ayarlar",
		"csrf":                     csrf.GetToken(c),
		"numberOfShares":           numberOfShares,
		"satis_birim_fiyati_1":     salesUnitPrice1,
		"satis_birim_fiyati_2":     salesUnitPrice2,
		"satis_birim_fiyati_3":     salesUnitPrice3,
		"alis_birim_fiyati_1":      purchaseUnitPrice1,
		"alis_birim_fiyati_2":      purchaseUnitPrice2,
		"alis_birim_fiyati_3":      purchaseUnitPrice3,
		"whatsAppMsg":              whatsAppMsg,
		"whatsAppMsgEk1":           whatsAppMsgEk1,
		"otomatik_sira_buyukbas":   otomatik_sira_buyukbas,
		"receipt_otomatik_sira_no": receipt_otomatik_sira_no,
		"cache_open_close":         cache_open_close,
		// "hayvan_dusuk_agirligi":  dusukagirlik,
		// "hayvan_orta_agirligi":   ortaagirlik,
		// "hayvan_yuksek_agirligi": yuksekagirlik,
		"kurban_yili":    yearSacrifice,
		"flashMsg":       flashMsg,
		"whatsAppMsgMap": whatsAppMsgMap,
		"locale":         locale,
		"localeMenus":    menuLanguage,
		"modulNameUrl":   modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathOptions+"index.html",
		viewData,
	)
}

// Update list
func (access *Options) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	access.IOption.SetOption("hisse_adeti", c.PostForm("hisse_adeti"))
	access.IOption.SetOption("satis_birim_fiyati_1", c.PostForm("satis_birim_fiyati_1"))
	access.IOption.SetOption("satis_birim_fiyati_2", c.PostForm("satis_birim_fiyati_2"))
	access.IOption.SetOption("satis_birim_fiyati_3", c.PostForm("satis_birim_fiyati_3"))
	access.IOption.SetOption("alis_birim_fiyati_1", c.PostForm("alis_birim_fiyati_1"))
	access.IOption.SetOption("alis_birim_fiyati_2", c.PostForm("alis_birim_fiyati_2"))
	access.IOption.SetOption("alis_birim_fiyati_3", c.PostForm("alis_birim_fiyati_3"))
	access.IOption.SetOption("whatsAppMsg", c.PostForm("whatsAppMsg"))
	access.IOption.SetOption("whatsAppMsgMap", c.PostForm("whatsAppMsgMap"))
	access.IOption.SetOption("whatsAppMsgEk1", c.PostForm("whatsAppMsgEk1"))
	// access.IOption.SetOption("hayvan_dusuk_agirligi", c.PostForm("hayvan_dusuk_agirligi"))
	// access.IOption.SetOption("hayvan_orta_agirligi", c.PostForm("hayvan_orta_agirligi"))
	// access.IOption.SetOption("hayvan_yuksek_agirligi", c.PostForm("hayvan_yuksek_agirligi"))
	access.IOption.SetOption("kurban_yili", c.PostForm("kurban_yili"))
	access.IOption.SetOption("otomatik_sira_buyukbas", c.PostForm("otomatik_sira_buyukbas"))
	access.IOption.SetOption("receipt_otomatik_sira_no", c.PostForm("receipt_otomatik_sira_no"))
	access.IOption.SetOption("cache_open_close", c.PostForm("cache_open_close"))
	stncsession.SetStore(c, "cache_open_close", c.PostForm("cache_open_close"))
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/options")
	return

}

// Receipt No generator  tr: makbuz no üretir
func (access *Options) ReceiptNo(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	receiptAutomaticQueueNo := access.IOption.GetOption("receipt_automatic_queue_no ")
	receiptAutomaticQueueNoint := stnccollection.StringToint(receiptAutomaticQueueNo) + 1
	receiptAutomaticQueueNostr := stnccollection.IntToString(receiptAutomaticQueueNoint)
	access.IOption.SetOption("receipt_automatic_queue_no", receiptAutomaticQueueNostr)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":                   "Otomatik Sıra No",
		"status":                  "ok",
		"receiptAutomaticQueueNo": receiptAutomaticQueueNo,
		"errMsg":                  "",
		"modulNameUrl":            modulName,
	}
	c.JSON(http.StatusOK, viewData)
	return

}
