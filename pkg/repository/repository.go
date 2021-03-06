package repository

import (
	todo_app "github.com/KrisInferno/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GetUser(username string, password int) (todo_app.User, error)
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
	GetAll(userId int, itemId int) ([]todo_app.TodoItem, error)
	GetById(userId, itemId int) (todo_app.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo_app.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
