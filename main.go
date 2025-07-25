package main

import (
	"fmt"
	configuration "go_fiber_core_project_api/configuration/app"
	database "go_fiber_core_project_api/configuration/database"
	redis "go_fiber_core_project_api/configuration/redis"
	custom_translate "go_fiber_core_project_api/configuration/translate"
	handler "go_fiber_core_project_api/handler"
	custom_logger "go_fiber_core_project_api/pkg/utils/loggers"
	routers "go_fiber_core_project_api/routers"
)

func main() {

	// INITIAL CONFIGURATION
	app_configuration := configuration.NewConfiguration()

	// INITIALIZE DATABASE
	db_pool := database.GetDB()

	// INITIALIZE ROUTER
	app := routers.New(db_pool)

	// INITIALIZE REDIS
	redis := redis.NewRedisClient()

	// INITIALIZE THE TRANSLATE
	if err := custom_translate.InitTranslate(); err != nil {
		custom_logger.NewCustomLog("Failed_initialize_i18n", err.Err.Error(), "error")
	}

	// HANDLER CONTROLL WHOLE ROUTE
	handler.NewFrontService(app, db_pool, redis)

	// START APPLICATION
	err := app.Listen(fmt.Sprintf("%s:%d", app_configuration.AppHost, app_configuration.AppPort))
	if err != nil {
		custom_logger.NewCustomLog("server_failed_to_start", err.Error(), "error")
		panic(err)
	}
}
