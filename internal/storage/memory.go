package storage

import (
	"fmt"
)

type MemoryStore struct {
	contacts []*Contact
	nextID   uint
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: []*Contact{},
		nextID:   1,
	}
}

func (m *MemoryStore) GetAll() ([]*Contact, error) {
	return m.contacts, nil
}

func (m *MemoryStore) GetUserById(id uint) (*Contact, error) {
	// ajoute une verif si l'id est au bon format

	for _, contact := range m.contacts {
		if contact.ID == id {
			return contact, nil
		}
	}

	return nil, fmt.Errorf("contact not found")
}

func (m *MemoryStore) Add(contact Contact) error {
	return nil
}

func (m *MemoryStore) Update(id, newName, newMail string) error {
	return nil
}

func (m *MemoryStore) Delete(id uint) error {
	contact, err := m.GetUserById(id)
	if err != nil {
		return err
	}

	for i, c := range m.contacts {
		if c == contact {
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("contact not found")
}
