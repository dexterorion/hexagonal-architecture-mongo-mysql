package todo

import (
	"database/sql"
	"fmt"

	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/domain"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/ports"
)

type toDoMysql struct {
	ID          string
	Title       string
	Description string
}

type toDoListMysql []toDoMysql

func (m *toDoMysql) ToDomain() *domain.ToDo {
	return &domain.ToDo{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
	}
}
func (m *toDoMysql) FromDomain(todo *domain.ToDo) {
	if m == nil {
		m = &toDoMysql{}
	}

	m.ID = todo.ID
	m.Title = todo.Title
	m.Description = todo.Description
}

func (m toDoListMysql) ToDomain() []domain.ToDo {
	todos := make([]domain.ToDo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

type todoMysqlRepo struct {
	db *sql.DB
}

func NewTodoMysqlRepo(db *sql.DB) ports.TodoRepository {
	return &todoMysqlRepo{
		db: db,
	}
}

func (m *todoMysqlRepo) Get(id string) (*domain.ToDo, error) {
	var todo toDoMysql = toDoMysql{}
	sqsS := fmt.Sprintf("SELECT id, title, description FROM todo WHERE id = '%s'", id)

	result := m.db.QueryRow(sqsS)
	if result.Err() != nil {
		return nil, result.Err()
	}

	if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (m *todoMysqlRepo) List() ([]domain.ToDo, error) {
	var todos toDoListMysql
	sqsS := "SELECT id, title, description FROM todo"

	result, err := m.db.Query(sqsS)
	if err != nil {
		return nil, err
	}

	if result.Err() != nil {
		return nil, result.Err()
	}

	for result.Next() {
		todo := toDoMysql{}

		if err := result.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos.ToDomain(), nil
}

func (m *todoMysqlRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	sqlS := "INSERT INTO todo (id, title, description) VALUES (?, ?, ?)"

	_, err := m.db.Exec(sqlS, todo.ID, todo.Title, todo.Description)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
