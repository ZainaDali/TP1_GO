package cmd

import (
	"fmt"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/spf13/cobra"
)

var name string
var email string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un contact",
	Run: func(cmd *cobra.Command, args []string) {
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
		fmt.Printf("Contact n°%d ajouté avec succès!", contact.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
}
