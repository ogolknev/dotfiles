package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type StartListEntry struct {
	Name    string   `json:"name"`
	Command string   `json:"cmd"`
	Args    []string `json:"args,omitempty"`
}

type Config struct {
	StartList []StartListEntry `json:"start-list"`
}

func main() {
	startListJSON, err := os.Open(filepath.Join(filepath.Dir(os.Args[0]), "start-list.json"))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer startListJSON.Close()

	var config Config

	decoder := json.NewDecoder(startListJSON)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	startListMap := make(map[string]StartListEntry)
	entryNamesBuffer := bytes.NewBufferString("")
	for _, entry := range config.StartList {
		startListMap[entry.Name] = entry
		entryNamesBuffer.WriteString(entry.Name + "\n")
	}

	startMenuCmd := exec.Command("rofi", "-i", "-dmenu", "-p", "Start", "-l", "8")
	startMenuCmd.Stdin = entryNamesBuffer

	chosenNameBytes, err := startMenuCmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	chosenName := strings.TrimSuffix(string(chosenNameBytes), "\n")
	chosenEntry := startListMap[chosenName]

	chosenCmd := exec.Command(chosenEntry.Command, chosenEntry.Args...)

	_, err = chosenCmd.Output()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
