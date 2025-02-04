package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func getNotificationIDFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("can't get user directory: %w", err)
	}
	cacheDir := filepath.Join(homeDir, ".cache/osd")

	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("can't create directory: %w", err)
	}

	notificationIDFile := filepath.Join(cacheDir, "notification-id.tmp")
	return notificationIDFile, nil
}

func getNotificationID() (string, error) {

	notificationIDFile, err := getNotificationIDFile()
	if err != nil {
		return "", fmt.Errorf("can't receive notification id file: %w", err)
	}

	var notificationID string
	fileInfo, err := os.Stat(notificationIDFile)
	if os.IsNotExist(err) || fileInfo.Size() == 0 {
		notificationID = "0"
	} else if err != nil {
		return "", fmt.Errorf("can't open notification id file: %w", err)
	} else {
		out, err := os.ReadFile(notificationIDFile)
		if err != nil {
			return "", fmt.Errorf("can't read notification id file: %w", err)
		}
		notificationID = strings.TrimSpace(string(out))
		_, err = strconv.Atoi(notificationID)
		if err != nil {
			notificationID = "0"
		}
	}

	return notificationID, nil
}

func setNotificationID(value []byte) error {
	notificationIDFile, err := getNotificationIDFile()
	if err != nil {
		return fmt.Errorf("can't receive notification id file: %w", err)
	}

	err = os.WriteFile(notificationIDFile, value, 0644)
	if err != nil {
		return fmt.Errorf("can't write notification id file: %w", err)
	}

	return nil
}

type NotifyOptions struct {
	App      string
	Category string
	Expire   string
	Value    string
}

func notify(msg string, options NotifyOptions) error {

	replaceID, err := getNotificationID()
	if err != nil {
		return fmt.Errorf("can't receive notification id: %w", err)
	}

	args := []string{msg, "-p", "-r", replaceID}
	if options.App != "" {
		args = append(args, "-a", options.App)
	}
	if options.Category != "" {
		args = append(args, "-c", options.Category)
	}
	if options.Expire != "" {
		args = append(args, "-t", options.Expire)
	}
	if options.Value != "" {
		args = append(args, "-h", options.Value)
	}

	cmd := exec.Command("notify-send", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't send notify: %s(%w)", out, err)
	}

	setNotificationID(out)

	return nil
}
