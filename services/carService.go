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

func (s *CarService) AddCar(car *models.Car) (int64,error) {
    car.Status = "available" 
    lastInsertID,err := s.carRepo.AddCar(car)
    if err != nil {
        return 0, err
    }

    
    return lastInsertID, nil
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


func (s *CarService) DeleteCarByModel(model string)(bool,error){
    return s.carRepo.DeleteCarByModel(model)
}

func (s *CarService) UpdateCarID(id int64,model string)(bool,error){
    return s.carRepo.UpdateCarID(id,model)
}