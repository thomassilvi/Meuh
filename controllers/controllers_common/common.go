/*
Meuh
Copyright (C) 2014 Thomas Silvi

This file is part of Meuh.

GoSimpleConfigLib is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

GoSimpleConfigLib is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar. If not, see <http://www.gnu.org/licenses/>.
*/

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

