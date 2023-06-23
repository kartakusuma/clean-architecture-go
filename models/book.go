package models

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(100)"`
	Author    string    `json:"author" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
