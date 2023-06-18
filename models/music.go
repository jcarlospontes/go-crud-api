package models

import (
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Musica string
	Autor  string
	Genero string
}
