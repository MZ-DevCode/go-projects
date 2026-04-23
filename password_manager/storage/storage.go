package storage

import (
	"os"
	"fmt"
)

func Save(data []byte){
	err := os.WriteFile("vault.db", data, 0600)
	if err != nil {
		fmt.Println("Failed to save vault:", err)
	}
}

func Load() ([]byte, error) {
	data, err := os.ReadFile("vault.db")
	if err != nil {
		return nil, err
	}
	return data, nil
}
