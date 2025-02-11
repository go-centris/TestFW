package region_mod

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/leonelquinteros/gotext"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/pkg/helpers/stnccollection"
	"stncCms/pkg/helpers/stnchelper"
	"stncCms/pkg/helpers/stncsession"
	Iregion "stncCms/app/services/regionServices_mod"

	"strconv"
)

// Branch constructor
type Branch struct {
	IBranch Iregion.BranchAppInterface
	IRegion Iregion.RegionAppInterface
}

const viewPathBranch = "admin/region/branch/"

func InitBranch(iBranchApp Iregion.BranchAppInterface, iRegionApp Iregion.RegionAppInterface) *Branch {
	return &Branch{
		IBranch: iBranchApp,
		IRegion: iRegionApp,
	}
}

// Index  list
func (access *Branch) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var total int64
	access.IBranch.GetAllPaginateCount(&total)
	postsPerPage := 15
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.IBranch.GetAllPaginate(postsPerPage, offset)
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
		viewPathBranch+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Branch) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var getText *gotext.Locale
	getText = gotext.NewLocale("public/locales", "tr_TR")
	getText.AddDomain("default")
	region, _ := access.IRegion.GetAll()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"locale":       getText,
		"regions":      region,
		"csrf":         csrf.GetToken(c),
		"flashMsg":     flashMsg,
		"kisiAdSoyad":  "",
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathBranch+"create.html",
		viewData,
	)
}

// Store save method
func (access *Branch) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	var data, _, _ = BranchModel(c)
	var savePostError = make(map[string]string)
	savePostError = data.Validate()
	region, _ := access.IRegion.GetAll()

	if len(savePostError) == 0 {
		saveData, saveErr := access.IBranch.Save(&data)
		if saveErr != nil {
			savePostError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)

		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/branch/edit/"+lastID)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	viewData := pongo2.Context{
		"csrf":         csrf.GetToken(c),
		"err":          savePostError,
		"regions":      region,
		"post":         data,
		"flashMsg":     flashMsg,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathBranch+"create.html",
		viewData,
	)
}

// Edit edit data
func (access *Branch) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		if posts, err := access.IBranch.GetByID(ID); err == nil {
			region, _ := access.IRegion.GetAll()
			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"post":         posts,
				"regions":      region,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"modulNameUrl": modulName,
			}

			c.HTML(
				http.StatusOK,
				viewPathBranch+"edit.html",
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
func (access *Branch) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var savePostError = make(map[string]string)
	BranchUpdate, id, _ := BranchModel(c)
	savePostError = BranchUpdate.Validate()
	region, _ := access.IRegion.GetAll()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	if len(savePostError) == 0 {
		_, saveErr := access.IBranch.Update(&BranchUpdate)
		if saveErr != nil {
			savePostError = saveErr
		}
		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/branch/edit/"+id)
		return
	} else {
		stncsession.SetFlashMessage("Zorunlu alanları lütfen doldurunuz", "danger", c)
	}
	flashMsg := stncsession.GetFlashMessage(c)

	viewData := pongo2.Context{
		"err":          savePostError,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"post":         BranchUpdate,
		"regions":      region,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathBranch+"edit.html",
		viewData,
	)
}

// ajax event
func (access *Branch) GetBranchListForRegion(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	// flashMsg := stncsession.GetFlashMessage(c)
	if regionID, err := strconv.ParseUint(c.Param("regionID"), 10, 64); err == nil {
		if jsonData, err := access.IBranch.GetByRegionID(regionID); err == nil {
			viewData := pongo2.Context{
				"csrf":     csrf.GetToken(c),
				"jsonData": jsonData,
				// "flashMsg": flashMsg,
			}
			c.JSON(http.StatusOK, viewData)
			return
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

/***  POST MODEL   ***/
func BranchModel(c *gin.Context) (data entity.Branches, idString string, err error) {
	id := c.PostForm("ID")
	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	data.ID = idN
	data.UserID = stncsession.GetUserID2(c)
	data.RegionID = stnccollection.StringToint(c.PostForm("RegionID"))
	data.Title = c.PostForm("Title")
	data.BranchCode = c.PostForm("BranchCode")
	data.ManagerName = c.PostForm("ManagerName")
	data.ManagerPhone = c.PostForm("ManagerPhone")
	data.ManagerMail = c.PostForm("ManagerMail")

	return data, id, nil
}
