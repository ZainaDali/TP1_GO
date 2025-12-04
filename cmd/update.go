package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	updateID    uint
	updateName  string
	updateEmail string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact",
	Long:  "Modifie le nom et/ou l'email d'un contact existant",
	Run: func(cmd *cobra.Command, args []string) {

		if updateName == "" && updateEmail == "" {
			fmt.Println("Aucune modification spécifiée. Utilisez --name ou --email")
			return
		}

		if err := store.Update(updateID, updateName, updateEmail); err != nil {
			fmt.Println("Erreur lors de la mise à jour:", err)
			return
		}

		fmt.Printf("\nContact n°%d mis à jour avec succès!", updateID)
		updatedContact, _ := store.GetUserById(updateID)
		displayContact(updatedContact)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().UintVarP(&updateID, "id", "i", 0, "ID du contact à modifier")
	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Nouveau nom (optionnel)")
	updateCmd.Flags().StringVarP(&updateEmail, "email", "e", "", "Nouvel email (optionnel)")

	updateCmd.MarkFlagRequired("id")
}
