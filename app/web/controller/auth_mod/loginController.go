package auth_mod

import (
	"fmt"
	"github.com/astaxie/beego/utils/pagination"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/pkg/helpers/lang"
	"stncCms/pkg/helpers/stnchelper"
	"stncCms/pkg/helpers/stncsession"
	"stncCms/pkg/infrastructure/security"
	Iauth "stncCms/app/services/authServices_mod"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

var (
	paginator = &pagination.Paginator{}
)

// Login constructor
type Login struct {
	IUser Iauth.UserAppInterface
}

// InitLogin login controller constructor
func InitLogin(iUser Iauth.UserAppInterface) *Login {
	return &Login{
		IUser: iUser,
	}
}

// sing in
func (login *Login) Login(c *gin.Context) {
	flashMsg := stncsession.GetFlashMessage(c)
	locale, _ := lang.LoadLanguages("user")
	//paginator := &pagination.Paginator{}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "Giriş",
		"flashMsg":     flashMsg,
		"locale":       locale,
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		"admin/auth/login/login.html",
		viewData,
	)
}

//func (login *Login) SifreVer(c *gin.Context) {
//	stncsession.IsLoggedInRedirect(c)
//	hak := c.Query("p")
//	name := c.Query("name")
//	sifre := security.PassGenerate(hak)
//	var post entity.Users
//	post.Email = name
//	post.Password = sifre
//	l.userApp.SaveUser(&post)
//	c.JSON(http.StatusOK, sifre)
//}

// sing in past data
func (login *Login) LoginPost(c *gin.Context) {
	locale, _ := lang.LoadLanguages("user")
	var user = entity.Users{}
	flashMsg := stncsession.GetFlashMessage(c)
	var savePostError = make(map[string]string)

	email := c.PostForm("Email")
	pass := c.PostForm("Password") //"111111-6" //sha1 hali cb5e6834e30cf762b38387db44c936daac667559
	user.Email = email
	user.Password = pass

	validateUser := user.ValidateLoginForm("login")
	if len(validateUser) > 0 {
		//	c.JSON(http.StatusUnprocessableEntity, validateUser)
		//stncsession.SetFlashMessage(validateUser, c)
		savePostError = validateUser
	} else {
		userData, result := login.IUser.GetUserByEmailAndPassword2(email, pass)
		if result == false {
			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
			stncsession.SetFlashMessage(locale.Get("Username or password is incorrect"), "warning", c)
			c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/login")
			return
		} else {
			stncsession.SetStoreUserID(c, userData.ID)
			stncsession.SetSession("UserName", userData.FirstName, c)
			//	c.SetCookie("username", "selmnn", 3600, "", "", false, true)
			c.Redirect(http.StatusMovedPermanently, "/")
			return
		}
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator": paginator,
		"err":       savePostError,
		"title":     "login",
		//"posts":     userData,
		"flashMsg":     flashMsg,
		"email":        email,
		"password":     pass,
		"locale":       locale,
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		"admin/auth/login/login.html",
		viewData,
	)
}

// sign in api implament
func (login *Login) LoginAPI(c *gin.Context) {
	var user = entity.Users{}
	email := "selmantunc@gmail.com"
	pass := "111111-6"
	user.Email = email
	user.Password = pass
	hashPassword := security.Hash(pass)
	fmt.Println("selman: " + string(hashPassword))
	//validate request:
	//var user *entity.User
	validateUser := user.ValidateLoginForm("login")
	if len(validateUser) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateUser)
		return
	}
	userData, userErr := login.IUser.GetUserByEmailAndPassword2(email, pass)
	if userErr != true {
		c.JSON(http.StatusInternalServerError, userErr)
		return
	} else {
		stncsession.SetStoreUserID(c, userData.ID)
		stncsession.SetSession("UserName", userData.FirstName, c)
		//	c.SetCookie("username", "selmnn", 3600, "", "", false, true)
	}
	fmt.Println(userData)
	c.Redirect(http.StatusMovedPermanently, "/")
	c.JSON(http.StatusOK, userData)
}

// Logout -sign out
func (au *Login) Logout(c *gin.Context) {
	stncsession.ClearUserID(c)
	c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
	//c.JSON(http.StatusOK, u)
}
