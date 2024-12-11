package domain

import "time"

type TodoList struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileID      string    `json:"file_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
