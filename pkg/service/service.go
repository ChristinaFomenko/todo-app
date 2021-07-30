package service

import (
	todo_app "github.com/KrisInferno/todo-app"
	"github.com/KrisInferno/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user interface{}) (int error, err error)
	GenerateToken(username, password int) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_app.TodoList) (int, error)
	GetAll(userId int) ([]todo_app.TodoList, error)
	GetById(userId, listId int) (todo_app.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type TodoItem interface {
	Create(userId int, id int, list todo_app.TodoItem) (int, error)
	GetAll(usedId, listId int) ([]todo_app.TodoItem, error)
	GetById(userId, itemId int) (todo_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
