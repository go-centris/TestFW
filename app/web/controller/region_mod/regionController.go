package region_mod

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/leonelquinteros/gotext"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stnchelper"
	"stncCms/app/domain/helpers/stncsession"
	Iregion "stncCms/app/services/regionServices_mod"
	"strconv"
)

// Region constructor
type Region struct {
	Iregion Iregion.RegionAppInterface
}

const viewPathRegion = "admin/region/region/"

func InitRegion(iregion Iregion.RegionAppInterface) *Region {
	return &Region{
		Iregion: iregion,
	}
}

// Index  list
func (access *Region) Index(c *gin.Context) {
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
		viewPathRegion+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Region) Create(c *gin.Context) {
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
		viewPathRegion+"create.html",
		viewData,
	)
}

// Store save method
func (access *Region) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)

	var data, _, _ = regionModel(c)
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
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/region/edit/"+lastID)
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
		viewPathRegion+"create.html",
		viewData,
	)
}

// Edit edit data
func (access *Region) Edit(c *gin.Context) {
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
				viewPathRegion+"edit.html",
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
func (access *Region) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var savePostError = make(map[string]string)
	regionUpdate, id, _ := regionModel(c)
	savePostError = regionUpdate.Validate()
	if len(savePostError) == 0 {
		_, saveErr := access.Iregion.Update(&regionUpdate)
		if saveErr != nil {
			savePostError = saveErr
		}
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/region/edit/"+id)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	flashMsg := stncsession.GetFlashMessage(c)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"err":          savePostError,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"post":         regionUpdate,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathRegion+"edit.html",
		viewData,
	)
}

/***  POST MODEL   ***/
func regionModel(c *gin.Context) (data entity.Region, idString string, err error) {

	id := c.PostForm("ID")

	idInt, _ := strconv.Atoi(id)

	var idN uint64

	idN = uint64(idInt)

	data.ID = idN
	data.UserID = stncsession.GetUserID2(c)
	data.Name = c.PostForm("Name")
	return data, id, nil
}
