package services

import (
	// "net/http"
	// "runners-mysql/models"
	"runners-mysql/repositories"
	// "strconv"
	// "time"
	"errors"
)

type AdminService struct {
	adminRepository *repositories.AdminRepository
	carRepository *repositories.CarRepository
	isLoggedIn bool
}

func NewAdminService(adminRepository *repositories.AdminRepository, carRepository *repositories.CarRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
		carRepository: carRepository,
		isLoggedIn: false,
	}
}

func (service *AdminService) Authenticate(username, password string) error {
    admin, err := service.adminRepository.AuthenticateAdmin(username, password)
    if err != nil || admin == nil {
        return errors.New("invalid credentials")
    }
    service.isLoggedIn = true // Set flag to true after successful login
    return nil
}

func (service *AdminService) IsAuthenticated() bool {
    return service.isLoggedIn
}


func (service *AdminService) Logout() {
    service.isLoggedIn = false
}