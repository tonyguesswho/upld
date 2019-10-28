package main

import (
	"github.com/bmizerany/pat"
	"net/http"

)

func (app *application) routes() http.Handler{
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Post("/upload", http.HandlerFunc(app.uploadFile))
	mux.Get("/uploads", http.HandlerFunc(app.allUploads))
	mux.Get("/upload/:id", http.HandlerFunc(app.showUpload))

	fileServer := http.FileServer(http.Dir("./tmp"))
	mux.Get("/tmp/", http.StripPrefix("/tmp", fileServer))

	return mux
}