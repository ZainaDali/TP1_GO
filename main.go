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
			fmt.Println("ajouter un contact fct")
		case "2":
			fmt.Println("voir les contacts fct")
		case "3":
			fmt.Println(contact.RemoveContact(1))
		case "4":
			fmt.Println("modifier un contact fct")
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir entre 1 et 5.")
		}
	}
}
