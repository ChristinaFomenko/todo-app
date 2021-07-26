package service

import "github.com/KrisInferno/todo-app/pkg/repository"

type Authorization interface {
	CreateUser(user interface{}) (int error, err error)
	GenerateToken(username, password string) (string, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
