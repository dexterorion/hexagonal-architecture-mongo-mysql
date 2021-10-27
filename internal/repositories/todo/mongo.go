package todo

import (
	"context"

	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/domain"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type toDoMongo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}

type toDoListMongo []toDoMongo

func (m *toDoMongo) ToDomain() *domain.ToDo {
	return &domain.ToDo{
		ID:          m.ID.Hex(),
		Title:       m.Title,
		Description: m.Description,
	}
}
func (m *toDoMongo) FromDomain(todo *domain.ToDo) {
	if m == nil {
		m = &toDoMongo{}
	}

	m.ID = helpers.SafeObjectIdFromString(todo.ID)
	m.Title = todo.Title
	m.Description = todo.Description
}

func (m toDoListMongo) ToDomain() []domain.ToDo {
	todos := make([]domain.ToDo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

type todoMongoRepo struct {
	col *mongo.Collection
}

func NewTodoMongoRepo(col *mongo.Collection) ports.TodoRepository {
	return &todoMongoRepo{
		col: col,
	}
}

func (m *todoMongoRepo) Get(id string) (*domain.ToDo, error) {
	var todo toDoMongo
	result := m.col.FindOne(context.Background(), bson.M{"id": helpers.SafeObjectIdFromString(id)})

	if err := result.Decode(&todo); err != nil {
		return nil, err
	}

	return todo.ToDomain(), nil
}

func (m *todoMongoRepo) List() ([]domain.ToDo, error) {
	var todos toDoListMongo
	result, err := m.col.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err := result.All(context.Background(), todos); err != nil {
		return nil, err
	}

	return todos.ToDomain(), nil
}

func (m *todoMongoRepo) Create(todo *domain.ToDo) (*domain.ToDo, error) {
	var tdMongo *toDoMongo
	tdMongo.FromDomain(todo)

	_, err := m.col.InsertOne(context.Background(), tdMongo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
