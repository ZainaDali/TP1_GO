package cmd

import (
	"fmt"
	"os"

	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/spf13/cobra"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use: "crm",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	store = storage.NewJSONStore("users.json")
}

func displayContact(c *storage.Contact) {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│   ID    : %-28d │\n", c.ID)
	fmt.Printf("│   Nom   : %-28s │\n", c.Name)
	fmt.Printf("│   Email : %-28s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}
