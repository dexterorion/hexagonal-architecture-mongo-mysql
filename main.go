package main

import (
	"flag"
	"net/http"

	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/helpers/logging"
	"github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/ports"
	usecases "github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/core/usecases"
	handlerTodo "github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/handlers/todo"
	repoTodo "github.com/dexterorion/hexagonal-architecture-mongo-mysql/internal/repositories/todo"
	restful "github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

var (
	repo    string
	binding string

	log *zap.SugaredLogger = logging.NewLogger()
)

func init() {
	flag.StringVar(&repo, "repo", "mysql", "Mongo or MySql")
	flag.StringVar(&binding, "httpbind", ":8080", "address/port to bind listen socket")

	flag.Parse()
}

func main() {
	var todoRepo ports.TodoRepository
	if repo == "mysql" {
		todoRepo = startMysqlRepo()
	} else {
		todoRepo = startMongoRepo()
	}

	todoUseCase := usecases.NewToDoUseCase(todoRepo)

	ws := new(restful.WebService)
	ws = ws.Path("/api")
	handlerTodo.NewTodoHandler(todoUseCase, ws)
	restful.Add(ws)

	log.Info("Listening...")

	log.Panic(http.ListenAndServe(binding, nil))
}

func startMongoRepo() ports.TodoRepository {
	return repoTodo.NewTodoMongoRepo(helpers.StartMongoDb())
}

func startMysqlRepo() ports.TodoRepository {
	return repoTodo.NewTodoMysqlRepo(helpers.StartMysqlDb())
}
