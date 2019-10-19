package controllers

import (
	"articlesystem/databases"
	"net/http"
)

func Arshow(w http.ResponseWriter, r *http.Request) {
	t.ExecuteTemplate(w, "addcontent", nil)
}

func Addcontent(w http.ResponseWriter, r *http.Request) {

	// if r.Method == "GET" {
	// 	addHTML, err := ioutil.ReadFile("./view/add.html")
	// 	checkError(err)
	// 	w.Write(addHTML)
	// 	return
	// }

	// title := r.FormValue("title")
	// content := r.FormValue("content")

	// if isEmpty(title, content) {
	// 	message(w, r, "标题或文章不为能空")
	// 	return
	// }

	// content1 := &models.Content{
	// 	Title:   title,
	// 	Content: content,
	// }
	// databases.AddContent(content1)
	// message(w, r, "发布成功")

	db := databases.Dbcoom()
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		insForm, err := db.Prepare("INSERT INTO Article(title,content)VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(title, content)
		http.Redirect(w, r, "/admin", 302)
	}

}
