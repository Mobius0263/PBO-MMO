package controllers

import (
    "context"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "pbommo/config"
    "pbommo/models"
    "pbommo/utils"
)

func Register(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user models.User
    if err := c.BodyParser(&user); err != nil {
        log.Printf("BodyParser error: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }
    
    log.Printf("Received registration data: %+v", user)
    log.Printf("Password length received: %d", len(user.Password))

    // Check if email already exists
    var existingUser models.User
    err := config.UserCollectionRef.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
    if err == nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email sudah terdaftar"})
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal hash password"})
    }
    user.Password = hashedPassword

    // Set timestamps
    now := time.Now()
    user.CreatedAt = now
    user.UpdatedAt = now
    
    // Set ID if not provided
    if user.ID.IsZero() {
        user.ID = primitive.NewObjectID()
    }

    // Default role if not provided
    if user.Role == "" {
        user.Role = "User"
    }

    _, err = config.UserCollectionRef.InsertOne(ctx, user)
    if err != nil {
        log.Printf("Database insertion error: %v", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan user"})
    }
    
    log.Printf("User registered successfully: %s", user.Email)

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Registrasi berhasil",
        "user": fiber.Map{
            "id":    user.ID,
            "nama":  user.Nama,
            "email": user.Email,
            "role":  user.Role,
        },
    })
}

func Login(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.BodyParser(&input); err != nil {
        log.Printf("Login BodyParser error: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }
    
    log.Printf("Login attempt for email: %s", input.Email)
    log.Printf("Password length received: %d", len(input.Password))

    var user models.User
    err := config.UserCollectionRef.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        log.Printf("User not found for email %s: %v", input.Email, err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email atau password salah"})
    }

    // Verify password
    log.Printf("Checking password for user: %s", user.Email)
    match := utils.CheckPasswordHash(input.Password, user.Password)
    if !match {
        log.Printf("Password mismatch for user: %s", user.Email)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email atau password salah"})
    }
    
    log.Printf("Login successful for user: %s", user.Email)

    // Generate JWT
    token, err := utils.GenerateJWT(user.ID.Hex(), user.Email)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal generate token"})
    }

    // User response without password
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Login berhasil",
        "token":   token,
        "user": fiber.Map{
            "id":           user.ID,
            "nama":         user.Nama,
            "email":        user.Email,
            "role":         user.Role,
            "profileImage": user.ProfileImage,
        },
    })
}