package controllers

import (
	"articlesystem/databases"
	"articlesystem/models"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "index", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	//原始
	// user, _ := session.GetSession(w, r).GetAttr("user")

	// t, err := template.ParseFiles("view/index1.html")
	// t.ExecuteTemplate(w, "index1", nil) //The is new
	// checkError(err)

	// t.Execute(w, user)
	// checkError(err)

	//这是新的
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

	// fmt.Println("标题:", ar)
	// fmt.Println("内容:", ar)
	t.ExecuteTemplate(w, "index", ra)
	// fmt.Println("前端：", t)
	defer db.Close()

}
