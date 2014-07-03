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
