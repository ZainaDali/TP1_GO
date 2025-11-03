package menu

import (
	"bufio"
	"fmt"
	"strings"
)

func DisplayMenu() {
	fmt.Println("\n=== Menu ===")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Voir les contacts")
	fmt.Println("3. Supprimer un contact")
	fmt.Println("4. Mettre à jour un contact")
	fmt.Println("5. Quitter")
	fmt.Print("Choisissez une option (1-5): ")
}

func displaySubMenu(title string) {
	fmt.Printf("\n=== %s ===\n", title)
	fmt.Println("Appuyez sur 'r' pour retourner au menu principal")
}

func HandleSubMenu(reader *bufio.Reader, title string) {
	displaySubMenu(title)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "r" {
			return
		}
		fmt.Println("Appuyez sur 'r' pour retourner au menu principal")
	}
}
