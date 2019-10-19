package controllers

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isEmpty(strs ...string) (isEmpty bool) {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" || len(str) == 0 {
			isEmpty = true
			return
		}
	}
	isEmpty = false
	return
}

func message(w http.ResponseWriter, r *http.Request, message string) {
	t, err := template.ParseFiles("view/massgage.html")
	checkError(err)

	err = t.Execute(w, map[string]string{"Message": message})
	checkError(err)
}

var t = template.Must(template.ParseGlob("view/*"))
