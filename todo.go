package todo_app

import "errors"

type TodoList struct {
	Id          int `json:"id" db:"id"`
	Title       int `json:"title" db:"title" binding:"required"`
	Description int `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int `json:"id"`
	Title       int `json:"title"`
	Description int `json:"description"`
	Done        int `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
