package repositories

import (
	"fmt"
	"user/models"

	"gorm.io/gorm"
)

func FindById(db *gorm.DB, id int) (models.User, error) {
	var user models.User

	result := db.First(&user, id)
	if result.Error != nil {
		return user, fmt.Errorf("user does not exist...")
	}

	return user, nil
}
