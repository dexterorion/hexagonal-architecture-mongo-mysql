package todo

import "github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/domain"

type ToDo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ToDoList []ToDo

func (td *ToDo) FromDomain(todo *domain.ToDo) {
	if td == nil {
		td = &ToDo{}
	}

	td.ID = todo.ID
	td.Title = todo.Title
	td.Description = todo.Description
}

func (td *ToDo) ToDomain() *domain.ToDo {
	if td == nil {
		td = &ToDo{}
	}

	return &domain.ToDo{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
	}
}

func (td ToDoList) FromDomain(tdms []domain.ToDo) {
	for _, t := range tdms {
		todo := ToDo{}
		todo.FromDomain(&t)
		td = append(td, todo)
	}
}
