package repositories

import (
    "database/sql"
    "runners-mysql/models"
  
)

type CarRepository struct {
    db *sql.DB
}

func NewCarRepository(db *sql.DB) *CarRepository {
    return &CarRepository{db: db}
}

func (r *CarRepository) AddCar(car *models.Car) error {
    query := "INSERT INTO cars ( model, status) VALUES (?, ?)"
    _, err := r.db.Exec(query, car.Model, car.Status)
    return err
}

func (r *CarRepository) GetAvailableCars() ([]models.Car, error) {
    query := "SELECT id,model, status FROM cars WHERE status = ?"
    
    rows, err := r.db.Query(query, "available")
  
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cars []models.Car
    for rows.Next() {
        var car models.Car
        if err := rows.Scan(&car.ID,  &car.Model, &car.Status); err != nil {
            return nil, err
        }
        cars = append(cars, car)
    }

    return cars, nil
}

func (r *CarRepository) GetAllCars() ([]models.Car, error) {
    query := "SELECT id, model, status FROM cars"
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cars []models.Car
    for rows.Next() {
        var car models.Car
        if err := rows.Scan(&car.ID,  &car.Model, &car.Status); err != nil {
            return nil, err
        }
        cars = append(cars, car)
    }

    return cars, nil
}


func (r *CarRepository) GetCarsByModel(model string) ([]models.Car, error) {
    query := "SELECT id, model,status FROM cars WHERE model = ?"
    rows, err := r.db.Query(query, model)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cars []models.Car
    for rows.Next() {
        var car models.Car
        if err := rows.Scan(&car.ID, &car.Model, &car.Status); err != nil {
            return nil, err
        }
        cars = append(cars, car)
    }

    return cars, nil
}
