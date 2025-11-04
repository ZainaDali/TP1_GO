package contact

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	ID    uint
	Name  string
	Email string
}

var (
	contacts = map[uint]*Contact{
		1: {
			ID:    1,
			Name:  "alexe",
			Email: "alexe@exemple.fr",
		},
	}
	nextID uint = 2
)

func newContact(name, email string) (*Contact, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, errors.New("Le nom doit être renseigné.")
	}
	if email == "" {
		return nil, errors.New("Le mail doit être renseigné.")
	}

	c := &Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
	}
	return c, nil
}

func (c *Contact) add() {
	contacts[c.ID] = c
	nextID++
}

func AddContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n Ajouter un contact")
	fmt.Println("====================")

	fmt.Print("Nom : ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	c, err := newContact(name, email)

	if err != nil {
		fmt.Println("Erreur lors de la création :", err)
		return
	}

	c.add()
	c.displayContact()

	fmt.Println("Contact ajouté avec succès!")
}

func (c Contact) displayContact() {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│   ID    : %-28d │\n", c.ID)
	fmt.Printf("│   Nom   : %-28s │\n", c.Name)
	fmt.Printf("│   Email : %-28s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}

func ListContacts() {
	fmt.Println("\nListe des contacts :")
	fmt.Println("========================")

	if len(contacts) == 0 {
		fmt.Printf("Aucun contact de disponible.\n")
		return
	}

	for _, contact := range contacts {
		contact.displayContact()
	}

	fmt.Printf("\n Total : %d contact(s)\n", len(contacts))
}

func (c *Contact) Remove() {
	delete(contacts, c.ID)
	c = nil
}

func RemoveContact() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID du contact à supprimer : ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println("Erreur : ID invalide.")
		return
	}
	id := uint(id64)

	contact, exists := contacts[id]
	if !exists {
		fmt.Printf("Erreur : aucun contact trouvé avec l'ID %d.\n", id)
		return
	}

	contact.Remove()
	fmt.Printf("Le contact avec l'ID %d a été supprimé avec succès.\n", id)
}

func (c *Contact) Update(newName, newEmail string) {
	if newName != "" {
		c.Name = newName
	}
	if newEmail != "" {
		c.Email = newEmail
	}
}

func UpdateContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nMettre à jour un contact")
	fmt.Println("============================")

	ListContacts()

	fmt.Print("\nID du contact à modifier : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	id64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		fmt.Println("Erreur : ID invalide.")
		return
	}
	id := uint(id64)

	contact, exists := contacts[id]
	if !exists {
		fmt.Printf("Erreur : aucun contact trouvé avec l'ID %d.\n", id)
		return
	}

	fmt.Println("\n Informations actuelles :")
	fmt.Printf("   Nom  : %s\n", contact.Name)
	fmt.Printf("   Email: %s\n", contact.Email)

	fmt.Print("\nNouveau nom (laisser vide pour ne pas changer) : ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)

	fmt.Print("Nouveau email (laisser vide pour ne pas changer) : ")
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	contact.Update(newName, newEmail)

	fmt.Println("\n Contact mis à jour avec succès!")
	fmt.Println("\n Nouvelles informations :")
	fmt.Printf("   ID   : %d\n", contact.ID)
	fmt.Printf("   Nom  : %s\n", contact.Name)
	fmt.Printf("   Email: %s\n", contact.Email)
}

func AddContactCLI(name, email string) {
	c, err := newContact(name, email)

	if err != nil {
		fmt.Println("Erreur lors de la création :", err)
		return
	}

	c.add()
	c.displayContact()

	fmt.Println(" Contact ajouté via CLI!")
}
