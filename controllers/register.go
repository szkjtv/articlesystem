package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"io/ioutil"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		registerHTML, err := ioutil.ReadFile("./view/register.html")
		checkError(err)
		w.Write(registerHTML)
		return
	}

	account := r.FormValue("account")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	// email := r.FormValue("email")

	if isEmpty(account, password, password2) {
		message(w, r, "账号或密码不能为空")
		return
	}
	if password != password2 {
		message(w, r, "两次密码不一样")
		return
	}

	user := &models.User{
		Account:  account,
		Password: password,
		// Email:    email,
	}
	databases.AddUser(user)
	message(w, r, "注册成功")
}
