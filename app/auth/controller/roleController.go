package auth_mod

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"stncCms/pkg/helpers/lang"
	"stncCms/pkg/helpers/stnccollection"
	"stncCms/pkg/helpers/stncdatetime"
	"stncCms/pkg/helpers/stnchelper"
	"strconv"

	"stncCms/pkg/helpers/stncsession"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	Iauth "stncCms/app/auth/services"
	Icommon "stncCms/app/services/commonServices_mod"
		  dtoAuth "stncCms/app/auth/dto"
			 modulesEntity "stncCms/app/modules/entity"
			 modulesDTO "stncCms/app/modules/dto"
			
			 
authEntity "stncCms/app/auth/entity"
)

// Permission constructor
type Roles struct {
	IPermission     Iauth.PermissionAppInterface
	IModules        Icommon.ModulesAppInterface
	IRole           Iauth.RoleAppInterface
	IrolePermission Iauth.RolePermissionAppInterface
}

const viewPathPermission = "admin/auth/roles/"

func InitRoles(iPermissionApp Iauth.PermissionAppInterface, iModulesApp Icommon.ModulesAppInterface, iRolesApp Iauth.RoleAppInterface, iRolePermApp Iauth.RolePermissionAppInterface) *Roles {
	return &Roles{
		IPermission:     iPermissionApp,
		IModules:        iModulesApp,
		IRole:           iRolesApp,
		IrolePermission: iRolePermApp,
	}
}

// Index list
func (access *Roles) Index(c *gin.Context) {

	stncsession.IsLoggedInRedirect(c)
	//rbac.RbacCheck(c, "post-index")
	locale, menuLanguage := lang.LoadLanguages("roles")
	flashMsg := stncsession.GetFlashMessage(c)
	var date stncdatetime.Inow
	var total int64
	access.IRole.Count(&total)
	postsPerPage := 3
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	data, _ := access.IRole.GetAllPagination(postsPerPage, offset)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "List",
		"dataList":     data,
		"date":         date,
		"csrf":         csrf.GetToken(c),
		"flashMsg":     flashMsg,
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathPermission+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Roles) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("roles")
	var data []modulesEntity.ModulesAndPermissionDTO
	data, _ = access.IModules.GetAllModulesMerge()
	for num, v := range data {
		var list = []authEntity.Permission{}
		list, _ = access.IPermission.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}

	// //#json formatter #stncjson
	// empJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "permissions",
		"datas":        data,
		"flashMsg":     flashMsg,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathPermission+"create.html",
		viewData,
	)
}

// store data
func (access *Roles) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	//once bi yere rolu kaydet
	//sonra kaydeidlern id yi alman lazim

	roleData := authEntity.Role{
		Title: c.PostForm("Title"),
		Slug:  c.PostForm("Title"),
	}

	saveRoleData, _ := access.IRole.Save(&roleData)
	roleID := saveRoleData.ID
	var data []modulesEntity.ModulesAndPermissionDTO
	data, _ = access.IModules.GetAllModulesMerge()
	for num, v := range data {
		var list = []authEntity.Permission{}
		list, _ = access.IPermission.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}

	for _, v := range data {
		for _, per := range v.Permissions {
			rolePermissondata := authEntity.RolePermisson{
				RoleID:       roleID,
				PermissionID: per.ID,
				Active:       0,
			}
			access.IrolePermission.Save(&rolePermissondata)
		}
	}

	names, _ := c.Request.PostForm["grant-caps[]"]
	for _, row := range names {
		grandPermissionID := stnccollection.StringToint(row)
		access.IrolePermission.UpdateActiveStatus(roleID, grandPermissionID, 1)
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin"+modulName+"/roles/edit/"+stnccollection.IntToString(roleID))
	return

	// grandPermissionID, e1 := strconv.Atoi(row)
	// fmt.Println(e1)
	// if e1 == nil {
	// 	fmt.Println("\nAfter:")
	// 	fmt.Printf("Type: %T ", grandPermissionID)
	// 	fmt.Printf("\nValue: %v", grandPermissionID)
	// }
	///*******************
	// names := c.PostFormMap("grant-caps")
	// for _, row := range names {
	// 	fmt.Println(row)
	// }
	// names := c.QueryMap("grant-caps[]")

	// names, _ := c.Request.PostForm["grant-caps[]"]
	// fmt.Println("grand list : ", names)
	// for _, row := range names {
	// 	fmt.Println("granst: ", row)

	// }

}

func (access *Roles) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("roles")
	flashMsg := stncsession.GetFlashMessage(c)
	var data []modulesDTO.ModulesAndPermissionRoleDTO
	if roleID, err := strconv.Atoi(c.Param("ID")); err == nil {
		data, _ = access.IModules.GetAllModulesMergePermission()
		roleData, _ := access.IRole.GetByID(roleID) //TODO: bu veri  access.IRole.EditList iicne de geliyor orada mi almak mantakli ??
		for num, v := range data {
			var list = []dtoAuth.RoleEditList{}
			list, _ = access.IRole.EditList(v.ID, roleID)
			// fmt.Println(v.ModulName)
			data[num].RoleEditList = list
		}
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

		viewData := pongo2.Context{
			"title":        "permissions",
			"roleData":     roleData,
			"datas":        data,
			"roleID":       roleID,
			"flashMsg":     flashMsg,
			"csrf":         csrf.GetToken(c),
			"locale":       locale,
			"localeMenus":  menuLanguage,
			"modulNameUrl": modulName,
		}

		c.HTML(
			http.StatusOK,
			viewPathPermission+"edit.html",
			viewData,
		)
	}

}

func (access *Roles) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	roleID := c.PostForm("roleID")
	roleIDint := stnccollection.StringToint(c.PostForm("roleID"))
	title := c.PostForm("Title")

	access.IRole.UpdateTitle(roleIDint, title)

	//TODO: neden calismadi
	//titleSlug := stnchelper.Slugify(title, 15)
	//access.IRole.Update(
	//	&entity.Role{
	//		ID:      roleIDint,
	//		Title:   title,
	//		Slug:    titleSlug,
	//		Context: titleSlug,
	//		Status:  1,
	//	})

	grants, _ := c.Request.PostForm["grant-caps[]"]
	for _, row := range grants {
		grandPermissionID := stnccollection.StringToint(row)
		access.IrolePermission.UpdateActiveStatus(roleIDint, grandPermissionID, 1)
	}

	deny, _ := c.Request.PostForm["deny-caps[]"]
	for _, row := range deny {
		grandPermissionID := stnccollection.StringToint(row)
		access.IrolePermission.UpdateActiveStatus(roleIDint, grandPermissionID, 0)
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
	c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"roles/edit/"+roleID)
	return
}

func (access *Roles) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		access.IRole.Delete(ID)
		stncsession.SetFlashMessage("delete", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"roles")
		return
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (access *Roles) IndexKnockout(c *gin.Context) {
	// allpermission, err := access.IPermission.GetAllPaginationermission()
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("roles")
	var date stncdatetime.Inow
	var data []modulesDTO.ModulesAndPermissionDTO
	data, _ = access.IModules.GetAllModulesMerge()
	for num, v := range data {
		var list = []authEntity.Permission{}
		list, _ = access.IPermission.GetAllPaginationermissionForModulID(int(v.ID))
		data[num].Permissions = list
	}
	// //#json formatter #stncjson https://github.com/TylerBrock/colorjson
	empJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("MarshalIndent funnction output\n %s\n", string(empJSON))
	// js, _ := json.Marshal(data)
	// fmt.Println((jsonData))
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        "permissions",
		"datas":        data,
		"json":         string(empJSON),
		"date":         date,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathPermission+"knockout.html",
		viewData,
	)
}
