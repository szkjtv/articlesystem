package controllers

import (
	"articlesystem/databases"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := databases.Dbcoom()
	ari := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Article WHERE id= ?")

	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(ari)
	// log.Println("delete 成功")
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)
}
