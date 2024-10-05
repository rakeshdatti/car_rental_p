// server/httpServer.go
package server

import (
    "database/sql"
    "log"
    "os"
    "runners-mysql/controllers"
    "runners-mysql/repositories"
    "runners-mysql/services"
    "runners-mysql/middlewares" // Ensure this is imported
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

type HttpServer struct {
    config            *viper.Viper
    router            *gin.Engine
    authController    *controllers.AuthController
    carsController    *controllers.CarsController
    bookingsController *controllers.BookingsController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
    userRepository := repositories.NewUserRepository(dbHandler)
    carRepository := repositories.NewCarRepository(dbHandler)
    bookingRepository := repositories.NewBookingRepository(dbHandler)

    jwtSecret := os.Getenv("JWT_SECRET")
    
    authService := services.NewAuthService(userRepository, jwtSecret)
    carService := services.NewCarService(carRepository)
    bookingService := services.NewBookingService(bookingRepository)

    authController := controllers.NewAuthController(authService)
    carsController := controllers.NewCarsController(carService)
    bookingsController := controllers.NewBookingsController(bookingService)

    router := gin.Default()
   
    router.POST("/register", authController.Register)
    router.POST("/login", authController.Login)
    router.GET("/cars", carsController.GetAvailableCars)
    router.POST("/bookings", middlewares.AuthMiddleware(jwtSecret), bookingsController.CreateBooking)
    router.GET("/cars/:model", carsController.GetCarsByModel)
    router.GET("/cars/available",carsController.GetAvailableCars)

    // Admin routes
    router.POST("/admin/login", authController.AdminLogin)
    router.POST("/admin/cars", carsController.AddCar) 
    router.GET("/admin/cars", carsController.GetAllCars) 
    

    return HttpServer{
        config:            config,
        router:            router,
        authController:    authController,
        carsController:    carsController,
        bookingsController: bookingsController,
    }
}

func (hs HttpServer) Start() {
    err := hs.router.Run(hs.config.GetString("http.server_address"))
    if err != nil {
        log.Fatalf("Error while starting HTTP server: %v", err)
    }
}
