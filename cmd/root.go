package cmd

import (
	"fmt"
	"os"

	"github.com/ZainaDali/TP1_GO.git/internal/config"
	"github.com/ZainaDali/TP1_GO.git/internal/storage"
	"github.com/spf13/cobra"
)

var store storage.Storer
var storageType string

var rootCmd = &cobra.Command{
	Use: "crm",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		switch storageType {
		case "json":
			store = storage.NewJSONStore("users.json")
		case "", "gorm":
			config.InitConfig()
			db, err := storage.ConnectDB()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			store = storage.NewGORMStore(db)
		default:
			fmt.Fprintf(os.Stderr, "Type de stockage inconnu '%s' (utiliser 'json' ou 'gorm')\n", storageType)
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	// par défaut, on utilise les données de la base de données SQLite
	rootCmd.PersistentFlags().StringVarP(&storageType, "storage", "s", "gorm", "Type de stockage (ex: gorm, json)")
}

func displayContact(c *storage.Contact) {
	fmt.Println("\n┌────────────────────────────────────────┐")
	fmt.Printf("│   ID    : %-28d │\n", c.ID)
	fmt.Printf("│   Nom   : %-28s │\n", c.Name)
	fmt.Printf("│   Email : %-28s │\n", c.Email)
	fmt.Println("└────────────────────────────────────────┘")
}
