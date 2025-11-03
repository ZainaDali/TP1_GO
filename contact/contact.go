package contact

import (
	"bufio"
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
	contacts = map[uint]Contact{
		1: {
			ID:    1,
			Name:  "alexe",
			Email: "alexe@exemple.fr",
		},
	}
	nextID uint = 2
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

	fmt.Println("✅ Contact ajouté avec succès!")
	// Afficher le contact qui vient d'être créé
	DisplayContact(newContact)
}

// DisplayContact affiche un seul contact de manière formatée
func DisplayContact(c Contact) {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│ 🆔 ID    : %-27d │\n", c.ID)
	fmt.Printf("│ 👤 Nom   : %-27s │\n", c.Name)
	fmt.Printf("│ 📧 Email : %-27s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}

func ListContacts() {
	fmt.Println("\n📋 Liste des contacts :")
	fmt.Println("========================")

	for _, contact := range contacts {
		DisplayContact(contact)
	}

	fmt.Printf("\n📊 Total : %d contact(s)\n", len(contacts))
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

	if _, exists := contacts[id]; !exists {
		fmt.Printf("Erreur : aucun contact trouvé avec l'ID %d.\n", id)
		return
	}

	delete(contacts, id)
	fmt.Printf("Le contact avec l'ID %d a été supprimé avec succès.\n", id)
}
