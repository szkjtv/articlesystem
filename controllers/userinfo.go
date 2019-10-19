package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"articlesystem/session"
	"log"
	"net/http"
	"text/template"
)

func Userinfo(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	user, exist := sess.GetAttr("user")
	if !exist {
		http.Redirect(w, r, "/", 302)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles("view/adminindex.html")
		checkError(err)
		err = t.Execute(w, user)
		checkError(err)
		return
	}

	// POST 更新用户信息
	account := r.FormValue("account")
	password := r.FormValue("password")
	// email := r.FormValue("email")

	if isEmpty(account, password) {
		message(w, r, "字段不能为空")
		return
	}

	switch user := user.(type) {
	case *models.User:
		user.Account = account
		user.Password = password
		// user.Email = email
		databases.UpdateUser(user)
	default:
		log.Println(":userinfo:user.(type)", user)
	}
	http.Redirect(w, r, "/userinfo", 302)
}
