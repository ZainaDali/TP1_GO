package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONStore struct {
	FilePath string
	nextID   uint
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{FilePath: path}
}

func (j *JSONStore) GetAll() ([]*Contact, error) {
	data, err := os.ReadFile(j.FilePath)
	if err != nil {
		return nil, fmt.Errorf("Impossible to read file %s : %w", j.FilePath, err)
	}

	var contacts []*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (j *JSONStore) Add(contact *Contact) error {
	return nil
}

func (j *JSONStore) Update(id uint, newName, newEmail string) error {
	return nil
}

func (j *JSONStore) Delete(id uint) error {
	return nil
}

func (j *JSONStore) GetUserById(id uint) (*Contact, error) {
	return nil, nil
}
