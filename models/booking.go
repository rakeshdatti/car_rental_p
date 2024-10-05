package models

import "time"


type Booking struct {
    ID      int       `json:"id"`
    UserID  int       `json:"user_id"`
    CarID   int       `json:"car_id"`
    Start   time.Time `json:"start"`
    End     time.Time `json:"end"`
}
