package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"articlesystem/session"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	user, _ := session.GetSession(w, r).GetAttr("user")
	t.ExecuteTemplate(w, "adminindex", user)

	db := databases.Dbcoom()
	selDB, err := db.Query("SELECT * FROM Article ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	ar := models.Article{}
	ra := []models.Article{}
	for selDB.Next() {
		var id int
		var title, content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		ar.Id = id
		ar.Title = title
		ar.Content = content
		ra = append(ra, ar)

	}
	t.ExecuteTemplate(w, "adminindex", ra)
	defer db.Close()

}

func Useradmin(w http.ResponseWriter, r *http.Request) {
	user, _ := session.GetSession(w, r).GetAttr("user")

	// t, err := template.ParseFiles("view/adminindex.html")
	t.ExecuteTemplate(w, "adminindex", user)
	// checkError(err)

	// err = t.Execute(w, user)
	// checkError(err)
}
