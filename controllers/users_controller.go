package controllers

import (
	"github.com/thomassilvi/Meuh/controllers/controllers_common"
	"github.com/thomassilvi/Meuh/models"
	"github.com/thomassilvi/Meuh/storage"
	"github.com/thomassilvi/Meuh/tools"
	"github.com/thomassilvi/Meuh/configuration"
	"net/http"
	"log"
)

//-------------------------------------------------------------------------------------------------

func InitUserRoutes() {
        http.HandleFunc("/sign_in", SignInPage)
        http.HandleFunc("/sign_out", SignOutPage)

	if configuration.AppConfiguration.Web.DisableUserRegistration {
		http.HandleFunc("/user/new", UnavailablePage)
	} else {
		http.HandleFunc("/user/new", UserNewPage)
	}

        http.HandleFunc("/user/view", UserViewPage)
        http.HandleFunc("/user/edit", UserEditPage)
        http.HandleFunc("/user/change_password", UserChangeUserPasswordPage)
        http.HandleFunc("/user/forgotten_password", UserForgottenPassword)
}

//-------------------------------------------------------------------------------------------------

func InitUserTemplates() {
	controllers_common.AddTemplate ("sign_in.html")
	controllers_common.AddTemplate ("user_new.html")
	controllers_common.AddTemplate ("user_view.html")
	controllers_common.AddTemplate ("user_edit.html")
	controllers_common.AddTemplate ("user_change_password.html")
	controllers_common.AddTemplate ("user_forgotten_password.html")
	controllers_common.AddTemplate ("user_forgotten_password_2.html")

	controllers_common.AddTemplate ("unavailable.html")
	controllers_common.AddTemplate ("header_c.html")
	controllers_common.AddTemplate ("header_d.html")
	controllers_common.AddTemplate ("header.html")
	controllers_common.AddTemplate ("footer.html")
}

//-------------------------------------------------------------------------------------------------

func UnavailablePage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage ("unavailable", w, req)
}

//-------------------------------------------------------------------------------------------------

func SignOutPage(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:SignOutPage:" + req.Method)
	sessionFS, _ := storage.StoreFS.Get(req, "sessionFS")
	delete(sessionFS.Values, "UserId")
	delete(sessionFS.Values, "UserLogin")
	sessionFS.Save(req, w)
	http.Redirect(w, req, "/", http.StatusFound)
}

//-------------------------------------------------------------------------------------------------

func SignInPage(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:SignInPage:" + req.Method)
        
	p := controllers_common.GetUserPage("sign_in", req)
	if p.IsLogged {
		http.Redirect(w, req, "/error", http.StatusFound)
	}

	if req.Method == "GET" {
		controllers_common.Templates.ExecuteTemplate(w, "sign_in", *p)
	} else if req.Method == "POST" {
		login := req.FormValue("Login")
		password := req.FormValue("Password")
		u := models.BasicUser{ Login: login }
		err := u.Authentificate(password)
		if err != nil {
			log.Println("ERROR:invalid login or pwd : ", err.Error())
			p.AddError(err.Error())
			log.Println("DEBUG:SignInPage:page:", p)
			controllers_common.Templates.ExecuteTemplate(w, "sign_in", *p)
			return
		}
		sessionFS, _ := storage.StoreFS.Get(req, "sessionFS")
		UserId, ok := sessionFS.Values["UserId"].(string)

		if (!ok) || (UserId == "") {
			sessionFS.Values["UserId"] = u.Id
			sessionFS.Values["UserLogin"] = u.Login
			sessionFS.Save(req, w)
		}

		http.Redirect(w, req, "/", http.StatusFound)

	} else {
		log.Println("DEBUG:user_new:Not Handled:", req.Header)
	}
}

//-------------------------------------------------------------------------------------------------

func UserNewPage(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:UserNewPage:" + req.Method)

	p := controllers_common.GetUserPage("user_new", req)
	if p.IsLogged {
		http.Redirect(w, req, "/error", http.StatusFound)
	}

	if req.Method == "GET" {
		controllers_common.Templates.ExecuteTemplate(w, "user_new", *p)
	} else if req.Method == "POST" {
		var err error

		login := req.FormValue("Login")
		email := req.FormValue("Email")
		password := req.FormValue("Password")
		passwordConfirmation := req.FormValue("PasswordConfirmation")

		err = tools.CheckEmailValidity(email)
		if err != nil {
			p.AddError(err.Error())
		}
		err = tools.CheckPasswordValidity(password, passwordConfirmation)
		if err != nil {
			p.AddError(err.Error())
		}

		if len(p.Errors) > 0 {
			p.User.Login = login
			p.User.Email = email
			controllers_common.Templates.ExecuteTemplate(w, "user_new", *p)
			return
		}
		// create user

		u := models.BasicUser{Login: login, Email: email}
		err = u.Create(password)

		if err != nil {
			p.AddError(err.Error())
			p.User.Login = login
			p.User.Email = email
			controllers_common.Templates.ExecuteTemplate(w, "user_new", *p)
			return
		}

		log.Println("DEBUG:user_new:POST:create user:", u)

		// if ok

		sessionFS, _ := storage.StoreFS.Get(req, "sessionFS")
		sessionFS.Values["UserId"] = u.Id
		sessionFS.Values["UserLogin"] = u.Login
		sessionFS.Save(req, w)

		http.Redirect(w, req, "/", http.StatusFound)

	} else {
		log.Println("DEBUG:user_new:Not Handled:", req.Header)
	}
}

//-------------------------------------------------------------------------------------------------

func UserEditPage(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:UserEditPage:" + req.Method)

	p := controllers_common.GetUserPage("user_edit",req)

	if p.User.Id == "" {
		// TODO
		log.Println("ERROR:UserEditPage")
		http.Redirect(w, req, "/", http.StatusFound)
	}

	if req.Method == "GET" {
		controllers_common.Templates.ExecuteTemplate(w, "user_edit", *p)
	} else if req.Method == "POST" {
		var err error

		// check email
		email := req.FormValue("Email")
		err = tools.CheckEmailValidity(email)
		if err != nil {
			p.AddError(err.Error())
		}
		p.User.Email = email

		// end of checks

		if len(p.Errors) > 0 {
			controllers_common.Templates.ExecuteTemplate(w, "user_edit", *p)
			return
		}

		err = p.User.Save()
		if err != nil {
			p.AddError(err.Error())
			controllers_common.Templates.ExecuteTemplate(w, "user_edit", p)
			return
		}

		http.Redirect(w, req, "/user/view", http.StatusFound)
	}
}

//-------------------------------------------------------------------------------------------------

func UserChangeUserPasswordPage(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:UserChangeUserPasswordPage:" + req.Method)

	p := controllers_common.GetUserPage("user_change_password",req)
	
	if p.User.Id == "" {
		// TODO
		log.Println("ERROR:UserEditPage")
		http.Redirect(w, req, "/error", http.StatusFound)
	}

	if req.Method == "GET" {
		controllers_common.Templates.ExecuteTemplate(w, "user_change_password", *p)
	} else if req.Method == "POST" {
		var err error

		password := req.FormValue("NewPassword")
		passwordConfirmation := req.FormValue("NewPasswordConfirmation")

		err = tools.CheckPasswordValidity(password, passwordConfirmation)
		if err != nil {
			p.AddError(err.Error())
		}

		if len(p.Errors) > 0 {
			controllers_common.Templates.ExecuteTemplate(w, "user_change_password", *p)
			return
		}

		err = p.User.ChangePassword(password)

		if err != nil {
			p.AddError("An error occured while saving new passord. Operation cancelled.")
			controllers_common.Templates.ExecuteTemplate(w, "user_change_password", *p)
			return
		}

		http.Redirect(w, req, "/user/view", http.StatusFound)
	}

}

//-------------------------------------------------------------------------------------------------

func UserViewPage(w http.ResponseWriter, req *http.Request) {
	controllers_common.StaticPage ("user_view", w, req)
}

//-------------------------------------------------------------------------------------------------

func UserForgottenPassword(w http.ResponseWriter, req *http.Request) {
	log.Println("DEBUG:UserForgottenPassword:" + req.Method)

	p := new(models.UserPage)
	p.Init("forgotten_password")

	if req.Method == "GET" {
		controllers_common.Templates.ExecuteTemplate(w, "user_forgotten_password", *p)
	} else if req.Method == "POST" {
		login := req.FormValue("Login")
		u := new(models.BasicUser)
		u.Login = login
		err := u.GetByLogin()
		if err != nil {
			log.Println("ERROR:UserForgottenPassword", err)
			if err.Error() == "NoRows" {
				p.AddError("Ce login n'existe pas.")
				controllers_common.Templates.ExecuteTemplate(w, "user_forgotten_password", *p)
				return
			}

			http.Redirect(w, req, "/error", http.StatusFound)
			return
		}

		log.Println("TODO:UserForgottenPassword:send email to >" + u.Email + "<")

		p.User.Login = u.Login
		controllers_common.Templates.ExecuteTemplate(w, "user_forgotten_password_2", *p)
	}
}

//-------------------------------------------------------------------------------------------------
