package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "runners-mysql/models"
    "runners-mysql/services"
)

type CarsController struct {
    carService *services.CarService
}

func NewCarsController(carService *services.CarService) *CarsController {
    return &CarsController{carService: carService}
}

func (ctrl *CarsController) AddCar(c *gin.Context) {
    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := ctrl.carService.AddCar(&car); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, car)
}

func (ctrl *CarsController) GetAvailableCars(c *gin.Context) {
    cars, err := ctrl.carService.GetAvailableCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cars)
}

func (ctrl *CarsController) GetAllCars(c *gin.Context) {
    cars, err := ctrl.carService.GetAllCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cars)
}

func (ctrl *CarsController) GetCarsByModel(c *gin.Context) {
    model := c.Param("model")
    cars ,err := ctrl.carService.GetCarsByModel(model)
    if err!=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        return
    }

    if len(cars)==0{
        c.JSON(http.StatusNotFound,gin.H{"error":"No cars found for given model"})
        return
    }
    c.JSON(http.StatusOK,cars)
}
