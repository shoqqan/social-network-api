package network

import "time"

type Post struct {
	Id        string    `json:"-"`
	UserId    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
