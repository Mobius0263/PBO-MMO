package main

import (
    "log"
    "os"

    "pbommo/config"
    "pbommo/routes"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ No .env file found, using environment variables")
    }

    // Set default port if not provided
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Create Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            // Log detailed error
            log.Printf("❌ ERROR: %v\nPath: %s, Method: %s", err, c.Path(), c.Method())

            // Return error response
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }

            return c.Status(code).JSON(fiber.Map{
                "error":  err.Error(),
                "path":   c.Path(),
                "method": c.Method(),
            })
        },
    })

    // Configure CORS
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:5173",
        AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
        AllowCredentials: true,
        ExposeHeaders:    "Content-Length, Content-Disposition",
    }))

    // Add logger middleware
    app.Use(logger.New())

    // Serve static files
    app.Static("/uploads", "./uploads")

    // Connect to database
    config.ConnectDB()
    log.Println("✅ Connected to database")

    // Setup routes
    routes.SetupRoutes(app)

    // Start server
    log.Printf("🚀 Server running at http://localhost:%s", port)
    if err := app.Listen(":" + port); err != nil {
        log.Fatal("❌ Failed to start server:", err)
    }
}