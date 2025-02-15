package common_mod

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/leonelquinteros/gotext"
	csrf "github.com/utrack/gin-csrf"
	"net/http"

	"stncCms/pkg/helpers/stnchelper"
	"stncCms/pkg/helpers/stncsession"

	Icommon "stncCms/app/modules/services"
	"strconv"

	modulesEntity "stncCms/app/modules/entity"
)

// Modules constructor
type Modules struct {
	Iregion Icommon.ModulesAppInterface
}

const viewPathModules = "admin/common/modules/"

func InitModules(iregion Icommon.ModulesAppInterface) *Modules {
	return &Modules{
		Iregion: iregion,
	}
}

// Index  list
func (access *Modules) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var total int64
	access.Iregion.GetAllPaginateCount(&total)
	postsPerPage := 15
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.Iregion.GetAllPaginate(postsPerPage, offset)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"posts":        posts,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathModules+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Modules) Create(c *gin.Context) {
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
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathModules+"create.html",
		viewData,
	)
}

// Store save method
func (access *Modules) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var data, _, _ = modulesModel(c)
	var savePostError = make(map[string]string)
	savePostError = data.Validate()
	if len(savePostError) == 0 {
		saveData, saveErr := access.Iregion.Save(&data)
		if saveErr != nil {
			savePostError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/common/modules/edit/"+lastID)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"csrf":         csrf.GetToken(c),
		"err":          savePostError,
		"post":         data,
		"flashMsg":     flashMsg,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathModules+"create.html",
		viewData,
	)
}

// Edit edit data
func (access *Modules) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		if posts, err := access.Iregion.GetByID(ID); err == nil {
			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"post":         posts,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"modulNameUrl": modulName,
			}
			c.HTML(
				http.StatusOK,
				viewPathModules+"edit.html",
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
func (access *Modules) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	var savePostError = make(map[string]string)
	regionUpdate, id, _ := modulesModel(c)
	savePostError = regionUpdate.Validate()
	if len(savePostError) == 0 {
		_, saveErr := access.Iregion.Update(&regionUpdate)
		if saveErr != nil {
			savePostError = saveErr
		}
		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/modules/edit/"+id)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	flashMsg := stncsession.GetFlashMessage(c)

	viewData := pongo2.Context{
		"err":          savePostError,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"post":         regionUpdate,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathModules+"edit.html",
		viewData,
	)
}

/***  modulesModel    ***/
func modulesModel(c *gin.Context) (data modulesEntity.Modules, idString string, err error) {
	id := c.PostForm("ID")
	idInt, _ := strconv.Atoi(id)
	data.ID = idInt
	data.UserID = stncsession.GetUserID2(c)
	data.ModulName = c.PostForm("ModulName")
	return data, id, nil
}
