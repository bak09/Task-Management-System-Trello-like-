package models

type Task struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Status      string `json:"status"`
	BoardID     int    `json:"board_id"`
	ColumnID    int    `json:"column_id"`
	UserID      int    `json:"user_id"`
}