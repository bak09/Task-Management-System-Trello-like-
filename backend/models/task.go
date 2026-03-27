package models

type task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	BoardID     int    `json:"board_id"`
	ColumnID    int    `json:"column_id"`
	UserID      int    `json:"user_id"`
}
