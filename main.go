package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.tmpl")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "トップページ")
}
func database(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/db.tmpl")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "入力")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/data", database)
	http.ListenAndServe(":8080", nil)
}
