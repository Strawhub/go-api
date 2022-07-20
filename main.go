package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/index.tmpl")
	if err == nil {
		log.Println("Start Server")
	}
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "トップページ")
}
func database(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/db.tmpl")
	if err == nil {
		log.Println("Access")
	}
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "入力")
}
func result(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/result.tmpl")
	if err == nil {
		log.Println("Access")
	}
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "入力確認画面")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/data", database)
	http.HandleFunc("/data/result", result)
	http.ListenAndServe(":8080", nil)
}
