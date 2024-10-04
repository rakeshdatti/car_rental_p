package repositories

import (
	"database/sql"
	// "fmt"
	// "net/http"
	"runners-mysql/models"
	// "strconv"
)

type AdminRepository struct {
	dbHandler   *sql.DB

}

func NewAdminRepository(dbHandler *sql.DB) *AdminRepository {
	return &AdminRepository{
		dbHandler: dbHandler,
	}
}


func (repo *AdminRepository) AuthenticateAdmin(username, password string) (*models.Admin, error) {
    var admin models.Admin
    query := "SELECT id, username, password FROM admins WHERE username = ? AND password = ?"
    err := repo.dbHandler.QueryRow(query, username, password).Scan(&admin.ID, &admin.Username, &admin.Password)
    if err != nil {
        return nil, err // Return nil if no matching admin found
    }
    return &admin, nil
}