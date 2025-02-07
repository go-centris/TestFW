package reportSacrife_mod

import (
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"net/http"
	"stncCms/app/domain/helpers/stnchelper"
	"stncCms/app/domain/helpers/stncsession"
	IreportSacrife "stncCms/app/services/reportSacrifeServices_mod"
)

// Kisiler constructor
type Report struct {
	IReport IreportSacrife.ReportAppInterface
}

const viewPathReport = "admin/report/"

// InitGKisiler post controller constructor
func InitReport(repApp IreportSacrife.ReportAppInterface) *Report {
	return &Report{
		IReport: repApp,
	}
}

// GetAllUsersWhoAddedMostUser  En cok kurban ekleyen hocamiz
func (access *Report) GetAllUsersWhoAddedMostUser(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var total int64
	access.IReport.GetAllUsersWhoAddedMostSacrifeAndUserCount(&total)
	postsPerPage := 15
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.IReport.GetAllUsersWhoAddedMostSacrifeAndUser(postsPerPage, offset)
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
		viewPathReport+"getAllUsersWhoAddedMostUsers.html",
		viewData,
	)
}

// GetAllUsersWhoAddedMostBranch  En cok kurban ekleyen subemiz
func (access *Report) GetAllUsersWhoAddedMostBranch(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var total int64
	access.IReport.GetAllUsersWhoAddedMostSacrifeAndBranchCount(&total)
	postsPerPage := 15
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.IReport.GetAllUsersWhoAddedMostSacrifeAndBranch(postsPerPage, offset)
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
		viewPathReport+"getAllUsersWhoAddedMostBranch.html",
		viewData,
	)
}

// GetAllCharitableWhoAddedMostSacrife  En cok kurban ekleyen hayirsever
func (access *Report) GetAllCharitableWhoAddedMostSacrife(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	var total int64
	access.IReport.GetAllCharitableWhoAddedMostSacrifeCount(&total)
	postsPerPage := 15
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.IReport.GetAllCharitableWhoAddedMostSacrife(postsPerPage, offset)
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
		viewPathReport+"getAllCharitableWhoAddedMostSacrife.html",
		viewData,
	)
}
