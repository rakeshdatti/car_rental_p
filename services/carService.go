package services

import (
    "runners-mysql/models"
    "runners-mysql/repositories"
)

type CarService struct {
    carRepo *repositories.CarRepository
}

func NewCarService(carRepo *repositories.CarRepository) *CarService {
    return &CarService{carRepo: carRepo}
}

func (s *CarService) AddCar(car *models.Car) error {
    car.Status = "available" // Default status
    return s.carRepo.AddCar(car)
}

func (s *CarService) GetAvailableCars() ([]models.Car, error) {
    return s.carRepo.GetAvailableCars()
}

func (s *CarService) GetAllCars() ([]models.Car, error) {
    return s.carRepo.GetAllCars()
}
func (s *CarService) GetCarsByModel(model string) ([]models.Car, error) {
    return s.carRepo.GetCarsByModel(model)
}
