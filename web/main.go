package main

// 12:48

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	models "github.com/tonyguesswho/upld/pkg"
	"github.com/tonyguesswho/upld/pkg/mysql"

	"github.com/joho/godotenv"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	uploads  *mysql.UploadModel
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	dbURL, exists := os.LookupEnv("DBURL")
	var dsn string
	if exists {
		dsn = fmt.Sprintf("%s?parseTime=true", dbURL)
	} else {
		log.Fatal("No database URL")
	}

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(models.CreateTable)
	if err != nil {
		errorLog.Fatal(err)
		return
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		uploads:  &mysql.UploadModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
