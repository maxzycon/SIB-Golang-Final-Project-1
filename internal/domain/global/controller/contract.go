package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/service"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/middleware"
)

const (
	todosBasePath   = "todos"
	todosBaseIdPath = "todos/:id"
)

type GlobalControllerParams struct {
	V1            fiber.Router
	Conf          *config.Config
	GlobalService service.GlobalService
	Middleware    middleware.GlobalMiddleware
}
type GlobalController struct {
	v1            fiber.Router
	conf          *config.Config
	globalService service.GlobalService
	middleware    middleware.GlobalMiddleware
}

func New(params *GlobalControllerParams) *GlobalController {
	return &GlobalController{
		v1:            params.V1,
		conf:          params.Conf,
		globalService: params.GlobalService,
		middleware:    params.Middleware,
	}
}

func (pc *GlobalController) Init() {
	// ---- Todos API
	pc.v1.Get(todosBasePath, pc.handlerGetAllTodos)
	pc.v1.Get(todosBaseIdPath, pc.handlerGetTodosById)
	pc.v1.Post(todosBasePath, pc.handlerCreateTodos)
	pc.v1.Put(todosBaseIdPath, pc.handlerUpdateTodos)
	pc.v1.Delete(todosBaseIdPath, pc.handlerDeleteTodos)
}
