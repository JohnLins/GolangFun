package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	//http.HandleFunc("/process", processor)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	//landing
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
}

/*func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}*/

	fname := r.FormValue("firstn")
	lname := r.FormValue("lastn")
	color := r.FormValue("color")
	fontSize := r.FormValue("fontSize")

	d := struct {
		First string
		Last string
		Color string
		Size string
	}{
		First: fname,
		Last: lname,
		Color: color,
		Size: fontSize,
	}

	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}