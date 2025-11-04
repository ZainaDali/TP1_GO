package storage

import (
	"fmt"
)

type MemoryStore struct {
	contacts map[uint]*Contact
	nextID   uint
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[uint]*Contact),
		nextID:   1,
	}
}

func (m *MemoryStore) GetAll() ([]*Contact, error) {
	contacts := make([]*Contact, 0, len(m.contacts))
	for _, contact := range m.contacts {
		contacts = append(contacts, contact)
	}
	return contacts, nil
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

func (m *MemoryStore) Add(contact *Contact) error {
	contact.ID = m.nextID
	m.contacts[contact.ID] = contact
	m.nextID++
	return nil
}

func (m *MemoryStore) Update(id, newName, newMail string) error {
	return nil
}

func (m *MemoryStore) Delete(id uint) error {
	if _, exists := m.contacts[id]; !exists {
		return fmt.Errorf("contact not found")
	}

	delete(m.contacts, id)
	return nil
}
