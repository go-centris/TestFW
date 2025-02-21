package cms_mod

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"stncCms/app/post/entity"
	languageEntity "stncCms/app/language/entity"
	"stncCms/pkg/helpers/lang"
	"stncCms/pkg/helpers/stnc2upload"
	stnccollection "stncCms/pkg/helpers/stnccollection"
	"stncCms/pkg/helpers/stncdatetime"
	Iauth "stncCms/app/auth/services"
	Icms "stncCms/app/post/services"
	ILanguage "stncCms/app/language/services"

	"stncCms/pkg/helpers/stnchelper"
	"stncCms/pkg/helpers/stncsession"
	"strconv"
	"strings"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2/v5"
	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	csrf "github.com/utrack/gin-csrf"
)

// Iregion "stncCms/app/services/regionServices_mod"
// Iauth "stncCms/app/services/authServices_mod"
// Icms "stncCms/app/services/cmsServices_mod"
// Icommon "stncCms/app/services/commonServices_mod"
// Isacrife "stncCms/app/services/sacrifeServices_mod"

// Post constructor
type Post struct {
	IPost     Icms.PostAppInterface
	ICatPost  Icms.CatPostAppInterface
	ICat      Icms.CatAppInterface
	IUser     Iauth.UserAppInterface
	ILanguage ILanguage.LanguageAppInterface
}

const viewPathPost = "admin/cms/post/"

func test(data string) string {
	return data
}

// InitPost post controller constructor
func InitPost(iPostApp Icms.PostAppInterface, ICatApp Icms.CatAppInterface, iCatPostApp Icms.CatPostAppInterface,
	iLangApp ILanguage.LanguageAppInterface, iUserApp Iauth.UserAppInterface) *Post {
	return &Post{
		IPost:     iPostApp,
		ICat:      ICatApp,
		ICatPost:  iCatPostApp,
		IUser:     iUserApp,
		ILanguage: iLangApp,
	}
}

// Index list
func (access *Post) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("post")
	var date stncdatetime.Inow
	var total int64
	access.IPost.Count(&total)
	postsPerPage := 10
	paginator := pagination.NewPaginator(c.Request, postsPerPage, total)
	offset := paginator.Offset()
	posts, _ := access.IPost.GetAllPagination(postsPerPage, offset)
	// var tarih stncdatetime.Inow
	// fmt.Println(tarih.TarihFullSQL("2020-05-21 05:08"))
	// fmt.Println(tarih.AylarListe("May"))
	// fmt.Println(tarih.Tarih())
	// //	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	//	tarih.FormatTarihForMysql("2020-05-17 05:08:40")
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"paginator":    paginator,
		"title":        locale.Get("Dashboard"),
		"flashMsg":     flashMsg,
		"posts":        posts,
		"date":         date,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}

	c.HTML(
		http.StatusOK,
		viewPathPost+"index.html",
		viewData,
	)
}

// Create all list f
func (access *Post) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("post")
	flashMsg := stncsession.GetFlashMessage(c)
	cats, _ := access.ICat.GetAll()
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "İçerik Ekleme",
		"flashMsg":     flashMsg,
		"catsData":     cats,
		"csrf":         csrf.GetToken(c),
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		viewPathPost+"create.html",
		viewData,
	)
}

// Store save method
func (access *Post) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("post")

	var post, _, _ = postModel(c)
	var savePostError = make(map[string]string)

	savePostError = post.Validate()

	sendFileName := "picture"
	filenameForm, _ := c.FormFile(sendFileName)
	stnc2upload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))

	/*//multiple
	form, _ := c.MultipartForm()
	files := form.File[sendFileName]
	stncupload.NewFileUpload().MultipleUploadFile(files, c.PostForm("Resim2"))
	*/

	// filename, uploadError := stncupload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))
	// if filename == "false" {
	// 	savePostError[sendFileName+"_error"] = uploadError
	// 	savePostError[sendFileName+"_valid"] = "is-invalid"
	// }

	// filename := "bos"
	fmt.Println(savePostError)
	catsPost := c.PostFormArray("cats")
	//fmt.Println(catsPost)
	catsData, _ := access.ICat.GetAll()
	// var list []entity.CategoriesSaveDTO
	fmt.Println(reflect.ValueOf(catsData).Kind())
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Name = row.Name
		//a, _ := strconv.Atoi(catsPost[key])
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPost, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}

	// for key, _ := range catsPost {
	// 	selectedID, _ := strconv.ParseUint(catsPost[key], 10, 64)
	// 	catsData[key].SelectedID = selectedID
	// }

	if len(savePostError) == 0 {
		post.Picture = "filename"
		saveData, saveErr := access.IPost.Save(&post)
		if saveErr != nil {
			savePostError = saveErr
		}

		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		var catPost = entity.CategoryPosts{}
		for _, row := range catsPost {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.CategoryID = catID
			catPost.PostID = saveData.ID
			saveCat, _ := access.ICatPost.Save(&catPost)
			catPost.ID = saveCat.ID + 1
		}
		lang := c.PostForm("languageSelect")
		var language = languageEntity.Languages{}
		language.PostID = saveData.ID
		language.Language = lang
		access.ILanguage.Save(&language)
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Kayıt başarı ile eklendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/post/edit/"+lastID)
		return
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "içerik ekleme",
		"catsPost":     catsPost,
		"catsData":     catsData,
		"csrf":         csrf.GetToken(c),
		"err":          savePostError,
		"post":         post,
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathPost+"create.html",
		viewData,
	)

}

// Edit edit data
func (access *Post) Edit(c *gin.Context) {
	//strconv.Atoi(c.Param("id"))
	//postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	locale, menuLanguage := lang.LoadLanguages("post")
	if postID, err := strconv.ParseUint(c.Param("postID"), 10, 64); err == nil {
		// Check if the article exists
		var catsPost []string

		catsPostData, _ := access.ICatPost.GetAllforPostID(postID)

		for _, row := range catsPostData {
			str := strconv.FormatUint(row.CategoryID, 10) //uint64 to stringS
			catsPost = append(catsPost, str)
		}

		catsData, _ := access.ICat.GetAll()
		for key, row := range catsData {
			catsData[key].ID = row.ID
			catsData[key].Name = row.Name
			//a, _ := strconv.Atoi(catsPost[key])
			finding := strconv.FormatInt(int64(row.ID), 10)
			_, found := stnccollection.FindSlice(catsPost, finding)
			if found {
				catsData[key].SelectedID = row.ID
			}
		}
		if posts, err := access.IPost.GetByID(postID); err == nil {
			modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

			viewData := pongo2.Context{
				"title":        "içerik ekleme",
				"catsPost":     catsPost,
				"catsData":     catsData,
				"post":         posts,
				"csrf":         csrf.GetToken(c),
				"flashMsg":     flashMsg,
				"locale":       locale,
				"localeMenus":  menuLanguage,
				"modulNameUrl": modulName,
			}
			c.HTML(
				http.StatusOK,
				viewPathPost+"edit.html",
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
func (access *Post) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	locale, menuLanguage := lang.LoadLanguages("post")
	var post, idN, id = postModel(c)

	var savePostError = make(map[string]string)

	savePostError = post.Validate()

	sendFileName := "picture"

	//	filename, uploadError := upload(c, sendFileName)
	//	filenameForm, _ := c.FormFile(sendFileName)
	// filename2, data2, data3 := c.Request.FormFile(sendFileName)

	//	filename, uploadError := stncupload.NewFileUpload().UploadFile(filenameForm)

	// form, _ := c.MultipartForm()
	// files := form.File[sendFileName]
	// stncupload.NewFileUpload().MultipleUploadFile(files, c.PostForm("Resim2"))

	filenameForm, _ := c.FormFile(sendFileName)
	filename, uploadError := stnc2upload.NewFileUpload().UploadFile(filenameForm, c.PostForm("Resim2"))

	if filename == "false" {
		savePostError[sendFileName+"_error"] = uploadError
		savePostError[sendFileName+"_valid"] = "is-invalid"
	}

	catsPostForm := c.PostFormArray("cats")

	access.ICatPost.DeleteForPostID(idN)

	catsData, _ := access.ICat.GetAll()

	// fmt.Println(reflect.ValueOf(catsData).Kind())
	for key, row := range catsData {
		catsData[key].ID = row.ID
		catsData[key].Name = row.Name
		finding := strconv.FormatInt(int64(row.ID), 10)
		_, found := stnccollection.FindSlice(catsPostForm, finding)
		if found {
			catsData[key].SelectedID = row.ID
		}
	}

	if len(savePostError) == 0 {
		post.Picture = filename
		saveData, saveErr := access.IPost.Update(&post)
		if saveErr != nil {
			savePostError = saveErr
		}

		var catPost = entity.CategoryPosts{}

		for _, row := range catsPostForm {
			catID, _ := strconv.ParseUint(row, 10, 64)
			catPost.CategoryID = catID
			catPost.PostID = saveData.ID
			saveCat, _ := access.ICatPost.Save(&catPost)
			catPost.ID = saveCat.ID + 1
		}
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Kayıt başarı ile düzenlendi", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/post/edit/"+id)
		return
	}
	modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))

	viewData := pongo2.Context{
		"title":        "içerik ekleme",
		"catsPost":     catsPostForm,
		"catsData":     catsData,
		"err":          savePostError,
		"csrf":         csrf.GetToken(c),
		"post":         post,
		"locale":       locale,
		"localeMenus":  menuLanguage,
		"modulNameUrl": modulName,
	}
	c.HTML(
		http.StatusOK,
		viewPathPost+"edit.html",
		viewData,
	)
}

// Delete data
func (access *Post) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if postID, err := strconv.ParseUint(c.Param("postID"), 10, 64); err == nil {
		access.IPost.Delete(postID)
		modulName := stnchelper.ModulNameUrlCheck(c.Param("ModulName"))
		stncsession.SetFlashMessage("Success Delete", "success", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/"+modulName+"/post/"+c.Param("postID"))
		return
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Delete data
func (access *Post) Upload(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	filenameForm, _ := c.FormFile("file")
	stnc2upload.NewFileUpload().UploadFile(filenameForm, "")
	c.JSON(http.StatusBadGateway, "hata")
	// c.JSON(http.StatusOK, "ok")
}

// form post model
func postModel(c *gin.Context) (post entity.Post, idD uint64, idStr string) {
	id := c.PostForm("ID")
	title := c.PostForm("PostTitle")
	content := c.PostForm("PostContent")
	excerpt := c.PostForm("PostExcerpt")
	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	//	var post = entity.Post{}
	post.ID = idN
	post.UserID = stncsession.GetUserID2(c)
	post.PostTitle = title
	post.PostSlug = stnchelper.Slugify(title, 15)
	post.PostType = 1
	post.PostStatus = 1
	post.PostContent = content
	post.PostExcerpt = excerpt
	post.MenuOrder = 1
	post.CommentCount = 1
	return post, idN, id
}

/*
kullanımı
	sendFileName := "picture"
	filename, uploadError := upload(c, sendFileName)
*/
//buradaki sıkıntı edit sırasında resimde bir işlem yapmazsan veritababından resimi siliyor
//TODO: boyutlandırma https://github.com/disintegration/imaging
func uploadIPTAL(c *gin.Context, formFilename string) (string, string) {

	var uploadFilePath string = "public/upl/"
	var filename string
	var errorReturn string

	file, header, err := c.Request.FormFile(formFilename)
	//fmt.Println(file)
	//fmt.Println(header)

	if header != nil {
		if err != nil {
			// c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			errorReturn = err.Error()
		}

		size := header.Size
		var size2 = strconv.FormatUint(uint64(size), 10)
		if size > int64(1024000*5) { // 1 MB
			// return "", errors.New("sorry, please upload an Image of 500KB or less")
			errorReturn = "Resim boyutu çok yüksek maximum 5 MB olmalıdır" + size2
			filename = "false"
		}

		filenameOrg := header.Filename

		filenameExtension := filepath.Ext(filenameOrg)

		realFilename := strings.Split(filenameOrg, ".")

		realFilenameSlug := stnchelper.GenericName(realFilename[0], 50)

		filename = realFilenameSlug + filenameExtension

		out, err := os.Create(uploadFilePath + filename)
		if err != nil {
			log.Fatal(err)
			errorReturn = err.Error()
			filename = "false"
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
			errorReturn = err.Error()
			filename = "false"
		}

		buf, _ := os.ReadFile(uploadFilePath + filename)

		if filetype.IsImage([]byte(buf)) {
			filename = realFilenameSlug + filenameExtension
			//TODO: resim boyutlandırma gelecek
			//https://github.com/disintegration/imaging
		} else {
			path := uploadFilePath + filename
			err := os.Remove(path)
			if err != nil {
				errorReturn = err.Error()
			}
			errorReturn = filenameOrg + " gerçek bir resim dosyası değildir"
			filename = "false"
		}

		return filename, errorReturn
	} else {
		return "", ""
	}
}
