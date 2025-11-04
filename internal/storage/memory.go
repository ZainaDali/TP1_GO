package storage

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

func (m *MemoryStore) GetUserById(id string) error {
	return nil
}
