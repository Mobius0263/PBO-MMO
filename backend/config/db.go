package config

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    MongoClient       *mongo.Client
    UserCollectionRef *mongo.Collection
)

func ConnectDB() {
    // Load env variables
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found")
    }

    // Set up connection context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Set connection string
    mongoURI := os.Getenv("MONGOSTRING")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017"
        log.Println("Using default MongoDB connection string")
    }

    // Connect to MongoDB
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("Failed to connect to MongoDB:", err)
    }

    // Check the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Failed to ping MongoDB:", err)
    }

    // Set global connection
    MongoClient = client

    // Initialize collection references
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "dbPBOMMO"
        log.Println("Using default database name:", dbName)
    }

    UserCollectionRef = MongoClient.Database(dbName).Collection(os.Getenv("USER_COLLECTION"))
    if UserCollectionRef == nil {
        log.Println("Warning: User collection reference is nil")
        UserCollectionRef = MongoClient.Database(dbName).Collection("users")
    }

    log.Println("Connected to MongoDB")
}

// Tambahkan di dalam package config

func GetDbName() string {
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "dbPBOMMO"
    }
    return dbName
}