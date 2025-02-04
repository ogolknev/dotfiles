package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getVolume() (string, error) {
	cmd := exec.Command("pactl", "get-sink-volume", "@DEFAULT_SINK@")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("can't get volume: %w", err)
	}

	re := regexp.MustCompile(`\d+%`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) == 0 {
		return "", fmt.Errorf("can't parse volume")
	}

	volume := strings.TrimSuffix(matches[0], "%")

	return volume, nil
}

func main() {
	option := os.Args[1]

	volumeUp := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "+5%")
	volumeDown := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", "-5%")

	var err error

	if option == "+" || option == "up" {
		err = volumeUp.Run()
	} else if option == "-" || option == "down" {
		err = volumeDown.Run()
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	volume, err := getVolume()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = exec.Command("osd", "volume", volume).Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
