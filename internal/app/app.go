package app

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
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
		DisplayMenu()
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

func DisplayMenu() {
	fmt.Println("\n=== Menu ===")
	fmt.Println("1. Ajouter un contact")
	fmt.Println("2. Voir les contacts")
	fmt.Println("3. Supprimer un contact")
	fmt.Println("4. Mettre à jour un contact")
	fmt.Println("5. Quitter")
	fmt.Print("Choisissez une option (1-5): ")
}

func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readID(reader *bufio.Reader, prompt string) (uint, error) {
	input := readInput(reader, prompt)
	id64, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("ID invalide")
	}
	return uint(id64), nil
}
func handleAddContact(reader *bufio.Reader, store storage.Storer) {

	fmt.Println("\n Ajouter un contact")
	fmt.Println("====================")

	name := readInput(reader, "Nom : ")
	email := readInput(reader, "Email : ")

	if name == "" || email == "" {
		fmt.Println("Erreur: nom et email requis")
		return
	}

	contact := &storage.Contact{
		Name:  name,
		Email: email,
	}

	err := store.Add(contact)
	if err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
		return
	}

	displayContact(contact)
	fmt.Println("Contact ajouté avec succès!")
}

func handleListContacts(store storage.Storer) {
	fmt.Println("\nListe des contacts :")
	fmt.Println("========================")

	// Récupère tous les contacts du store
	contacts, err := store.GetAll()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des contacts :", err)
		return
	}

	if len(contacts) == 0 {
		fmt.Println("Aucun contact de disponible.")
		return
	}

	for _, contact := range contacts {
		displayContact(contact)
	}

	fmt.Printf("\n Total : %d contact(s)\n", len(contacts))
}

func handleRemoveContact(reader *bufio.Reader, store storage.Storer) {
	id, err := readID(reader, "ID du contact à supprimer : ")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	err = store.Delete(id)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Printf("Le contact avec l'ID %d a été supprimé avec succès.\n", id)
}

func handleUpdateContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Println("\nMettre à jour un contact")
	fmt.Println("============================")

	// Affiche liste
	handleListContacts(store)

	id, err := readID(reader, "\nID du contact à modifier : ")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	contact, err := store.GetUserById(id)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("\n Informations actuelles :")
	fmt.Printf("   Nom  : %s\n", contact.Name)
	fmt.Printf("   Email: %s\n", contact.Email)

	newName := readInput(reader, "\nNouveau nom (laisser vide pour ne pas changer) : ")
	newEmail := readInput(reader, "Nouveau email (laisser vide pour ne pas changer) : ")

	// Appel au store (au lieu de contact.update())
	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("\n Contact mis à jour avec succès!")
}

func handleAddContactCLI(store storage.Storer, name, email string) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		fmt.Println("Erreur: nom et email requis")
		return
	}

	contact := &storage.Contact{
		Name:  name,
		Email: email,
	}

	err := store.Add(contact)
	if err != nil {
		fmt.Println("Erreur lors de la création :", err)
		return
	}

	displayContact(contact)
	fmt.Println(" Contact ajouté via CLI!")
}

func displayContact(c *storage.Contact) {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│   ID    : %-28d │\n", c.ID)
	fmt.Printf("│   Nom   : %-28s │\n", c.Name)
	fmt.Printf("│   Email : %-28s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}
