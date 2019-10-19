package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"net/http"
)

//查看详情
func ShowContent(w http.ResponseWriter, r *http.Request) {
	db := databases.Dbcoom()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Article WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := models.Article{}
	for selDB.Next() {
		var id int
		var title, content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Title = title
		emp.Content = content
	}
	t.ExecuteTemplate(w, "contentt", emp)
	defer db.Close()

}
