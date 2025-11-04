package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ZainaDali/TP1_GO.git/contact"
	"github.com/ZainaDali/TP1_GO.git/menu"
)

func main() {

	name := flag.String("name", "", "Nom du contact à ajouter")
	email := flag.String("email", "", "Email du contact à ajouter")
	flag.Parse()
	if *name != "" && *email != "" {
		contact.AddContactCLI(*name, *email)
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
			contact.AddContact()
		case "2":
			contact.ListContacts()
		case "3":
			contact.RemoveContact()
		case "4":
			contact.UpdateContact()
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir entre 1 et 5.")
		}
	}
}
