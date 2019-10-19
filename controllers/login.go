package controllers

import (
	"articlesystem/databases"
	"articlesystem/session"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		loginHTML, err := ioutil.ReadFile("view/login.html")
		checkError(err)
		w.Write(loginHTML)
		return
	}

	account := r.FormValue("account")
	password := r.FormValue("password")

	if isEmpty(account, password) {
		message(w, r, "账号或密码不能为空")
		return
	}

	user := databases.FindUserByUsernameAndPassword(account, password)
	if user == nil {
		message(w, r, "登录失败")
		return
	}

	// 登陆成功
	sess := session.GetSession(w, r)
	sess.SetAttr("user", user)
	// get, _ := sess.GetAttr("user")
	http.Redirect(w, r, "/admin", 302)

	getuser, _ := sess.GetAttr("user")
	if getuser == nil {
		log.Println("登录有效")
		http.Redirect(w, r, "login", 302)
	}

}

func Logout(w http.ResponseWriter, r *http.Request) {
	sess := session.GetSession(w, r)
	sess.DelAttr("user")
	http.Redirect(w, r, "/", 302)
}
