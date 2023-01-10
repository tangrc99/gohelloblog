package model

import "time"

// Note describes mysql tables notes' row
type Note struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedOn time.Time `json:"created_on"`
	Content   string    `json:"content"`
}
