// logs/wal.go
package logs

import (
	"fmt"
	"os"
	"sync"
)

type WAL struct {
	NodeID string
	File   *os.File
	Mutex  sync.Mutex
}

func SetupWAL(nodeID string) (*WAL, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s.wal", nodeID), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &WAL{
		NodeID: nodeID,
		File:   file,
	}, nil
}

func (wal *WAL) Write(entry []byte) error {
	wal.Mutex.Lock()
	defer wal.Mutex.Unlock()

	_, err := wal.File.Write(entry)
	return err
}
