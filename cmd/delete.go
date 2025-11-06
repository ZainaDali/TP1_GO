package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteID uint

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact",
	Long:  "Supprime un contact en fonction de son ID",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Vérifier que l'ID est fourni (déjà géré par MarkFlagRequired)

		// 2. Optionnel : Afficher le contact avant suppression
		contact, err := store.GetUserById(deleteID)
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		fmt.Println("\nContact à supprimer :")
		displayContact(contact)

		// 3. Effectuer la suppression
		err = store.Delete(deleteID)
		if err != nil {
			fmt.Println("Erreur lors de la suppression:", err)
			return
		}

		// 4. Confirmation
		fmt.Printf("\nContact ID %d supprimé avec succès.\n", deleteID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintVarP(&deleteID, "id", "i", 0, "ID du contact à supprimer")
	deleteCmd.MarkFlagRequired("id")
}
