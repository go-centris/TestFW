package auth_mod

import (
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/lang"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncdatetime"
	"stncCms/app/domain/helpers/stnchelper"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/infrastructure/security"
	Iauth "stncCms/app/services/authServices_mod"
	Iregion "stncCms/app/services/regionServices_mod"
	"strconv"
)

// userControl constructor
type UserControl struct {
	IUser   Iauth.UserAppInterface
	IRegion Iregion.RegionAppInterface
	RoleApp Iauth.RoleAppInterface
}

//TODO: mail ile daha once uye olmusmu kontrolu olacak

const viewPathuserControl = "admin/auth/user/"

// InitUserControl userControl controller constructor
func InitUserControl(iUserApp Iauth.UserAppInterface, iIRegionApp Iregion.RegionAppInterface, iRolesApp Iauth.RoleAppInterface) *UserControl {
	return &UserControl{
		IUser:   iUserApp,
		IRegion: iIRegionApp,
		RoleApp: iRolesApp,
	}
}

// Index list
func (access *UserControl) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)

	var tarih stncdatetime.Inow
	var total int64
	access.IUser.Count(&total)
	userControlsPerPage := 10
	paginator := pagination.NewPaginator(c.Request, userControlsPerPage, total)
	offset := paginator.Offset()
	allData, _ := access.IUser.GetAllPagination(userControlsPerPage, offset)

	// var tarih stncdatetime.Inow
	// fmt.Println(tarih.TarihFullSQL("2020-05-21 05:08"))
	// fmt.Println(tarih.AylarListe("May"))
	// fmt.Println(tarih.Tarih())
	// //	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	//	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "İçerik Ekleme",
		"allData":      allData,
		"tarih":        tarih,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathuserControl+"index.html",
		viewData,
	)
}

// Create all list f
func (access *UserControl) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)
	region, _ := access.IRegion.GetAll()
	roles, _ := access.RoleApp.GetAll()

	//#json formatter #stncjson
	//empJSON, err := json.MarshalIndent(region, "", "  ")
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "İçerik Ekleme",
		"regions":      region,
		"roles":        roles,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"create.html",
		viewData,
	)
}

// Store save method
func (access *UserControl) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)
	roles, _ := access.RoleApp.GetAll()
	var userSave = userModel(c, "create", "")
	var userSavePostError = make(map[string]string)
	userSavePostError = userSave.Validate()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	if len(userSavePostError) == 0 {
		saveData, saveErr := access.IUser.Save(&userSave)
		if saveErr != nil {
			userSavePostError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/user/edit/"+lastID)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}

	viewData := pongo2.Context{
		"title":        "content add",
		"csrf":         csrf.GetToken(c),
		"err":          userSavePostError,
		"data":         userSave,
		"flashMsg":     flashMsg,
		"roles":        roles,
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"create.html",
		viewData,
	)

}

// Edit edit data
func (access *UserControl) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	flashMsg := stncsession.GetFlashMessage(c)
	var branchID, regionID uint64
	if userID, err := strconv.ParseUint(c.Param("UserID"), 10, 64); err == nil {
		if userData, err := access.IUser.GetByID(userID); err == nil {
			roles, _ := access.RoleApp.GetAll()
			region, _ := access.IRegion.GetAll()
			fmt.Println(userData.BranchID)

			if userData.BranchID == 0 {
				stncsession.SetFlashMessage("Bu kullanicinin branch adresi bulunmuyor lutfen tanimlayiniz", "danger", c)
				branchID = 0
				regionID = 0
			} else {
				dataUserForBranchID, _ := access.IUser.GetByUserForBranchID(userData.BranchID)
				branchID = dataUserForBranchID.BranchID
				regionID = dataUserForBranchID.RegionID
			}
			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"title":        "kullanıcı düzenleme",
				"data":         userData,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"regions":      region,
				"branchID":     branchID,
				"regionID":     regionID,
				"roles":        roles,
				"locale":       locale,
				"localeMenus":  menuLanguage,
				"modulNameUrl": modulName,
			}
			c.HTML(
				http.StatusOK,
				viewPathuserControl+"edit.html",
				viewData,
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Update data
// TODO: resim silmeyi unutma
// Delete data
func (access *UserControl) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var pass string
	var branchID, regionID uint64
	var userSavePostError = make(map[string]string)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("user")
	id := c.PostForm("userID")
	id2 := stnccollection.StringtoUint64(id)
	if userData, err := access.IUser.GetByID(id2); err == nil {
		pass = userData.Password
	}
	var userData = userModel(c, "edit", pass)
	userSavePostError = userData.Validate()
	roles, _ := access.RoleApp.GetAll()
	region, _ := access.IRegion.GetAll()

	if userData.BranchID == 0 {
		stncsession.SetFlashMessage("Bu kullanicinin branch adresi bulunmuyor lutfen tanimlayiniz", "danger", c)
		branchID = 0
		regionID = 0
	} else {
		dataUserForBranchID, _ := access.IUser.GetByUserForBranchID(userData.BranchID)
		branchID = dataUserForBranchID.BranchID
		regionID = dataUserForBranchID.RegionID
	}
	if len(userSavePostError) == 0 {
		_, saveErr := access.IUser.Update(&userData)
		if saveErr != nil {
			userSavePostError = saveErr
		}
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/"+viewPathuserControl+"edit/"+id)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":    "User Edit",
		"err":      userSavePostError,
		"csrf":     csrf.GetToken(c),
		"flashMsg": flashMsg,

		"data":         userData,
		"roles":        roles,
		"regions":      region,
		"branchID":     branchID,
		"regionID":     regionID,
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathuserControl+"edit.html",
		viewData,
	)
}

// OdemeEkleCreateModalBox takistler
func (access *UserControl) NewPasswordModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	userID := c.Query("userID")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "New Password",
		"userID":       userID,
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathuserControl+"NewPasswordModalBox.html",
		viewData,
	)
}

// ajax save method
func (access *UserControl) NewPasswordCreateModalBox(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, _ := lang.LoadLanguages("user")
	//TODO: sahte veri girisi kontrolu olcak
	//var KurbanID uint64
	//var kalanKurbanFiyati float64
	//KurbanID, _ = strconv.ParseUint(c.PostForm("kurbanID"), 10, 64)
	//kalanKurbanFiyati = access.kurbanApp.KalanUcret(KurbanID)
	//db := repository.DB
	//appKurban := repository.KurbanRepositoryInit(db)
	//var kisiID int
	//appKurban.GetKurbanKisiVarmi(KurbanID, &kisiID)
	//if kisiID == 1 {
	// sahte veri girişi yani kişi atanmamış kurbana ödeme yapmaya çalışıyor  TODO: bunun loglanması lazım
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "passpart change",
		"csrf":         csrf.GetToken(c),
		"status":       "err",
		"err":          "fk", // sahte veri girişi TODO: bunun loglanması lazım
		"errMsg":       "unexpected error ",
		"locale":       locale,
		"modulNameUrl": modulName,
	}
	c.JSON(http.StatusOK, viewData)
	//} else {
	//	if odeme, _, errorR := odemelerModel(kalanKurbanFiyati, c); errorR == nil {
	//		var savePostError = make(map[string]string)
	//		savePostError = odeme.Validate()
	//		fmt.Printf("%+v\n", odeme)
	//		if len(savePostError) == 0 {
	//			_, saveErr := access.OdemelerApp.Save(&odeme)
	//			stncsession.SetFlashMessage("Password changed successfully", "success", c)
	//			if saveErr != nil {
	//				savePostError = saveErr
	//			}
	//			viewData := pongo2.Context{
	//				"title":  "Password Change",
	//				"csrf":   csrf.GetToken(c),
	//				"err":    savePostError,
	//				"status": "ok",
	//				"path":   "/admin/kurban/edit/" + c.PostForm("kurbanID"),
	//				"id":     c.PostForm("kurbanID"),
	//				"post":   odeme,
	//			}
	//			c.JSON(http.StatusOK, viewData)
	//
	//		}
	//	}
	//}
}

// referansEkleAjax save method
func (access *UserControl) PassportChange(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	//locale, _ := lang.LoadLanguages("user")
	userID := c.PostForm("userID")
	pass := c.PostForm("pass")
	var status string = "ok"
	if pass == "" {
		status = "err"
	}

	encryptPass := security.PassGenerate(pass)
	access.IUser.SetUserPassword(stnccollection.StringtoUint64(userID), encryptPass)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "Password Change",
		"csrf":         csrf.GetToken(c),
		"status":       status,
		"modulNameUrl": modulName,
		//"errMsg": "beklenmeyen bir hata oluştu",
		//"locale": locale,
	}
	c.JSON(http.StatusOK, viewData)
}

// Delete data
func (access *UserControl) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	if postID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		access.IUser.Delete(postID)
		stncsession.SetFlashMessage("Success Delete", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/user")
		return
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// form post model
func userModel(c *gin.Context, form string, pass string) (user entity.Users) {
	//	var post = entit.Post{}
	user.ID = stnccollection.StringtoUint64(c.PostForm("ID"))
	user.Username = c.PostForm("Username")
	user.Email = c.PostForm("Email")
	user.FirstName = c.PostForm("FirstName")
	user.LastName = c.PostForm("LastName")
	user.Phone = c.PostForm("Phone")
	if form == "create" {
		pass := c.PostForm("PasswordNew")
		encryptPass := security.PassGenerate(pass)
		user.Password = encryptPass
	} else {
		user.Password = pass
	}
	user.RoleID = stnccollection.StringToint(c.PostForm("RoleID"))
	user.BranchID = stnccollection.StringToint(c.PostForm("branchID"))
	return user
}
