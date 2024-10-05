package repositories

import (
    "database/sql"
    "runners-mysql/models"
    "errors"
    
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) (int64,error) {
    query := "INSERT INTO users (username, password, role) VALUES (?, ?, ?)"
    result, err := r.db.Exec(query, user.Username, user.Password, user.Role)
    if err!=nil{
        return 0,err
    }
    userID,err := result.LastInsertId()
    if err!=nil{
        return 0,err
    }
    return userID,nil
}


func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
    query := "SELECT id, username, password, role FROM users WHERE username = ?"
    row := r.db.QueryRow(query, username)

    var user models.User
    if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found") 
         }
        return nil, err
    }

    return &user, nil
}
