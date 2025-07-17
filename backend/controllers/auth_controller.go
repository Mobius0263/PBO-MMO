package controllers

import (
    "context"
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
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }

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
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menyimpan user"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Registrasi berhasil",
        "user": fiber.Map{
            "id":    user.ID,
            "nama":  user.Nama,
            "email": user.Email,
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
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }

    var user models.User
    err := config.UserCollectionRef.FindOne(ctx, bson.M{"email": input.Email}).Decode(&user)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email atau password salah"})
    }

    // Verify password
    match := utils.CheckPasswordHash(input.Password, user.Password)
    if !match {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email atau password salah"})
    }

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