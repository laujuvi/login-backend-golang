package user

import "gorm.io/gorm"

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	return user, result.Error
}
