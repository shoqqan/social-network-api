package repository

type Authorization interface{}

type Posts interface{}

type Users interface{}

type Repository struct {
	Authorization
	Posts
	Users
}

func NewRepository() *Repository {
	return &Repository{}
}
