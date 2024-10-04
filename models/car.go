package models


type Car struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Model    string  `json:"model"`
    Price    float64 `json:"price"`
    Available bool   `json:"available"`
}
