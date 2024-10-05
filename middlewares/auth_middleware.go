// middleware/auth.go
package middlewares

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "fmt"
    "log"
)
func AuthMiddleware(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Entering AuthMiddleware")
        authHeader := c.Request.Header.Get("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
            c.Abort()
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        tokenString := parts[1]
        claims := &Claims{} 
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil {
            fmt.Printf("Error parsing token: %v\n", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        log.Println(claims.UserID,claims.Role)
        c.Set("userID", claims.UserID)
        c.Set("role", claims.Role)
         
        c.Next()
    }
}
