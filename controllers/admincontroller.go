package controllers

import (
	// "encoding/json"
	// "io"
	"log"
	"net/http"
	// "runners-mysql/models"
	"github.com/google/uuid"
	"runners-mysql/services"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

func (ctrl *AdminController) Login(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Error while reading result request body", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
    }

    if err := ctrl.adminService.Authenticate(input.Username, input.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
	sessionID := uuid.New().String()
	c.SetCookie("session_id", sessionID, 3600, "/", "", false, true) 
	c.JSON(http.StatusOK, gin.H{
        "message":    "Login successful",
        "session_id": sessionID,
    })
}

// Admin logout
func (ctrl *AdminController) Logout(c *gin.Context) {
    ctrl.adminService.Logout()
    c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}



// curl -X POST http://localhost:8080/admin/login \
//   -H "Content-Type: application/json" \
//   -d '{"username": "admin1", "password": "admin123"}'
