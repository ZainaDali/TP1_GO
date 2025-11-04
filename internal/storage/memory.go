package storage

type MemoryStore struct {
	contacts []*Contact
	nextID   uint
}

func (m *MemoryStore) GetAll() ([]*Contact, error) {
	return m.contacts, nil
}

func (m *MemoryStore) GetUserById(id string) error {
	return nil
}
