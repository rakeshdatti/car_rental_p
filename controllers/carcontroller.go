package controllers

import (
	// "log"
	"net/http"
	"strconv" // Import strconv for string to int conversion
	"runners-mysql/models"
	"runners-mysql/services"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	carService *services.CarService
}

func NewCarController(carService *services.CarService) *CarController {
	return &CarController{
		carService: carService,
	}
}

func (controller *CarController) GetAllCars(c *gin.Context) {
	// Check for session_id cookie
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	cars, err := controller.carService.GetAllCars()
	if err != nil {
		// Assuming err is a custom error type, handle accordingly.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cars"})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func (controller *CarController) GetCarByID(c *gin.Context) {
    // Check for session_id cookie
    sessionID, err := c.Cookie("session_id")
    if err != nil || sessionID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    // Get car ID from the URL parameters
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    // Fetch the car using the service
    car, err := controller.carService.GetCarByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
        return
    }

    c.JSON(http.StatusOK, car)
}


func (controller *CarController) CreateCar(c *gin.Context) {
	// Check for session_id cookie
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Message: "Invalid input", Status: http.StatusBadRequest})
		return
	}

	// Call CreateCar from the car service
	err = controller.carService.CreateCar(car) // Only error is returned
	if err != nil {
		// Handle your custom error response here if necessary
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Car created successfully"}) // Respond with success message
}

func (controller *CarController) UpdateCar(c *gin.Context) {
	// Check for session_id cookie
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Message: "Invalid input", Status: http.StatusBadRequest})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert the ID to an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	err = controller.carService.UpdateCar(id, car) // Pass the id and car
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car updated successfully"})
}

func (controller *CarController) DeleteCar(c *gin.Context) {
	// Check for session_id cookie
	sessionID, err := c.Cookie("session_id")
	if err != nil || sessionID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Convert the ID to an integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	err = controller.carService.DeleteCar(id) // Pass the id
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete car"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Car deleted successfully"})
}
