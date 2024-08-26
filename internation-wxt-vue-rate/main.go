package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Message struct {
	ExtName struct {
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"extName"`
	Description struct {
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"description"`
	ActionTitle struct {
		Message     string `json:"message"`
		Description string `json:"description"`
	} `json:"actionTitle"`
}

type DescriptionUpdates struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

func main() {
	updates, err := loadDescriptionUpdates("./updates.json")
	if err != nil {
		fmt.Println("Error loading updates:", err)
		return
	}

	localesPath := "./public/_locales"
	err = filepath.WalkDir(localesPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("err1: %v\n", err)
			return err
		}
		if d.IsDir() || filepath.Base(path) != "messages.json" {
			fmt.Println("messages.json bak")
			return nil
		}

		langCode := filepath.Base(filepath.Dir(path))
		update, exists := updates[langCode]
		if !exists {
			fmt.Println("exists")
			return nil
		}

		err = updateMessageFile(path, update.Text)
		if err != nil {
			fmt.Println("Error updating file:", path, err)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through locales:", err)
	}
}

func loadDescriptionUpdates(filePath string) (map[string]DescriptionUpdates, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("err3: %v\n", err)
		return nil, err
	}

	var updates map[string]DescriptionUpdates
	err = json.Unmarshal(data, &updates)
	if err != nil {
		fmt.Printf("err5: %v\n", err)
		return nil, err
	}

	return updates, nil
}

func updateMessageFile(filePath, newText string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("err7: %v\n", err)
		return err
	}

	var message Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		fmt.Printf("err9: %v\n", err)
		return err
	}

	message.Description.Message = newText

	updatedData, err := json.MarshalIndent(message, "", "    ")
	if err != nil {
		fmt.Printf("err11: %v\n", err)
		return err
	}

	return os.WriteFile(filePath, updatedData, 0644)
}
