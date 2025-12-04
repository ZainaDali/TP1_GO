package storage

import "fmt"

type Storer interface {
	GetAll() ([]*Contact, error)
	Add(c *Contact) error
	Update(id uint, newName, newMail string) error
	Delete(id uint) error
	GetUserById(id uint) (*Contact, error)
}

type Contact struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ErrContactNotFound = fmt.Errorf("L'utilisateur est introuvable.")
