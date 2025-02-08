package sacrifice_mod

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"
	"stncCms/app/domain/helpers/stnchelper"
	"stncCms/app/domain/helpers/stncsession"
	Isacrife "stncCms/app/services/sacrifeServices_mod"

	"strconv"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/leonelquinteros/gotext"
	csrf "github.com/utrack/gin-csrf"
)

// Kisiler constructor
type Kisiler struct {
	IKisiler Isacrife.KisilerAppInterface
}

const viewPathGKisiler = "admin/sacrifece/kisiler/"

// InitGKisiler post controller constructor
func InitGKisiler(iKisiler Isacrife.KisilerAppInterface) *Kisiler {
	return &Kisiler{
		IKisiler: iKisiler,
	}
}

// Index list
func (access *Kisiler) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	lng := c.Query("lng")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	var getText *gotext.Locale
	// getText = gotext.NewLocale("locales", "en_US")

	if lng == "en" {
		getText = gotext.NewLocale("public/locales", "en_US")
	} else {
		getText = gotext.NewLocale("public/locales", "tr_TR")
	}

	// // Load domain '/path/to/locales/root/dir/es_UY/default.po'
	getText.AddDomain("default")
	// fmt.Println(getText.Get("Person List"))
	viewData := pongo2.Context{
		"paginator":    paginator,
		"flashMsg":     flashMsg,
		"locale":       getText,
		"modulNameUrl": modulName,
	}
	// c.JSON(http.StatusOK, viewData)

	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"index.html",
		viewData,
	)
}

// Index list
func (access *Kisiler) ListDataTable(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if result, err := access.IKisiler.ListDataTable(c); err == nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.AbortWithStatus(404)
		log.Println(err)
		return
	}

}

// Index list
func (access *Kisiler) IndexV1(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var tarih stncdatetime.Inow
	// var total int64
	// access.KisilerApp.Count(&total)
	// postsPerPage := 5
	// paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	// offset := paginator.Offset()
	// posts, _ := access.KisilerApp.GetAllPagination(postsPerPage, offset)
	posts, _ := access.IKisiler.GetAll()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "Kişi Ekleme",
		"posts":        posts,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"tarih":        tarih,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"indexList.html",
		viewData,
	)
}

// Create all list f
func (access *Kisiler) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var getText *gotext.Locale
	// getText = gotext.NewLocale("locales", "en_US")
	getText = gotext.NewLocale("public/locales", "tr_TR")
	getText.AddDomain("default")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"locale":       getText,
		"csrf":         csrf.GetToken(c),
		"flashMsg":     flashMsg,
		"kisiAdSoyad":  "",
		"MainMenu":     "krbnMenu",
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"create.html",
		viewData,
	)
}

// Store save method
func (access *Kisiler) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	// var AdSoyadint int
	var adSoyadstr string
	adSoyadstr = c.PostForm("AdSoyad")
	// AdSoyadint = 0
	// fmt.Println(AdSoyadstr)
	//burda gelen veri eğer daha once sistemde kayıtlı ise int tipinde gelecek
	//kayıtlı değilse string gelecek , store kısımında string veri gorunmesı ıcındır
	//dataControl ısımlı verının kotrol edılmesı geekıyor eger bu empty ıse kısı modalbox ı kapatmıstır yanı hıle yapıyor
	//data control ok ıse kısı dogru yoldadır asagıdakı degerler calıacak
	if adSoyadint, err := strconv.ParseUint(adSoyadstr, 10, 64); err == nil {
		kisiBilgileri, _ := access.IKisiler.GetByID(adSoyadint)
		adSoyadstr = kisiBilgileri.AdSoyad
	} else {
		adSoyadstr = c.PostForm("AdSoyad")
	}

	var kisiKaydet, _, _ = gKisilerModel(adSoyadstr, c)
	var savePostError = make(map[string]string)
	savePostError = kisiKaydet.Validate()
	if len(savePostError) == 0 {
		saveData, saveErr := access.IKisiler.Save(&kisiKaydet)
		if saveErr != nil {
			savePostError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/kisiler/edit/"+lastID)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}

	//fmt.Println(c.PostForm("ReferansKisi1"))
	var referansData *entity.Kisiler

	if c.PostForm("ReferansKisi1") != "0" {
		referansKisi1 := stnccollection.StringtoUint64(c.PostForm("ReferansKisi1"))
		referansData, _ = access.IKisiler.GetByID(referansKisi1)
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":         "Kisi Ekleme",
		"csrf":          csrf.GetToken(c),
		"err":           savePostError,
		"post":          kisiKaydet,
		"kisiAdSoyad":   adSoyadstr,
		"flashMsg":      flashMsg,
		"referansData1": referansData,
		"modulName":     referansData,
		"modulNameUrl":  modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"create.html",
		viewData,
	)
}

// Edit genel Kisiler düzenleme işler
// Edit edit data
func (access *Kisiler) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var tarih stncdatetime.Inow
	action := c.DefaultQuery("action", "ekle")
	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		if posts, err := access.IKisiler.GetByIDRel(ID); err == nil {
			kisiler, _ := access.IKisiler.GetAll()

			referansData, _ := access.IKisiler.GetByID(posts.ReferansKisi1)

			adSoyadstr := posts.AdSoyad

			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"title":        "Kişi ekleme",
				"post":         posts,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"tarih":        tarih,
				"kisiler":      kisiler,
				"referansData": referansData,
				"kisiAdSoyad":  adSoyadstr,
				"actionHref":   action,
				"modulNameUrl": modulName,
			}
			c.HTML(
				http.StatusOK,
				viewPathGKisiler+"edit.html",
				viewData,
			)

			return
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Update data
func (access *Kisiler) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	KisilerUpdate, id, _ := gKisilerModel(c.PostForm("AdSoyad"), c)
	adSoyadstr := c.PostForm("AdSoyad")
	var savePostError = make(map[string]string)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	savePostError = KisilerUpdate.Validate()
	kisiler, _ := access.IKisiler.GetAll()
	if len(savePostError) == 0 {

		_, saveErr := access.IKisiler.Update(&KisilerUpdate)
		if saveErr != nil {
			savePostError = saveErr
		}

		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/kisiler/edit/"+id)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}

	flashMsg := stncsession.GetFlashMessage(c)

	viewData := pongo2.Context{
		"title":        "Kisiler Düzenleme",
		"err":          savePostError,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"post":         KisilerUpdate,
		"kisiAdSoyad":  adSoyadstr,
		"kisiler":      kisiler,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"edit.html",
		viewData,
	)
}

// ReferansCreateModalBox create modalbox
func (access *Kisiler) ReferansCreateModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	viewID := c.Param("viewID")
	adSoyad := c.Query("metin")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "Kişi Ekleme",
		"viewID":       viewID,
		"adSoyad":      adSoyad,
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathGKisiler+"referansCreateModalBox.html",
		viewData,
	)
}

func (access *Kisiler) PersonComment(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		var aciklama string
		access.IKisiler.GetPersonComment(ID, &aciklama)
		// if posts, err := access.KisilerApp.GetPersonComment(ID, &aciklama); err == nil {
		// } else {
		// 	c.AbortWithError(http.StatusNotFound, err)
		// }
		viewData := pongo2.Context{
			"title":    "Açıklama",
			"status":   "ok",
			"aciklama": aciklama,
			"errMsg":   "",
		}
		c.JSON(http.StatusOK, viewData)
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Update data
func (access *Kisiler) PersonCommentEdit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	id := stnccollection.StringtoUint64(c.Query("id"))
	aciklama := c.Query("aciklama")
	access.IKisiler.SetAciklama(id, aciklama)
	return
	// viewData := pongo2.Context{
	// 	"title":  "Açıklama Düzenleme",
	// 	"status": "ok",
	// 	"csrf":   csrf.GetToken(c),
	// }
	// c.JSON(http.StatusOK, viewData)
}

// SearchAjax ajax search
func (access *Kisiler) KisiAraAjax(c *gin.Context) {
	q := c.PostForm("q")

	stncsession.IsLoggedInRedirect(c)

	posts, _ := access.IKisiler.Search(q)

	c.JSON(http.StatusOK, posts)

}

// referansEkleAjax save buraası referans kişi eklerken kullanılıyor
func (access *Kisiler) KisiEkleAjax(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	// var kurban, _, _ = gkurbanModel("referans", c)
	var kisiler = entity.Kisiler{}

	kisiler.UserID = stncsession.GetUserID2(c)
	adSoyad := c.PostForm("AdSoyad")

	kisiler.AdSoyad = adSoyad

	kisiler.Telefon = c.PostForm("Telefon")
	kisiler.Email = c.PostForm("Email")
	kisiler.Aciklama = c.PostForm("Aciklama")
	kisiler.Adres = c.PostForm("Adres")
	kisiler.Origins = stnccollection.StringToint(c.PostForm("Origins"))
	kisiler.AffinityType = stnccollection.StringToint(c.PostForm("AffinityType"))

	var savePostError = make(map[string]string)

	savePostError = kisiler.Validate()

	if len(savePostError) == 0 {
		saveData, saveErr := access.IKisiler.Save(&kisiler)
		if saveErr != nil {
			savePostError = saveErr
		}

		lastID := strconv.FormatUint(uint64(saveData.ID), 10)

		viewData := pongo2.Context{
			"title":    "Kişi Ekleme",
			"csrf":     csrf.GetToken(c),
			"lastID":   lastID,
			"viewID":   c.PostForm("viewID"),
			"username": saveData.AdSoyad,
			"tel":      saveData.Telefon,
			"err":      savePostError,
			"status":   "ok",
			"msg":      "Kayıt Başarı ile Eklendi",
		}
		c.JSON(http.StatusOK, viewData)
		return
	} else {
		viewData := pongo2.Context{
			"title":  "Kişi Ekleme",
			"csrf":   csrf.GetToken(c),
			"status": "error",
			"err":    savePostError,
		}
		c.JSON(http.StatusOK, viewData)
		return
	}

}

func (access *Kisiler) KisiGosterModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		if posts, err := access.IKisiler.GetByID(ID); err == nil {

			empJSON, err := json.MarshalIndent(posts, "", "  ")
			if err != nil {
				log.Fatalf(err.Error())
			}
			fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))

			referansData, _ := access.IKisiler.GetByID(posts.ReferansKisi1)

			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"title":        "Kişi ekleme",
				"post":         posts,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"modulNameUrl": modulName,
				"referansData": referansData,
			}
			c.HTML(
				http.StatusOK,
				viewPathGKisiler+"kisiBilgileriModalBox.html",
				viewData,
			)
			return
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

/***  POST MODEL   ***/
func gKisilerModel(adSoyad string, c *gin.Context) (data entity.Kisiler, idString string, err error) {

	id := c.PostForm("ID")

	idInt, _ := strconv.Atoi(id)

	var idN uint64

	idN = uint64(idInt)

	data.ID = idN
	data.UserID = stncsession.GetUserID2(c)
	// adSoyad := c.PostForm("AdSoyad")
	data.AdSoyad = adSoyad
	data.Telefon = c.PostForm("Telefon")
	data.Email = c.PostForm("Email")
	data.Adres = c.PostForm("Adres")
	data.Aciklama = c.PostForm("Aciklama")
	data.ReferansKisi1 = stnccollection.StringtoUint64(c.PostForm("ReferansKisi1"))
	data.Origins = stnccollection.StringToint(c.PostForm("Origins"))
	data.AffinityType = stnccollection.StringToint(c.PostForm("AffinityType"))
	// data.ReferansKisi2 = stnccollection.StringtoUint64(c.PostForm("ReferansKisi2"))
	return data, id, nil
}

// Delete data
func (access *Kisiler) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if postID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		var toplamKisi int
		var toplamRefKisi int
		access.IKisiler.HasKisiKurban(postID, &toplamKisi)
		access.IKisiler.HasKisiReferans(postID, &toplamRefKisi)
		fmt.Println("toplamKisi")
		fmt.Println(toplamKisi)

		fmt.Println("toplamRefKisi")
		fmt.Println(toplamRefKisi)
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		if toplamRefKisi > 0 {
			stncsession.SetFlashMessage("Kişi Referans Olarak Kayıtlıdır,Silemezsiniz", "danger", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/kisiler")
			return
		}
		if toplamKisi > 0 {
			stncsession.SetFlashMessage("Kişi Kurbana Kayıtlıdır,Silemezsiniz", "danger", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/kisiler")
			return
		} else {
			access.IKisiler.Delete(postID)
			stncsession.SetFlashMessage("Silindi", "success", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/kisiler")
			return
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}
