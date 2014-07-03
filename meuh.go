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

package meuh

import (
        "github.com/thomassilvi/Meuh/configuration"
        "github.com/thomassilvi/Meuh/controllers"
        "github.com/thomassilvi/Meuh/storage"
	"log"
	"net/http"
	"strconv"
)

//-------------------------------------------------------------------------------------------------

func Init(configfilename string) {
	var err error

	err = configuration.InitConfiguration(configfilename)
	if err != nil {
		log.Fatal(err)
	}
	err = storage.InitStorage()
	if err != nil {
		log.Fatal(err)
	}
}

//-------------------------------------------------------------------------------------------------

func InitDefault() {
	http.HandleFunc("/", controllers.RootPage)

	assetsTmp := http.Dir(configuration.AppConfiguration.Web.PublicAssetsPath)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assetsTmp)))

	http.HandleFunc("/error", controllers.ErrorPage)
	http.HandleFunc("/about", controllers.AboutPage)
	http.HandleFunc("/terms", controllers.TermsPage)
	http.HandleFunc("/privacy", controllers.PrivacyPage)
	http.HandleFunc("/contacts", controllers.ContactPage)

}

//-------------------------------------------------------------------------------------------------

func Run() {
	port := ":" + strconv.FormatUint(uint64(configuration.AppConfiguration.Web.Port), 10)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//-------------------------------------------------------------------------------------------------
