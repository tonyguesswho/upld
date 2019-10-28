package main

import (
	models "github.com/tonyguesswho/upld/pkg"
)

type templateData struct {
	Upload  *models.Upload
	Uploads []*models.Upload
	Token   string
}
