package models

type Column struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"not null"`
	BoardID int    `json:"board_id"`
}