package auth

import (
	"errors"
	"os"
	"time"

	"github.com/laujuvi/login-system/internal/database"
	"github.com/laujuvi/login-system/internal/user"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("")

func Init() {
	jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
}

func LoginUser(input LoginRequest) (LoginResponse, error) {
	u, err := user.GetUserByEmail(database.DB, input.Email)
	if err != nil {
		return LoginResponse{}, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return LoginResponse{}, errors.New("invalid email or password")
	}

	accessToken, err := GenerateToken(u.ID, 15*time.Minute)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, err := GenerateRefreshToken(u.ID, 7*24*time.Hour)
	if err != nil {
		return LoginResponse{}, err
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
