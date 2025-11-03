package contact

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	ID    uint
	Name  string
	Email string
}

var (
	contacts      = make(map[uint]Contact)
	nextID   uint = 1
)

func AddContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n➕ Ajouter un contact")
	fmt.Println("====================")

	fmt.Print("Nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	contacts[nextID] = newContact
	nextID++

	fmt.Println("Contact ajouté avec succès!")
}

func AddContactDirect(name, email string) {
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
	}

	contacts[nextID] = newContact
	nextID++
}

func ListContacts() {
	if len(contacts) == 0 {

		newContact := Contact{
			ID:    nextID,
			Name:  "alexe",
			Email: "alexe@exemple.fr",
		}

		contacts[nextID] = newContact
		nextID++
	}

	fmt.Println("\n📋 Liste des contacts :")
	fmt.Println("========================")

	for id, contact := range contacts {
		fmt.Printf("ID: %d | Nom: %-20s | Email: %s\n",
			id, contact.Name, contact.Email)
	}

	fmt.Printf("\n Total : %d contact(s)\n", len(contacts))
}
