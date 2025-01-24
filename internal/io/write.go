package write

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var dbPath = filepath.Join(os.Getenv("LOCALAPPDATA"), "sleeper", "data.csv")

func CreateDB() error {
	wasCreated, err := createFilePath(dbPath)

	if err != nil {
		return err
	}

	if !wasCreated {
		return nil
	}

	header := "DateBefore,DateAfter,Duration"

	if err := writeLineToFile(dbPath, header); err != nil {
		return err
	}

	return nil
}

func AddRecord(before time.Time, after time.Time, duration string) error {
	format := "20060102150405"

	record := fmt.Sprintf("%s,%s,%s", before.Format(format), after.Format(format), duration)

	if err := writeLineToFile(dbPath, record); err != nil {
		return err
	}

	return nil
}

func createFilePath(filePath string) (bool, error) {
	dir := filepath.Dir(filePath)

	// Creates all the directories in the path that don't exist. 0755 is the permission mode
	if err := os.MkdirAll(dir, 0755); err != nil {
		return false, err
	}

	// os.Stat() tries to retieve the information of the filePath and throws an error if that is not possible
	// os.IsNotExist() parses this error to check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return false, err
		}
		defer file.Close()

		return true, nil // File was created
	} else if err != nil {
		return false, err
	}

	// File already exists, do nothing
	return false, nil
}

func writeLineToFile(filePath, line string) error {
	// Ensure file exists
	if _, err := createFilePath(filePath); err != nil {
		return err
	}

	// Open file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the line
	_, err = file.WriteString(line + "\n")
	return err
}
