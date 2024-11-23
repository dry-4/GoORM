package services

import (
	"fmt"
	"user/models"
	"user/repositories"

	"gorm.io/gorm"
)

func CreateUserService(user models.User, db *gorm.DB) (models.User, error) {
	if validationErr := user.Validate(); validationErr != nil {
		return user, validationErr
	}

	result := db.Create(&user)
	if result.Error != nil {
		return user, fmt.Errorf("user can not be created...")
	}

	return user, nil
}

func GetUsersService(db *gorm.DB) ([]models.User, error) {
	var users []models.User

	res := db.Find(&users)
	if res.Error != nil {
		return nil, fmt.Errorf("user can not be created...")
	}

	return users, nil
}

func GetUserByIdsService(db *gorm.DB, id int) (models.User, error) {
	var user models.User

	user, err := repositories.FindById(db, id)

	if err != nil {
		return user, fmt.Errorf("user not found...")
	}

	return user, nil
}

func UpdateUserById(db *gorm.DB, id int, user models.User) (models.User, error) {
	var upatedUser models.User

	res := db.First(&upatedUser, id)
	if res.Error != nil {
		return user, fmt.Errorf("user not found...")
	}

	res = db.Model(&upatedUser).Updates(user)
	if res.Error != nil {
		return upatedUser, fmt.Errorf("user updation failed...")
	}

	return upatedUser, nil
}

func DeleteUserById(db *gorm.DB, id int) error {
	res := db.Delete(&models.User{}, id)

	if res.Error != nil {
		return fmt.Errorf("user delation fails...")
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("User not found...")
	}

	return nil
}
