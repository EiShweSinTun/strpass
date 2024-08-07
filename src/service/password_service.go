package service

import (
    "strpass/src/repository"
    "unicode"
)

type PasswordService struct {
    PasswordRepository *repository.PasswordRepository
}

func NewPasswordService(passwordRepository *repository.PasswordRepository) *PasswordService {
    return &PasswordService{PasswordRepository: passwordRepository}
}

func (s *PasswordService) GetStrongPasswordSteps(password string) int {
    var steps int

    if len(password) < 6 {
        steps += 6 - len(password)
    } else if len(password) > 20 {
        steps += len(password) - 20
    }

    hasLower, hasUpper, hasDigit := false, false, false
    for _, char := range password {
        if unicode.IsLower(char) {
            hasLower = true
        } else if unicode.IsUpper(char) {
            hasUpper = true
        } else if unicode.IsDigit(char) {
            hasDigit = true
        }
    }

    if !hasLower {
        steps++
    }
    if !hasUpper {
        steps++
    }
    if !hasDigit {
        steps++
    }

    // Adjusted the logic for counting consecutive repeating characters
    for i := 2; i < len(password); i++ {
        if password[i] == password[i-1] && password[i] == password[i-2] {
            steps++
            i++
        }
    }

    return steps
}
