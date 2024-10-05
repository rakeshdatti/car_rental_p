package models



type Car struct {
    ID     int    `json:"id"`
    Model  string `json:"model"`
    Status string `json:"status"` // 'available' or 'booked'
}
