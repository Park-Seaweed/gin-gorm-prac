package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"uniqueIndex;size:255"`
	Password     string `json:"-"`
	RefreshToken string `json:"-"`

	Posts []Post `json:"posts,omitempty"`
}
