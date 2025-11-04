package menu

import (
	"fmt"
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
