package app

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/ZainaDali/TP1_GO.git/menu"
)

func Run(store storage.Storer) {
	name := flag.String("name", "", "Nom du contact à ajouter")
	email := flag.String("email", "", "Email du contact à ajouter")
	flag.Parse()

	if *name != "" && *email != "" {
		handleAddContactCLI(store, *name, *email)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var choice string

	for {
		menu.DisplayMenu()
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleAddContact(reader, store)
		case "2":
			handleListContacts(store)
		case "3":
			handleRemoveContact(reader, store)
		case "4":
			handleUpdateContact(reader, store)
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir entre 1 et 5.")
		}
	}
}

func handleAddContact(reader *bufio.Reader, store storage.Storer) {
	// Votre logique ici
}

func handleListContacts(store storage.Storer) {
	// Votre logique ici
}

func handleRemoveContact(reader *bufio.Reader, store storage.Storer) {
	// Votre logique ici
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	// Votre logique ici
}

func handleAddContactCLI(store storage.Storer, name, email string) {
	// Votre logique ici
}
