package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ZainaDali/TP1_GO.git/contact"
	"github.com/ZainaDali/TP1_GO.git/menu"
)

func main() {
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
