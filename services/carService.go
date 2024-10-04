package services

import (
	"runners-mysql/models"
	"runners-mysql/repositories"
)

type CarService struct {
	adminRepository *repositories.AdminRepository
	carRepository   *repositories.CarRepository
	admin           *models.Admin 
}
func NewCarService(adminRepository *repositories.AdminRepository, carRepository *repositories.CarRepository) *CarService {
	return &CarService{
		adminRepository: adminRepository,
		carRepository:   carRepository,
	}
}


func (s *CarService) SetAdmin(admin *models.Admin) {
	s.admin = admin
}

func (s *CarService) CreateCar(car models.Car) *models.ResponseError {
	if s.admin == nil {
		return &models.ResponseError{Message: "Unauthorized", Status: 401}
	}
	_, err := s.carRepository.Create(car) // Calls the Create function in CarRepository
	return err
}


func (s *CarService) GetAllCars() ([]models.Car, *models.ResponseError) {
	
	if s.admin == nil {
		return nil, &models.ResponseError{Message: "Unauthorized", Status: 401}
	}
	return s.carRepository.FindAll() 
}


func (s *CarService) GetCarByID(id int) (models.Car, *models.ResponseError) {
	if s.admin == nil {
		return models.Car{}, &models.ResponseError{Message: "Unauthorized", Status: 401}
	}
	return s.carRepository.FindByID(id) // Calls the FindByID function in CarRepository
}


func (s *CarService) UpdateCar(id int, car models.Car) *models.ResponseError {
	if s.admin == nil {
		return &models.ResponseError{Message: "Unauthorized", Status: 401}
	}
	return s.carRepository.Update(id, car) // Calls the Update function in CarRepository
}


func (s *CarService) DeleteCar(id int) *models.ResponseError {
	if s.admin == nil {
		return &models.ResponseError{Message: "Unauthorized", Status: 401}
	}
	return s.carRepository.Delete(id) // Calls the Delete function in CarRepository
}
