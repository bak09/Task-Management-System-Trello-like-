package models

type Board struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
}