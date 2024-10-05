package services

import (
    "runners-mysql/models"
    "runners-mysql/repositories"
)

type BookingService struct {
    bookingRepo *repositories.BookingRepository
}

func NewBookingService(bookingRepo *repositories.BookingRepository) *BookingService {
    return &BookingService{bookingRepo: bookingRepo}
}



func (s *BookingService) CreateBooking(booking *models.Booking) (*models.Booking, error) {
    return s.bookingRepo.CreateBooking(booking)
}

