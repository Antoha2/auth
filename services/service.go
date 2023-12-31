package services

import (
	"context"
	"errors"

	AuthRepository "github.com/Antoha2/auth/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type AuthService interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
}

type servImpl struct {
	//log *slog.Logger
	AuthService
	rep AuthRepository.RepAuth
}
