package utils

import (
    "errors"
    "os"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "go.mongodb.org/mongo-driver/bson/primitive"
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

// GetUserIDFromToken extracts user ID from JWT token in fiber context
func GetUserIDFromToken(c *fiber.Ctx) (primitive.ObjectID, error) {
    // Get token from Authorization header
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return primitive.NilObjectID, errors.New("authorization header missing")
    }

    // Extract token from "Bearer <token>"
    tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
    if tokenString == authHeader {
        return primitive.NilObjectID, errors.New("invalid authorization header format")
    }

    // Get JWT secret
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "your_super_secret_key_for_jwt_should_be_long_and_complex"
    }

    // Parse and validate token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return []byte(secret), nil
    })

    if err != nil {
        return primitive.NilObjectID, err
    }

    // Extract claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        if userIDStr, exists := claims["id"].(string); exists {
            userID, err := primitive.ObjectIDFromHex(userIDStr)
            if err != nil {
                return primitive.NilObjectID, errors.New("invalid user ID format")
            }
            return userID, nil
        }
        return primitive.NilObjectID, errors.New("user ID not found in token")
    }

    return primitive.NilObjectID, errors.New("invalid token")
}