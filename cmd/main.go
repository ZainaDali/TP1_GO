package main

import (
	"github.com/ZainaDali/TP1_GO.git/internal/app"
	"github.com/ZainaDali/TP1_GO.git/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	app.Run(store)
}
