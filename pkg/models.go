package models

import (
	"errors"
	"time"
)

// ErrNoRecord : No record found
var ErrNoRecord = errors.New("models: no matching record found")

// Upload :
type Upload struct {
	ID              int
	Filename        string
	Filesize        string
	Filetype        string
	Initialfilename string
	Created         time.Time
}

// CreateTable : create uploads table
var CreateTable = `CREATE TABLE IF NOT EXISTS uploads (    
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,    
	filename VARCHAR(100) NOT NULL,    
	filesize INT NOT NULL,   
	filetype VARCHAR(100) NOT NULL,
	initialfilename VARCHAR(100) NOT NULL, 
	created DATETIME NOT NULL )`
