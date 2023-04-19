package model

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}
