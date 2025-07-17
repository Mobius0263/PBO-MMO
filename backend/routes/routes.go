package routes

import (
    "pbommo/controllers"
    "pbommo/middleware"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    // Health check
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("API Server is running")
    })

    // Auth routes (unprotected)
    app.Post("/register", controllers.Register)
    app.Post("/login", controllers.Login)

    // Get all users (unprotected for demo purposes)
    app.Get("/users", controllers.GetUsers)

    // API routes (protected)
    api := app.Group("/api")
    api.Use(middleware.Protected())

    // User routes
    api.Get("/users/:id", controllers.GetUserById)
    api.Put("/users/:id", controllers.UpdateUser)
    api.Post("/users/:id/profile-image", controllers.UploadProfileImage)

    // Meeting routes
    api.Post("/meetings", controllers.CreateMeeting)
    api.Get("/meetings", controllers.GetMeetings)
    api.Get("/meetings/:id", controllers.GetMeetingById)
    api.Put("/meetings/:id", controllers.UpdateMeeting)
    api.Delete("/meetings/:id", controllers.DeleteMeeting)
}