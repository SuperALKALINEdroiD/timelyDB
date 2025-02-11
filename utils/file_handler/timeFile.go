package filehandler

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SuperALKALINEdroiD/timelyDB/config"
)

func GetFilePath(databaseConfig *config.DatabaseConfig) string {
	now := time.Now()

	var filename string
	if databaseConfig.TimelyConfig.IsEnabled {
		switch databaseConfig.TimelyConfig.Type {
		case 'h':
			filename = fmt.Sprintf("data_%d_%02d_%02d_%02d.json",
				now.Year(), now.Month(), now.Day(), now.Hour())
		case 'd':
			filename = fmt.Sprintf("data_%d_%02d_%02d.json",
				now.Year(), now.Month(), now.Day())
		case 'm':
			filename = fmt.Sprintf("data_%d_%02d.json",
				now.Year(), now.Month())
		default:
			filename = "data.json"
		}
	} else {
		filename = "data.json"
	}

	return filename
}

func CheckIfFileExists(databaseConfig *config.DatabaseConfig) (bool, error) {
	filename := GetFilePath(databaseConfig)

	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, fmt.Errorf("Failed to check for file: %v", err)
	}

	return true, nil
}

func AppendOrCreateFile(databaseConfig *config.DatabaseConfig, data []byte) error {
	exists, err := CheckIfFileExists(databaseConfig)

	if err != nil {
		return err
	}

	if !exists {
		log.Println("File does not exist, creating a new one...")
	}

	file, err := os.OpenFile(GetFilePath(databaseConfig), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	defer file.Close()
	_, err = file.Write(data)

	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}
