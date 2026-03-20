package service

import (
	todo "github.com/lockok/rest-api"
	"github.com/lockok/rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, ListId int) (todo.TodoList, error)
	Delete(userId, ListId int) error
	Update(userId, Listid int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
	}
}