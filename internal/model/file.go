package model

import "time"

// File describes mysql tables files' row
type File struct {
	ID       int       `gorm:"primary_key" json:"id"`
	FileName string    `json:"file_name"`
	Ext      string    `json:"ext"`
	Path     string    `json:"path"`
	UserName string    `json:"user_name"`
	CreateON time.Time `json:"create_on"`
}
