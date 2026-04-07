package main

import (
	"fmt"

	"github.com/PPlaner/Backend/internal/config"
	"github.com/PPlaner/Backend/internal/database"
)

func main() {
	fmt.Println("Backend is running")

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := database.Connect(cfg.DB)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Connected to database")

}
