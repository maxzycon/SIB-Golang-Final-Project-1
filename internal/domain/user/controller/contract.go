package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/user/service"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/constant/role"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/middleware"
)

const (
	Create = "/users"

	Update            = "/users/:id"
	GetById           = "/users/:id"
	Delete            = "/users/:id"
	GetUsersPaginated = "/users"
	UpdateUserProfile = "/profile"

	Login        = "/login"
	GetUserLogin = "/user"
)

type UsersControllerParams struct {
	V1          fiber.Router
	Conf        *config.Config
	UserService service.UserService
	Middleware  middleware.GlobalMiddleware
}
type usersController struct {
	v1          fiber.Router
	conf        *config.Config
	userService service.UserService
	middleware  middleware.GlobalMiddleware
}

func New(params *UsersControllerParams) *usersController {
	return &usersController{
		v1:          params.V1,
		conf:        params.Conf,
		userService: params.UserService,
		middleware:  params.Middleware,
	}
}
func (pc *usersController) Init() {
	pc.v1.Post(Create, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerCreateUser)
	pc.v1.Put(Update, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerUpdateUser)
	pc.v1.Get(GetUserLogin, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerUser)
	pc.v1.Get(GetUsersPaginated, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerGetUsersPaginated)
	pc.v1.Delete(Delete, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerDeleteUserById)
	pc.v1.Get(GetById, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handlerGetUserById)

	pc.v1.Put(UpdateUserProfile, pc.middleware.Protected([]uint{role.ROLE_SYSTEM, role.ROLE_MANAGER}), pc.handerUpdateUserProfile)

	pc.v1.Post(Login, pc.handlerLogin)
}
