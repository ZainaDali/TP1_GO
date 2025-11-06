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
	store = storage.NewMemoryStore()
	testContacts := []*storage.Contact{
		{Name: "Alice Dupont", Email: "alice@mail.com"},
		{Name: "Bob Martin", Email: "bob@mail.com"},
		{Name: "Charlie Durand", Email: "charlie@mail.com"},
		{Name: "Diana Prince", Email: "diana@mail.com"},
		{Name: "Ethan Hunt", Email: "ethan@mail.com"},
	}

	for _, contact := range testContacts {
		store.Add(contact)
	}
}
