package usecases

import (
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers/logging"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/domain"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/ports"
)

var (
	log = logging.NewLogger()
)

type todoUseCase struct {
	todoRepo ports.TodoRepository
}

func NewToDoUseCase(todoRepo ports.TodoRepository) ports.TodoUseCase {
	return &todoUseCase{
		todoRepo: todoRepo,
	}
}

func (t *todoUseCase) Get(id string) (*domain.ToDo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		log.Errorw("Error getting from repo", logging.KeyID, id, logging.KeyErr, err)
		return nil, err
	}

	return todo, nil
}

func (t *todoUseCase) List() ([]domain.ToDo, error) {
	todos, err := t.todoRepo.List()
	if err != nil {
		log.Errorw("Error listing from repo", logging.KeyErr, err)
		return nil, err
	}

	return todos, nil
}

func (t *todoUseCase) Create(title, description string) (*domain.ToDo, error) {
	todo := domain.NewToDo(helpers.RandomUUIDAsString(), title, description)

	_, err := t.todoRepo.Create(todo)
	if err != nil {
		log.Errorw("Error creating from repo", "todo", todo, logging.KeyErr, err)
		return nil, err
	}

	return todo, nil
}
