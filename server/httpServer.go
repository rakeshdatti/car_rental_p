package server

import (
	"database/sql"
	"log"
	"runners-mysql/controllers"
	"runners-mysql/repositories"
	"runners-mysql/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	adminController *controllers.AdminController
	carController *controllers.CarController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	adminRepository := repositories.NewAdminRepository(dbHandler)
	carRepository := repositories.NewCarRepository(dbHandler)
	adminService := services.NewAdminService(adminRepository, carRepository)
	carService := services.NewCarService(adminRepository, carRepository)
	adminController := controllers.NewAdminController(adminService)
	carController := controllers.NewCarController(carService)

	router := gin.Default()

	router.POST("/admin/login", adminController.Login)
    router.POST("/admin/logout", adminController.Logout)


	router.GET("/cars", carController.GetAllCars)
	router.GET("/cars/:id", carController.GetCarByID)
	router.POST("/cars", carController.CreateCar)
	router.PUT("/cars/:id", carController.UpdateCar)
	router.DELETE("/cars/:id", carController.DeleteCar)

	return HttpServer{
		config:            config,
		router:            router,
		adminController: adminController,
		carController: carController,
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
