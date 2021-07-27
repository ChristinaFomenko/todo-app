package todo_app

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
