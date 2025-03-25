package network

import "time"

type User struct {
	Id        string    `json:"-"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Posts     []Post    `json:"posts"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
