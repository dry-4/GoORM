package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique; not null"`
	Age   int    `json:"age"`
}

func (u *User) Validate() error {
	if strings.Trim(u.Name, " ") == "" {
		return fmt.Errorf("Invalid user name..")
	}

	if strings.Trim(u.Email, " ") == "" {
		return fmt.Errorf("Invalid user email..")
	}

	return nil
}
