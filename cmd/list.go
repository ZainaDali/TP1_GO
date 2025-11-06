package cmd

import (
	"fmt"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func displayContact(c *storage.Contact) {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│   ID    : %-28d │\n", c.ID)
	fmt.Printf("│   Nom   : %-28s │\n", c.Name)
	fmt.Printf("│   Email : %-28s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}
