package routers

import (
	"articlesystem/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/add", controllers.Addcontent)
	http.HandleFunc("/addshow", controllers.Arshow)
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/content", controllers.ShowContent)
	http.HandleFunc("/editor", controllers.Editor)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/admin", controllers.Admin)
	http.HandleFunc("/userinfo", controllers.Userinfo)
	http.ListenAndServe(":3006", nil)
}
