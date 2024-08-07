package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "strpass/src/handler"
    "strpass/src/model"
    "strpass/src/repository"
    "strpass/src/service"
)

func main() {
    router := gin.Default()

    db, err := model.SetupDB()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    passwordRepo := repository.NewPasswordRepository(db)
    passwordService := service.NewPasswordService(passwordRepo)
    passwordHandler := handler.NewPasswordHandler(passwordService)

    router.POST("/api/strong_password_steps", passwordHandler.GetStrongPasswordSteps)

    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
