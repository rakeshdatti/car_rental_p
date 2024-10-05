package services

import (
    "errors"
    "time"
    "github.com/dgrijalva/jwt-go"
    "runners-mysql/models"
    "runners-mysql/repositories"
)

type AuthService struct {
    userRepo *repositories.UserRepository
    secret   []byte
}

func NewAuthService(userRepo *repositories.UserRepository, secret string) *AuthService {
    return &AuthService{
        userRepo: userRepo,
        secret:   []byte(secret), 
    }
}

func (s *AuthService) Register(user *models.User) (int64,error) {
   
    return s.userRepo.CreateUser(user)
}

func (s *AuthService) Login(username, password string) (string, int, error) {
    user, err := s.userRepo.GetUserByUsername(username)
    if err != nil {
        return "", 0, err 
    }

    
    if user.Password != password { 
        return "", 0, errors.New("invalid credentials") }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":    user.ID,
        "role":  user.Role,
        "exp":   time.Now().Add(time.Hour * 72).Unix(),
    })

  
    tokenString, err := token.SignedString(s.secret)
    if err != nil {
        return "", 0, err
    }

    return tokenString, user.ID, nil 
}
