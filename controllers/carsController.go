package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "runners-mysql/models"
    "runners-mysql/services"
    "log"
    "strconv"
    
)

type CarsController struct {
    carService *services.CarService
}

func NewCarsController(carService *services.CarService) *CarsController {
    return &CarsController{carService: carService}
}

func (ctrl *CarsController) AddCar(c *gin.Context) {
  
      
    role,exists:= c.Get("role")
    log.Println(role)

    if exists==false {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
        return
    }

    if !exists || role !="admin"{
        c.JSON(http.StatusForbidden,gin.H{"error":"Access denied.Admins only can add cars"})
        return
    }
    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    id,err := ctrl.carService.AddCar(&car);     
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    car.ID=int(id)
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
func (ctrl *CarsController) DeleteCarByModel(c *gin.Context){
    role,exists:= c.Get("role")
    log.Println(role)

    if exists==false {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
        return
    }

    if !exists || role !="admin"{
        c.JSON(http.StatusForbidden,gin.H{"error":"Access denied.Admins only can add cars"})
        return
    }
    model := c.Param("model")
    if model == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Model is required"})
        return
    }
    deleted,err :=ctrl.carService.DeleteCarByModel(model)
    if err!=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        return
    }
    
    if !deleted {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car with the given model not found"})
        return
    }


    c.JSON(http.StatusOK,gin.H{"message":"Deleted Sucessfully"})
}


func (ctrl *CarsController) UpdateCarID(c *gin.Context){
    role,exists:= c.Get("role")
    log.Println(role)

    if exists==false {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Please login"})
        return
    }

    if !exists || role !="admin"{
        c.JSON(http.StatusForbidden,gin.H{"error":"Access denied.Admins only can add cars"})
        return
    }
  
    id := c.Param("id") 
    carID, err := strconv.ParseInt(id, 10, 64)  
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updated,err :=ctrl.carService.UpdateCarID(carID,car.Model)
    
    if err!=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        return
    }
    
    if !updated {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car with the given id not found"})
        return
    }
    c.JSON(http.StatusOK,gin.H{"message":"Updated Sucessfully"})
}