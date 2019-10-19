package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"net/http"
)

//更新内容
func Editor(w http.ResponseWriter, r *http.Request) {
	db := databases.Dbcoom()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM	Article WHERE id=? ", nId)
	if err != nil {
		panic(err.Error())
	}
	ar := models.Article{}
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
	}
	t.ExecuteTemplate(w, "editor", ar)
	defer db.Close()
}

//更新
func Update(w http.ResponseWriter, r *http.Request) {
	db := databases.Dbcoom()
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Article SET title=?, content=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, content, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)
}
