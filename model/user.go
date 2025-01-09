package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12
	Active       string = "active"
)

func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
