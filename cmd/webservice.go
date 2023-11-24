package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/config"
	UserController "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/user/controller"
	UserRepository "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/user/repository/impl"
	UserService "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/user/service/impl"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/database"
	middleware2 "github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/middleware"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/model"

	GlobalController "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/controller"
	GlobalRepository "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/repository/impl"
	GlobalService "github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/service/impl"
)

type InitWebserviceParam struct {
	Conf *config.Config
}

func InitWebservice(params *InitWebserviceParam) {
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // set 100MB
	})

	db, err := database.InitMariaDB(&database.InitMariaDBParams{
		Conf: &params.Conf.MariaDBConfig,
	})

	db.AutoMigrate(&model.Todos{})

	if err != nil {
		log.Errorf("[webservice.go][InitWebservice] err init mysql :%+v", err)
		return
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowOriginsFunc: func(origin string) bool {
			return params.Conf.ENVIRONMENT == "dev"
		},
	}))

	app.Use(logger.New())  // Logger middleware
	app.Use(recover.New()) // --- recover panic

	api := app.Group("/api") // /api
	v1 := app
	api.Get("/health", monitor.New())

	// ------- User
	userRepository := UserRepository.New(&UserRepository.NewUserRepository{Conf: params.Conf, Db: db})
	userService := UserService.New(&UserService.NewUserServiceParams{Conf: params.Conf, UserRepository: userRepository})

	// ------  middleware
	middleware := middleware2.GlobalMiddleware{
		UserService: userService,
		Conf:        params.Conf,
	}

	userController := UserController.New(&UserController.UsersControllerParams{
		V1:          v1,
		Conf:        params.Conf,
		UserService: userService,
		Middleware:  middleware,
	})

	// --------- Main
	globalRepository := GlobalRepository.New(&GlobalRepository.NewGlobalRepository{
		Conf: params.Conf,
		Db:   db,
	})

	globalService := GlobalService.New(&GlobalService.NewGlobalServiceParams{
		Conf:             params.Conf,
		GlobalRepository: globalRepository,
		Db:               db,
	})

	globalController := GlobalController.New(&GlobalController.GlobalControllerParams{
		V1:            v1,
		Conf:          params.Conf,
		GlobalService: globalService,
		Middleware:    middleware,
	})

	userController.Init()
	globalController.Init()
	app.Listen(fmt.Sprintf(":%s", params.Conf.AppAddress))
}
