package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey") // Troque por uma chave segura

func GenerateJWT(userID string, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "email": email,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
