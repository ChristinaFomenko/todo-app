package service

import (
	todo_app "github.com/KrisInferno/todo-app"
	"github.com/KrisInferno/todo-app/pkg/repository"
)

//объеденить структуру, добавить конструктор. аналогично сервису работы со списками

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, list repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(userId, itemId int, item todo_app.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, itemId)
	if err != nil {
		//list does not exists or does not belongs to user
		return 0, err
	}
	return s.repo.Create(itemId, 0, item)
}

func (s *TodoItemService) GetAll(userId, itemId int) ([]todo_app.TodoItem, error) {
	return s.repo.GetAll(userId, itemId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todo_app.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input todo_app.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
