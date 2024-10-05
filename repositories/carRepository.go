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

func (r *CarRepository) AddCar(car *models.Car) (int64, error) {
    query := "INSERT INTO cars (model, status) VALUES (?, ?)"
    
    result, err := r.db.Exec(query, car.Model, car.Status)
    if err != nil {
        return 0, err
    }

    lastInsertID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    
    return lastInsertID, nil
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


func (r *CarRepository) DeleteCarByModel(model string)(bool,error){
    query := "DELETE FROM cars WHERE model=?"
    result,err := r.db.Exec(query,model)
    if err != nil {
        return false, err
    }
 
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, err
    }
 
    return rowsAffected > 0, nil

}


func (r *CarRepository) UpdateCarID(id int64,model string) (bool, error) {
    query := "UPDATE cars SET Model = ? WHERE id= ?"
    result, err := r.db.Exec(query, model,id)
    if err != nil {
        return false, err
    }
 
 
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, err
    }

    return rowsAffected > 0, nil
}