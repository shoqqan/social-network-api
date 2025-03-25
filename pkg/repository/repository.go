package repository

import "github.com/jmoiron/sqlx"

type Authorization interface{}

type Posts interface{}

type Users interface{}

type Repository struct {
	Authorization
	Posts
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
