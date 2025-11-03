package contact

import "fmt"

type Contact struct {
	ID    uint
	Name  string
	Email string
}

var (
	contacts      = make(map[uint]Contact)
	nextID   uint = 1
)

func AddContact(name, email string) {
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	contacts[nextID] = newContact

	nextID++
}

func RemoveContact(id uint) string {
	return fmt.Sprintf("Le contact avec l'ID : %d est supprimé.", id)
}
