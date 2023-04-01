package main

import (
	"fmt"

	env "github.com/caarlos0/env/v6"
	validator "github.com/go-playground/validator/v10"
	fiber "github.com/gofiber/fiber/v2"
	dotenv "github.com/joho/godotenv"
	configs "github.com/joseesco24/users_crud_api_go/src/configs"
	routers "github.com/joseesco24/users_crud_api_go/src/routers"
	logging "github.com/sirupsen/logrus"
)

func main() {

	const envFilePath string = "../.env"

	var validate *validator.Validate
	var app *fiber.App
	var err error

	validate = validator.New()
	app = fiber.New()

	err = dotenv.Load(envFilePath)
	if err != nil {
		logging.Warn(fmt.Sprintf("error loading env file from %q", envFilePath))
	}

	err = env.Parse(configs.Config)
	if err != nil {
		logging.Warn(fmt.Sprintf("%+v", err.Error()))
		return
	}

	err = validate.Struct(configs.Config)
	if err != nil {
		logging.Warn(fmt.Sprintf("%+v", err.Error()))
		return
	}

	app.Post("/graphql", routers.MainRouter.Handler.ServeHTTP)

	logging.Info(fmt.Sprintf("starting api on %q mode", configs.Config.GetAppMode()))
	var appListenString string = fmt.Sprintf(":%d", configs.Config.GetAppPort())

	err = app.Listen(appListenString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
