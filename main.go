package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
			menu.HandleSubMenu(reader, "Ajouter un contact")
		case "2":
			menu.HandleSubMenu(reader, "Voir les contacts")
		case "3":
			menu.HandleSubMenu(reader, "Supprimer un contact")
		case "4":
			menu.HandleSubMenu(reader, "Mettre à jour un contact")
		case "5":
			fmt.Println("Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuillez choisir entre 1 et 5.")
		}
	}
}
