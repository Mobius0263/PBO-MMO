package controllers

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "pbommo/config"
    "pbommo/models"
)

// GetUsers returns all users
func GetUsers(c *fiber.Ctx) error {
    var users []models.UserResponse
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := config.UserCollectionRef.Find(ctx, bson.M{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user models.User
        if err := cursor.Decode(&user); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        // Convert to user response without password
        userResponse := models.UserResponse{
            ID:           user.ID,
            Nama:         user.Nama,
            Email:        user.Email,
            Role:         user.Role,
            Bio:          user.Bio,
            ProfileImage: user.ProfileImage,
            Status:       "Online", // Placeholder
            LastActive:   time.Now(),
        }

        users = append(users, userResponse)
    }

    return c.Status(200).JSON(users)
}

// GetUserById returns a specific user by ID
func GetUserById(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var user models.User
    err = config.UserCollectionRef.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    // Convert to user response without password
    userResponse := models.UserResponse{
        ID:           user.ID,
        Nama:         user.Nama,
        Email:        user.Email,
        Role:         user.Role,
        Bio:          user.Bio,
        ProfileImage: user.ProfileImage,
    }

    return c.Status(200).JSON(userResponse)
}

// UpdateUser updates user information
func UpdateUser(c *fiber.Ctx) error {
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // Remove password from update data if present (should be updated separately)
    delete(updateData, "password")
    
    // Set updated timestamp
    updateData["updatedAt"] = time.Now()

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    result, err := config.UserCollectionRef.UpdateOne(
        ctx,
        bson.M{"_id": id},
        bson.M{"$set": updateData},
    )

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    if result.MatchedCount == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    return c.Status(200).JSON(fiber.Map{"message": "User updated successfully"})
}

// UploadProfileImage uploads and saves a profile image
func UploadProfileImage(c *fiber.Ctx) error {
    // Get user ID from params
    idParam := c.Params("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    // Check if the user exists
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    var user models.User
    err = config.UserCollectionRef.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "User not found"})
    }

    // Get the file from request
    file, err := c.FormFile("profileImage")
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
    }

    // Create uploads directory if it doesn't exist
    uploadDir := "./uploads/profiles"
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        err = os.MkdirAll(uploadDir, 0755)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to create uploads directory"})
        }
    }

    // Generate file name with timestamp to avoid duplicates
    timeStamp := time.Now().UnixNano()
    fileName := fmt.Sprintf("%d_%s", timeStamp, file.Filename)
    filePath := fmt.Sprintf("%s/%s", uploadDir, fileName)

    // Save the file
    if err := c.SaveFile(file, filePath); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to save the file"})
    }

    // Create image URL - path must start with '/'
    imageURL := fmt.Sprintf("/uploads/profiles/%s", fileName)

    // Update user profile image in database
    _, err = config.UserCollectionRef.UpdateOne(
        ctx,
        bson.M{"_id": id},
        bson.M{"$set": bson.M{
            "profileImage": imageURL,
            "updatedAt":    time.Now(),
        }},
    )

    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile image"})
    }

    return c.Status(200).JSON(fiber.Map{
        "message": "Profile image uploaded successfully",
        "imageUrl": imageURL,
    })
}