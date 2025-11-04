package storage

type Storer interface {
	GetAll(c []*Contact) error
	Add(c *Contact) error
	Update(id, newName, newMail string) error
	Delete(c *Contact) error
	GetUserById(id string) error
}

type Contact struct {
	ID    uint
	Name  string
	Email string
}
