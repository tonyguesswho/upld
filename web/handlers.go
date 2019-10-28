package main

import (
	"net/http"
	"os"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/upload.tmpl",
		"./ui/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	token := os.Getenv("TOKEN")
	err = ts.Execute(w, token)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) uploadFile(w http.ResponseWriter, r *http.Request) {

}

func (app *application) allUploads(w http.ResponseWriter, r *http.Request) {

}

func (app *application) showUpload(w http.ResponseWriter, r *http.Request) {

}
