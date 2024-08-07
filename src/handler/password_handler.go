package handler

import (
    "strpass/src/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

type PasswordHandler struct {
    PasswordService *service.PasswordService
}

func NewPasswordHandler(passwordService *service.PasswordService) *PasswordHandler {
    return &PasswordHandler{PasswordService: passwordService}
}

func (h *PasswordHandler) GetStrongPasswordSteps(c *gin.Context) {
    var request struct {
        InitPassword string `json:"init_password"`
    }

    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    steps := h.PasswordService.GetStrongPasswordSteps(request.InitPassword)
    c.JSON(http.StatusOK, gin.H{"num_of_steps": steps})
}
