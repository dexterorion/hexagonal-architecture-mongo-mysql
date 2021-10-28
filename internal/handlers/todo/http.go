package todo

import (
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/ports"
	restful "github.com/emicklei/go-restful/v3"
)

type TodoHandler struct {
	todoUseCase ports.TodoUseCase
}

func NewTodoHandler(todoUseCase ports.TodoUseCase, ws *restful.WebService) *TodoHandler {
	handler := &TodoHandler{
		todoUseCase: todoUseCase,
	}

	ws.Route(ws.GET("/todo/{id}").To(handler.Get).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.GET("/todo").To(handler.List).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))
	ws.Route(ws.POST("/todo").To(handler.Create).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON))

	return handler
}

func (tdh *TodoHandler) Get(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")

	result, err := tdh.todoUseCase.Get(id)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo *ToDo = &ToDo{}

	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}

func (tdh *TodoHandler) Create(req *restful.Request, resp *restful.Response) {
	var data = new(ToDo)
	if err := req.ReadEntity(data); err != nil {
		resp.WriteError(500, err)
		return
	}

	result, err := tdh.todoUseCase.Create(data.Title, data.Title)
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todo ToDo = ToDo{}
	todo.FromDomain(result)
	resp.WriteAsJson(todo)
}

func (tdh *TodoHandler) List(req *restful.Request, resp *restful.Response) {
	result, err := tdh.todoUseCase.List()
	if err != nil {
		resp.WriteError(500, err)
		return
	}

	var todos ToDoList = ToDoList{}

	todos = todos.FromDomain(result)
	resp.WriteAsJson(todos)
}
