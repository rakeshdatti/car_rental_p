package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "runners-mysql/models"
    "runners-mysql/services"
)

type AuthController struct {
    authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
    return &AuthController{authService: authService}
}

func (ctrl *AuthController) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // user.Role = "user" 
    userID,err := ctrl.authService.Register(&user)
    if err!=nil{
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "message":"User registered sucessfully",
        "user_id": userID,
        "user_name": user.Username,
        "role": user.Role,
    })
}

func (ctrl *AuthController) AdminLogin(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    token,userID, err := ctrl.authService.Login(input.Username, input.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userID})

}

func (ctrl *AuthController) Login(c *gin.Context) {
    var loginDetails struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, userID, err := ctrl.authService.Login(loginDetails.Username, loginDetails.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userID})

}

