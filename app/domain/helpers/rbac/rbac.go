package rbac

import (
	"fmt"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"stncCms/app/domain/helpers/stnccollection"
	"stncCms/app/domain/helpers/stncsession"
	repository "stncCms/app/domain/repository/cacheRepository"
)

// RbacCheck rbac kontrolu yapar
func RbacCheck(c *gin.Context, permissionName string) {
	stncsession.IsLoggedInRedirect(c)

	db := repository.DB
	appGrup := repository.UserRepositoryInit(db)
	userID := stncsession.GetUserID2(c)
	//fmt.Println(userID)

	userData, _ := appGrup.GetUser(userID)
	roleID := userData.RoleID
	//fmt.Println(roleID)
	found := RbacCheckRolePermission(roleID, permissionName)
	if found {
		viewDataPer := pongo2.Context{"title": "List"}
		c.HTML(
			http.StatusOK,
			"admin/roles/permission.html",
			viewDataPer,
		)
	}
	return
}

// RbacCheckRolePermission roleid ye gore izinleri getirir
func RbacCheckRolePermission(RoleID int, permissionName string) bool {

	db := repository.DB
	appGrup := repository.PermissionRepositoryInit(db)
	userPermissionData, _ := appGrup.GetUserPermission(RoleID)

	words := []string{}
	for _, row := range userPermissionData {
		fmt.Println(row.PermissionName)
		words = append(words, row.PermissionName)
	}
	fmt.Println(words)
	fmt.Println(userPermissionData)

	_, found := stnccollection.FindSlice(words, permissionName)
	if found {
		return true
	} else {
		return false
	}
}

/*
yapialcak islem make slice ile type component olanlar gelecek
bir sorgu lazimhg
*/

// RbacCheckForComponent Component lerer rbac kontrolu yapar
func RbacCheckForComponent(c *gin.Context, componentBaseName string) map[string]bool {

	db := repository.DB
	appGrup := repository.UserRepositoryInit(db)
	userID := stncsession.GetUserID2(c)

	userData, _ := appGrup.GetUser(userID)
	roleID := userData.RoleID

	permissionList := RbacCheckRolePermissionComponent(roleID, componentBaseName)
	return permissionList
}

// RbacCheckRolePermissionComponent roleid ye gore izinleri getirir
func RbacCheckRolePermissionComponent(RoleID int, componentBaseName string) map[string]bool {

	db := repository.DB
	appGrup := repository.PermissionRepositoryInit(db)
	userPermissionData, _ := appGrup.GetUserPermissionForComponent(RoleID, componentBaseName)

	permissionList := make(map[string]bool)
	for _, row := range userPermissionData {
		permissionList[row.Function] = true
	}
	return permissionList
}
