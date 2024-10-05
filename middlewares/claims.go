
package middlewares

import (
    "github.com/dgrijalva/jwt-go"
)


type Claims struct {
    UserID int `json:"id"`
    Role   string `json: "role"`
    jwt.StandardClaims
}
