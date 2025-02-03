package storage

import (
	"errors"
	"log"
	"os"
	"sync"
)

type LocalWAL struct {
	file    *os.File
	mutex   sync.Mutex
	maxSize int
}

func openLocalStorageFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
}

func (localWAL *LocalWAL) Connect(config map[string]interface{}) error {
	path, ok := config["path"].(string)
	if !ok {
		return errors.New("invalid WAL storage path")
	}

	maxSize, ok := config["maxSize"].(int) // TODO: see if i can do something with this
	if !ok {
		return errors.New("invalid max size for WAL")
	}

	file, err := openLocalStorageFile(path)

	if err != nil {
		log.Println("Error opening file:", err)
		return err
	}

	localWAL.file = file
	localWAL.maxSize = maxSize

	return nil
}

func (localWAL *LocalWAL) Close() error {
	return localWAL.file.Close()
}

func (localWAL *LocalWAL) GetSize() (int, error) {
	fileInfo, err := localWAL.file.Stat()
	if err != nil {
		return 0, err
	}
	return int(fileInfo.Size()), nil
}

func (localWAL *LocalWAL) WriteLog(data []byte) error {
	localWAL.mutex.Lock()
	defer localWAL.mutex.Unlock()

	_, err := localWAL.file.Write(data)
	if err != nil {
		log.Println("Error writing log:", err)
		return err
	}

	localWAL.file.Sync() // TODO: find a better time to sync changes to stable storage

	return nil
}

type LocalKVStore struct {
}

func (localKVStore LocalKVStore) Connect() {

}

type LocalLogStore struct {
}

func (localLogStore LocalLogStore) Connect() {

}
