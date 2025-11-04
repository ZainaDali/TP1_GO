package storage

type Storer interface {
	GetAll() ([]*Contact, error)
	Add(c *Contact) error
	Update(id uint, newName, newMail string) error
	Delete(id uint) error
	GetUserById(id uint) (*Contact, error)
}

type Contact struct {
	ID    uint
	Name  string
	Email string
}
