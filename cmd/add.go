package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/spf13/cobra"
)

var name string
var email string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un contact",
	Run: func(cmd *cobra.Command, args []string) {
		// si l'user utilise les flags
		if name != "" && email != "" {
			addContact(name, email)
			return
		}

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Nom: ")
		n, _ := reader.ReadString('\n')
		n = strings.TrimSpace(n)

		fmt.Print("Email: ")
		m, _ := reader.ReadString('\n')
		m = strings.TrimSpace(m)

		if n == "" || m == "" {
			fmt.Println("Erreur : nom et email requis")
			return
		}

		addContact(n, m)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
}

func addContact(n, m string) {
	c := &storage.Contact{Name: n, Email: m}
	if err := store.Add(c); err != nil {
		fmt.Println("Erreur lors de l'ajout :", err)
		return
	}
	fmt.Printf("Contact n°%d ajouté avec succès !\n", c.ID)
}
