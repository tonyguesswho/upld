package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	models "github.com/tonyguesswho/upld/pkg"
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

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		app.clientError(w, http.StatusForbidden, "Invalid request")
		return
	}
	auth := r.PostForm.Get("auth")
	if auth != os.Getenv("TOKEN") {
		app.clientError(w, http.StatusForbidden, "UnAuthorized")
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		app.clientError(w, http.StatusForbidden, "Invalid request")
		return
	}
	if header.Size > 1024*8 {
		app.clientError(w, http.StatusForbidden, "Maximum file size exceeded")
		return
	}
	defer file.Close()

	// read uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// check file type
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpeg" && fileType != "image/jpg" &&
		fileType != "image/gif" && fileType != "image/png" {
		app.clientError(w, http.StatusForbidden, "Invalid file type")
		return
	}
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		app.serverError(w, err)
		return
	}

	tempFile, err := ioutil.TempFile("./tmp", fmt.Sprintf("upload-*%s", fileEndings[0]))
	if err != nil {
		app.serverError(w, err)
	}
	defer tempFile.Close()
	fileName := filepath.Base(tempFile.Name())

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	fileSize := header.Size
	initialFileName := header.Filename

	id, err := app.uploads.Insert(fileSize, fileName, fileType, initialFileName)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/upload/%d", id), http.StatusSeeOther)
}

func (app *application) showUpload(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	u, err := app.uploads.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	files := []string{
		"./ui/show.tmpl",
		"./ui/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Upload: u}
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) allUploads(w http.ResponseWriter, r *http.Request) {
	u, err := app.uploads.All()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Uploads: u}
	files := []string{
		"./ui/uploads.tmpl",
		"./ui/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
