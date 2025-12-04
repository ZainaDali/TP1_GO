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
		// Si le fichier n'existe pas encore, on considère qu'il n'y a aucun contact
		if os.IsNotExist(err) {
			return []*Contact{}, nil
		}
		return nil, fmt.Errorf("impossible de lire le fichier %s : %w", j.FilePath, err)
	}

	if len(data) == 0 {
		return []*Contact{}, nil
	}

	var contacts []*Contact
	if err := json.Unmarshal(data, &contacts); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage JSON : %w", err)
	}
	return contacts, nil
}

func (j *JSONStore) Add(contact *Contact) error {
	contacts, err := j.GetAll()
	if err != nil {
		return err
	}

	contact.ID = uint(len(contacts) + 1)
	contacts = append(contacts, contact)

	return j.saveAll(contacts)
}

func (j *JSONStore) Update(id uint, newName, newEmail string) error {
	contacts, err := j.GetAll()
	if err != nil {
		return err
	}

	var found bool
	for _, c := range contacts {
		if c.ID == id {
			if newName != "" {
				c.Name = newName
			}
			if newEmail != "" {
				c.Email = newEmail
			}
			found = true
			break
		}
	}

	if !found {
		return ErrContactNotFound
	}

	return j.saveAll(contacts)
}

func (j *JSONStore) Delete(id uint) error {
	contacts, err := j.GetAll()
	if err != nil {
		return err
	}

	var newContacts []*Contact
	for _, c := range contacts {
		if c.ID == id {
			continue
		}
		newContacts = append(newContacts, c)
	}

	return j.saveAll(newContacts)
}

func (j *JSONStore) GetUserById(id uint) (*Contact, error) {
	contacts, err := j.GetAll()
	if err != nil {
		return nil, err
	}

	for _, c := range contacts {
		if c.ID == id {
			return c, nil
		}
	}

	return nil, ErrContactNotFound
}

// saveAll sérialise et réécrit la liste complète des contacts dans le fichier JSON.
func (j *JSONStore) saveAll(contacts []*Contact) error {
	data, err := json.MarshalIndent(contacts, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON : %w", err)
	}

	if err := os.WriteFile(j.FilePath, data, 0644); err != nil {
		return fmt.Errorf("impossible d'écrire dans le fichier %s : %w", j.FilePath, err)
	}

	return nil
}
