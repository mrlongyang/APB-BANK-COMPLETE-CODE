
package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("your-secret-key")

func GenerateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(SecretKey)
}
