package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/yosssi/gohtml"
)

func SaveHTML(filename string, content string) error {
	timestamp := time.Now().Format("2006-01-02---15:04:05")
	fileNameWithTime := fmt.Sprintf("%s(%s).html", filename, timestamp)

	dirPath := filepath.Join("snapshots")
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create snapshots dir: %w", err)
	}

	filePath := filepath.Join(dirPath, fileNameWithTime)

	fmt.Println("Saving content in to:", filePath)

	prettyContent := gohtml.Format(content)

	if err := os.WriteFile(filePath, []byte(prettyContent), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
