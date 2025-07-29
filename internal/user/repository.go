package user

import (
	"github.com/laujuvi/login-system/internal/user/model"
	"gorm.io/gorm"
)

func GetUserByEmail(db *gorm.DB, email string) (model.User, error) {
	var user model.User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}
