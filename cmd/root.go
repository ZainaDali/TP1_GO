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
}
