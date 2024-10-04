package repositories

import (
	"database/sql"
	"fmt"
	// "net/http"
	"runners-mysql/models"
	// "strconv"
)

type CarRepository struct {
	dbHandler   *sql.DB
	
}

func NewCarRepository(dbHandler *sql.DB) *CarRepository {
	return &CarRepository{
		dbHandler: dbHandler,
	}
}

func (repo *CarRepository) FindAll() ([]models.Car, *models.ResponseError) {
    query := "SELECT id, name, model, price, available FROM cars"
    rows, err := repo.dbHandler.Query(query)
    if err != nil {
        return nil, &models.ResponseError{
            Message: "Failed to fetch cars",
            Status:  500,
        }
    }
    defer rows.Close()

    var cars []models.Car
    for rows.Next() {
        var car models.Car
        if err := rows.Scan(&car.ID, &car.Name, &car.Model, &car.Price, &car.Available); err != nil {
            return nil, &models.ResponseError{
                Message: "Failed to scan car data",
                Status:  500,
            }
        }
        cars = append(cars, car) // Append the scanned car to the slice
    }

    // Check for errors encountered while iterating over rows
    if err := rows.Err(); err != nil {
        return nil, &models.ResponseError{
            Message: "Failed to iterate over car data",
            Status:  500,
        }
    }

    return cars, nil
}



func (repo *CarRepository) FindByID(id int) (models.Car, *models.ResponseError) {
	query := "SELECT id, name, model, price, available FROM cars WHERE id = ?"
	row := repo.dbHandler.QueryRow(query, id)

	var car models.Car
	err := row.Scan(&car.ID, &car.Name, &car.Model, &car.Price, &car.Available)
	if err == sql.ErrNoRows {
		return car, &models.ResponseError{
			Message: fmt.Sprintf("Car with ID %d not found", id),
			Status:  404,
		}
	} else if err != nil {
		return car, &models.ResponseError{
			Message: "Failed to fetch car",
			Status:  500,
		}
	}

	return car, nil
}


func (repo *CarRepository) Create(car models.Car) (int, *models.ResponseError) {
	query := "INSERT INTO cars (name, model, price, available) VALUES (?, ?, ?, ?)"
	result, err := repo.dbHandler.Exec(query, car.Name, car.Model, car.Price, car.Available)
	if err != nil {
		return 0, &models.ResponseError{
			Message: "Failed to create car",
			Status:  500,
		}
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, &models.ResponseError{
			Message: "Failed to retrieve new car ID",
			Status:  500,
		}
	}

	return int(id), nil
}


func (repo *CarRepository) Update(id int, car models.Car) *models.ResponseError {
	query := "UPDATE cars SET name = ?, model = ?, price = ?, available = ? WHERE id = ?"
	_, err := repo.dbHandler.Exec(query, car.Name, car.Model, car.Price, car.Available, id)
	if err != nil {
		return &models.ResponseError{
			Message: "Failed to update car",
			Status:  500,
		}
	}
	return nil
}

func (repo *CarRepository) Delete(id int) *models.ResponseError {
	query := "DELETE FROM cars WHERE id = ?"
	_, err := repo.dbHandler.Exec(query, id)
	if err != nil {
		return &models.ResponseError{
			Message: "Failed to delete car",
			Status:  500,
		}
	}
	return nil
}