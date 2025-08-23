package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func SaveHTML(filename string, content string) error {
	timestamp := time.Now().Format("2006-01-02---15:04:05")
	fileNameWithTime := fmt.Sprintf("%s(%s).html", filename, timestamp)

	rootDir, err := filepath.Abs(filepath.Join("../.."))
	if err != nil {
		return fmt.Errorf("failed to get project root path: %w", err)
	}

	dirPath := filepath.Join(rootDir, "snapshots")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create snapshots dir: %w", err)
	}

	filePath := filepath.Join(dirPath, fileNameWithTime)

	fmt.Println("Saving content in to:", filePath)

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
