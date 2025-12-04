package storage

import (
	"fmt"

	"gorm.io/gorm"
)

type Storer interface {
	GetAll() ([]*Contact, error)
	Add(c *Contact) error
	Update(id uint, newName, newMail string) error
	Delete(id uint) error
	GetUserById(id uint) (*Contact, error)
}

type Contact struct {
	gorm.Model
	// ID    uint   `json:"id"`
	// Name  string `json:"name"`
	// Email string `json:"email"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);unique;not null"`
}

var ErrContactNotFound = fmt.Errorf("L'utilisateur est introuvable.")
