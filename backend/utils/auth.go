package utils

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
)

// HashPassword menghasilkan password terenkripsi
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPasswordHash membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// GenerateJWT menghasilkan token JWT untuk user
func GenerateJWT(userID string, email string) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "your_super_secret_key_for_jwt_should_be_long_and_complex"
    }

    // Set waktu kedaluwarsa token (1 hari)
    expirationTime := time.Now().Add(24 * time.Hour)
    
    // Buat klaim untuk token
    claims := jwt.MapClaims{
        "id":    userID,
        "email": email,
        "exp":   expirationTime.Unix(),
    }
    
    // Buat token dengan algoritma HS256
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Tanda tangani token dengan secret key
    tokenString, err := token.SignedString([]byte(secret))
    
    return tokenString, err
}