package storage

import (
	"fmt"
)

var ErrContactNotFound = fmt.Errorf("L'utilisateur est introuvable.")

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
	if c, ok := m.contacts[id]; ok {
		return c, nil
	}
	return nil, ErrContactNotFound
}

func (m *MemoryStore) Add(contact *Contact) error {
	contact.ID = m.nextID
	m.contacts[contact.ID] = contact
	m.nextID++
	return nil
}

func (m *MemoryStore) Update(id uint, newName, newMail string) error {
	contact, exists := m.contacts[id]
	if !exists {
		return ErrContactNotFound
	}

	if newName != "" {
		contact.Name = newName
	}
	if newMail != "" {
		contact.Email = newMail
	}

	return nil
}

func (m *MemoryStore) Delete(id uint) error {
	if _, exists := m.contacts[id]; !exists {
		return ErrContactNotFound
	}

	delete(m.contacts, id)
	return nil
}
