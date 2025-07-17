package middleware

import (
    "fmt"
    "os"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
)

// Protected middleware untuk routes yang memerlukan autentikasi
func Protected() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get authorization header
        authHeader := c.Get("Authorization")
        
        // Check if auth header exists and has Bearer prefix
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized: missing or invalid authorization token",
            })
        }
        
        // Extract token from header
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
        // Parse and verify token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            
            // Get secret key from environment
            secret := os.Getenv("JWT_SECRET")
            if secret == "" {
                secret = "your_super_secret_key_for_jwt_should_be_long_and_complex"
            }
            
            return []byte(secret), nil
        })
        
        // Handle parsing errors
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": fmt.Sprintf("Unauthorized: %v", err),
            })
        }
        
        // Check if token is valid
        if !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized: invalid token",
            })
        }
        
        // Extract claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Unauthorized: failed to parse token claims",
            })
        }
        
        // Set user ID in locals for later use in handlers
        c.Locals("userID", claims["id"])
        c.Locals("userEmail", claims["email"])
        
        // Continue to the next handler
        return c.Next()
    }
}