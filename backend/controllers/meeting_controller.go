package controllers

import (
    "context"
    "time"

    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"

    "pbommo/config"
    "pbommo/models"
    "pbommo/utils"
)

// CreateMeeting creates a new meeting
func CreateMeeting(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var meeting models.Meeting
    if err := c.BodyParser(&meeting); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }

    // Get user ID from JWT token
    userID, err := utils.GetUserIDFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
    }

    // Set meeting metadata
    meeting.ID = primitive.NewObjectID()
    meeting.CreatedBy = userID
    meeting.CreatedAt = time.Now()
    meeting.UpdatedAt = time.Now()

    // Validate required fields
    if meeting.Title == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Judul meeting diperlukan"})
    }

    if meeting.StartTime.IsZero() {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Waktu mulai meeting diperlukan"})
    }

    if meeting.Duration <= 0 {
        meeting.Duration = 60 // Default 60 minutes
    }

    // Insert meeting into database
    result, err := config.MeetingCollectionRef.InsertOne(ctx, meeting)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat meeting"})
    }

    meeting.ID = result.InsertedID.(primitive.ObjectID)

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Meeting berhasil dibuat",
        "meeting": meeting,
    })
}

// GetMeetings gets all meetings with optional filtering
func GetMeetings(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get user ID from JWT token
    userID, err := utils.GetUserIDFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
    }

    // Build filter
    filter := bson.M{
        "$or": []bson.M{
            {"createdBy": userID},
            {"participants": userID},
        },
    }

    // Optional date filter
    if dateStr := c.Query("date"); dateStr != "" {
        date, err := time.Parse("2006-01-02", dateStr)
        if err == nil {
            startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
            endOfDay := startOfDay.Add(24 * time.Hour)
            filter["startTime"] = bson.M{
                "$gte": startOfDay,
                "$lt":  endOfDay,
            }
        }
    }

    cursor, err := config.MeetingCollectionRef.Find(ctx, filter)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data meeting"})
    }
    defer cursor.Close(ctx)

    var meetings []models.Meeting
    if err := cursor.All(ctx, &meetings); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal memproses data meeting"})
    }

    return c.JSON(fiber.Map{
        "meetings": meetings,
    })
}

// GetMeetingById gets a specific meeting by ID
func GetMeetingById(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get meeting ID from URL parameter
    meetingID, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID meeting tidak valid"})
    }

    // Get user ID from JWT token
    userID, err := utils.GetUserIDFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
    }

    var meeting models.Meeting
    err = config.MeetingCollectionRef.FindOne(ctx, bson.M{"_id": meetingID}).Decode(&meeting)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Meeting tidak ditemukan"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data meeting"})
    }

    // Check if user has access to this meeting
    hasAccess := meeting.CreatedBy == userID
    for _, participant := range meeting.Participants {
        if participant == userID {
            hasAccess = true
            break
        }
    }

    if !hasAccess {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Akses ditolak"})
    }

    return c.JSON(fiber.Map{
        "meeting": meeting,
    })
}

// UpdateMeeting updates an existing meeting
func UpdateMeeting(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get meeting ID from URL parameter
    meetingID, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID meeting tidak valid"})
    }

    // Get user ID from JWT token
    userID, err := utils.GetUserIDFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
    }

    // Check if meeting exists and user is the creator
    var existingMeeting models.Meeting
    err = config.MeetingCollectionRef.FindOne(ctx, bson.M{"_id": meetingID}).Decode(&existingMeeting)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Meeting tidak ditemukan"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data meeting"})
    }

    // Check if user is the creator
    if existingMeeting.CreatedBy != userID {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Hanya pembuat meeting yang dapat mengubah"})
    }

    var updateData models.Meeting
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
    }

    // Prepare update document
    update := bson.M{
        "$set": bson.M{
            "updatedAt": time.Now(),
        },
    }

    // Update fields if provided
    if updateData.Title != "" {
        update["$set"].(bson.M)["title"] = updateData.Title
    }
    if updateData.Description != "" {
        update["$set"].(bson.M)["description"] = updateData.Description
    }
    if !updateData.StartTime.IsZero() {
        update["$set"].(bson.M)["startTime"] = updateData.StartTime
    }
    if updateData.Duration > 0 {
        update["$set"].(bson.M)["duration"] = updateData.Duration
    }
    if updateData.Participants != nil {
        update["$set"].(bson.M)["participants"] = updateData.Participants
    }
    update["$set"].(bson.M)["emotionTracking"] = updateData.EmotionTracking

    result, err := config.MeetingCollectionRef.UpdateOne(ctx, bson.M{"_id": meetingID}, update)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengupdate meeting"})
    }

    if result.MatchedCount == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Meeting tidak ditemukan"})
    }

    return c.JSON(fiber.Map{
        "message": "Meeting berhasil diupdate",
    })
}

// DeleteMeeting deletes a meeting
func DeleteMeeting(c *fiber.Ctx) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Get meeting ID from URL parameter
    meetingID, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID meeting tidak valid"})
    }

    // Get user ID from JWT token
    userID, err := utils.GetUserIDFromToken(c)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid"})
    }

    // Check if meeting exists and user is the creator
    var existingMeeting models.Meeting
    err = config.MeetingCollectionRef.FindOne(ctx, bson.M{"_id": meetingID}).Decode(&existingMeeting)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Meeting tidak ditemukan"})
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengambil data meeting"})
    }

    // Check if user is the creator
    if existingMeeting.CreatedBy != userID {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Hanya pembuat meeting yang dapat menghapus"})
    }

    result, err := config.MeetingCollectionRef.DeleteOne(ctx, bson.M{"_id": meetingID})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal menghapus meeting"})
    }

    if result.DeletedCount == 0 {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Meeting tidak ditemukan"})
    }

    return c.JSON(fiber.Map{
        "message": "Meeting berhasil dihapus",
    })
}
