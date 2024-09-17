package model

import (
	"ekolo/pkg/storage"
	"ekolo/pkg/xlog"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Organization is the organization model
type Organization struct {
	storage.BaseModel
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Phone *string `json:"phone"`
}

// User is the user model
type User struct {
	storage.BaseModel
	Email      string       `json:"email" gorm:"primaryKey;not null"`
	Password   *string      `json:"password" `
	FirstName  *string      `json:"first_name" `
	LastName   *string      `json:"last_name" `
	BirthDate  *string      `json:"birth_date" `
	BirthPlace *string      `json:"birth_place" `
	Address    *string      `json:"address" `
	Phone      *string      `json:"phone" `
	Type       *string      `json:"type" `
	OrgUUID    uuid.UUID    `json:"org"`
	Org        Organization `json:"-"`
}

func (u *User) SetPassword(password string) error {
	hashBytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		xlog.Error("user", "set-password", err)
		return err
	}
	hashPassword := string(hashBytePassword)
	u.Password = &hashPassword
	return nil
}

func (u User) Authenticate(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(*u.Password), []byte(password)); err != nil {
		xlog.Error("user", "authenticate", err)
		return err
	}
	return nil
}

func GetModels() []any {
	return []any{
		Organization{}, User{},
	}
}
