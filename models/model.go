package models

import "gorm.io/gorm"

type UrlModel struct {
	gorm.Model
	Id  int64 `json:"-" gorm:"-;primary_key;AUTO_INCREMENT"`
	Url string `json:"url" binding:"required,url" gorm:"type:text;not null"`
}

type RedirectRequest struct {
	Code string `uri:"code" binding:"required"`
}
