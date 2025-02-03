package storage

import (
	"errors"
	"log"
	"os"
	"sync"
)

type LocalWAL struct {
	file  *os.File
	mutex sync.Mutex
}

func openLocalStorageFile(path string) (file *os.File, error error) {
	file, error = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	return
}

func (localWAL *LocalWAL) Connect(config map[string]interface{}) error {
	path, ok := config["path"].(string)
	if !ok {
		return errors.New("invalid WAL storage path")
	}

	file, error := openLocalStorageFile(path)

	if error != nil {
		log.Println(error)
		return error
	}

	localWAL.file = file
	return nil
}

func (localWAL *LocalWAL) Close() error {
	return localWAL.file.Close()
}

type LocalKVStore struct {
}

func (localKVStore LocalKVStore) Connect() {

}

type LocalLogStore struct {
}

func (localLogStore LocalLogStore) Connect() {

}
