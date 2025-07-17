package models

import (
    "time"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Nama         string             `bson:"nama" json:"nama"`
    Email        string             `bson:"email" json:"email"`
    Password     string             `bson:"password" json:"-"` // Password tidak di-expose ke JSON
    Role         string             `bson:"role" json:"role,omitempty"`
    Bio          string             `bson:"bio" json:"bio,omitempty"`
    ProfileImage string             `bson:"profileImage" json:"profileImage,omitempty"`
    CreatedAt    time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
    UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt,omitempty"`
}

type UserResponse struct {
    ID           primitive.ObjectID `json:"id,omitempty"`
    Nama         string             `json:"nama"`
    Email        string             `json:"email"`
    Role         string             `json:"role,omitempty"`
    Bio          string             `json:"bio,omitempty"`
    ProfileImage string             `json:"profileImage,omitempty"`
    Status       string             `json:"status,omitempty"`
    LastActive   time.Time          `json:"lastActive,omitempty"`
}

type Meeting struct {
    ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
    Title       string               `bson:"title" json:"title"`
    Description string               `bson:"description" json:"description,omitempty"`
    StartTime   time.Time            `bson:"startTime" json:"startTime"`
    Duration    int                  `bson:"duration" json:"duration"` // in minutes
    CreatedBy   primitive.ObjectID   `bson:"createdBy" json:"createdBy"`
    Participants []primitive.ObjectID `bson:"participants" json:"participants"`
    EmotionTracking bool             `bson:"emotionTracking" json:"emotionTracking"`
    CreatedAt    time.Time           `bson:"createdAt" json:"createdAt"`
    UpdatedAt    time.Time           `bson:"updatedAt" json:"updatedAt"`
}