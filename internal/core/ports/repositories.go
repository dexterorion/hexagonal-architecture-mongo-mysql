package ports

import (
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/domain"
)

type TodoRepository interface {
	Get(id string) (*domain.ToDo, error)
	List() ([]domain.ToDo, error)
	Create(todo *domain.ToDo) (*domain.ToDo, error)
}
