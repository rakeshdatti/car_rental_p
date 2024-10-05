package repositories

import (
    "database/sql"
    "runners-mysql/models"
)

type BookingRepository struct {
    db *sql.DB
}

func NewBookingRepository(db *sql.DB) *BookingRepository {
    return &BookingRepository{db: db}
}

// func (repo *BookingRepository) CreateBooking(booking *models.Booking) error {
//     _, err := repo.db.Exec("INSERT INTO bookings (user_id, car_id, start, end) VALUES (?, ?, ?, ?)", booking.UserID, booking.CarID, booking.Start, booking.End)
//     return err
// }

func (repo *BookingRepository) CreateBooking(booking *models.Booking) (*models.Booking, error) {
  
    result, err := repo.db.Exec("INSERT INTO bookings (user_id, car_id, start, end) VALUES (?, ?, ?, ?)", booking.UserID, booking.CarID, booking.Start, booking.End)
    if err != nil {
        return nil, err
    }

    
    bookingID, err := result.LastInsertId()
    if err != nil {
        return nil, err
    }

   
    booking.ID = int(bookingID) 
    return booking, nil
}

