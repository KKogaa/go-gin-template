package model

import "gorm.io/gorm"

type Manhwa struct {
	gorm.Model
	Title  string
	Author string
}

func (m *Manhwa) TableName() string {
	return "manhwas"
}
