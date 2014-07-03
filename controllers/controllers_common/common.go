
package controllers_common

import (
	"github.com/thomassilvi/Meuh/models"
	"github.com/thomassilvi/Meuh/storage"
	"github.com/thomassilvi/Meuh/configuration"
        "html/template"
        "log"
        "net/http"
)

var Templates  *template.Template

//-------------------------------------------------------------------------------------------------

func GetUserPage (title string, req *http.Request) *models.UserPage {
	result := new(models.UserPage)
	result.Title = title

	u := GetUserBySession(req)

	if u == nil {
		result.User.Id = ""
		return result
	}

	result.User = *u
	result.IsLogged = (result.User.Id != "")

	return result
}

//-------------------------------------------------------------------------------------------------

func GetUserBySession(req *http.Request) *models.BasicUser {
        sessionFS, _ := storage.StoreFS.Get(req, "sessionFS")
        UserId, isOk := sessionFS.Values["UserId"].(string)

        if !isOk || UserId == "" {
                return nil
        }

        u := new(models.BasicUser)
        u.Id = UserId
        err := u.GetById()
        if err != nil {
                return nil
                log.Println("ERROR:GetUserBySession:", err.Error())
        }
        return u
}

//-------------------------------------------------------------------------------------------------

func StaticPage (title string , w http.ResponseWriter, req *http.Request) {
        p := GetUserPage(title,req)
        err := Templates.ExecuteTemplate(w, p.Title, *p)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
        }
}


//-------------------------------------------------------------------------------------------------

func AddTemplate (name string) {
	filename := configuration.AppConfiguration.Web.TemplatesPath + "/" + name
	tmp, err := Templates.ParseFiles(filename)

	if (err == nil) {
		Templates = tmp
	}
}

//-------------------------------------------------------------------------------------------------

