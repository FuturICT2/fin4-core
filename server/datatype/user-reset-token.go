package datatype

import (
	"math/rand"
	"time"
)

// UserResetToken reset token struct
type UserResetToken struct {
	ID        ID
	UserID    ID
	Token     string
	CreatedAt time.Time
}

// NewUserResetToken creates a new reset password token
func NewUserResetToken(userID ID) *UserResetToken {
	return &UserResetToken{
		UserID:    userID,
		Token:     generateRandomString(64),
		CreatedAt: time.Now(),
	}
}

// generates new random token
func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	space := "abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLNOPQRSTUVWXYZ0123456789_$!*-"
	spaceLength := len(space)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = space[rand.Intn(spaceLength)]
	}
	return string(result)
}
