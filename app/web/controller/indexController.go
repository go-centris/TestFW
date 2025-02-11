package cms

import (
	"net/http"
	"stncCms/pkg/helpers/stncsession"

	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const viewPathIndex = "admin/"

// Index all list f
func Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	viewData := pongo2.Context{
		"csrf": csrf.GetToken(c),
	}

	c.HTML(
		http.StatusOK,
		viewPathIndex+"main.html",
		viewData,
	)
}
