package models

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Id        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Descricao string `json:"descricao"`
	Url       string `json:"url"`
}

var Videos []Video
