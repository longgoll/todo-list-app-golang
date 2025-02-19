package models

import "time"

type TodoItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:Doing"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Add this function to specify the table name
func (TodoItem) TableName() string {
	return "todos"
}
