package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/Hoomanfr/messaging/mailbox/internal/common"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	UserName  string
	Password  string
	CreatedAt time.Time
}

func NewUser(userName string, password string) (*User, error) {
	if strings.TrimSpace(userName) == "" {
		return nil, errors.New(common.ErrInvalidUserName)
	}
	if strings.TrimSpace(password) == "" {
		return nil, errors.New(common.ErrInvalidPassword)
	}
	return &User{
		ID:        uuid.New().String(),
		UserName:  userName,
		Password:  password,
		CreatedAt: time.Now(),
	}, nil
}
