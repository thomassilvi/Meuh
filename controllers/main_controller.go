package controllers

import (
	"github.com/thomassilvi/Meuh/controllers_common"
	"net/http"
)

//-------------------------------------------------------------------------------------------------

func InitMainTemplates() {
	controllers_common.AddTemplate("header_c.html")
	controllers_common.AddTemplate("header_d.html")
	controllers_common.AddTemplate("header.html")
	controllers_common.AddTemplate("footer.html")
	controllers_common.AddTemplate("/index.html")
	controllers_common.AddTemplate("/welcome.html")
	controllers_common.AddTemplate("/home.html")
	controllers_common.AddTemplate("/about.html")
	controllers_common.AddTemplate("/error.html")
	controllers_common.AddTemplate("/terms.html")
	controllers_common.AddTemplate("/privacy.html")
	controllers_common.AddTemplate("/contacts.html")
}

//-------------------------------------------------------------------------------------------------

func ErrorPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("error",w,req)
}

//-------------------------------------------------------------------------------------------------

func AboutPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("about",w,req)
}

//-------------------------------------------------------------------------------------------------

func RootPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("index",w,req)
}

//-------------------------------------------------------------------------------------------------

func TermsPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("terms",w,req)
}

//-------------------------------------------------------------------------------------------------

func PrivacyPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("privacy",w,req)
}

//-------------------------------------------------------------------------------------------------

func ContactPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage("contacts",w,req)
}

//-------------------------------------------------------------------------------------------------





