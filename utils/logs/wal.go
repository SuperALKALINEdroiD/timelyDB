package logs

import (
	"encoding/json"
	"time"

	"github.com/SuperALKALINEdroiD/timelyDB/utils/storage"
	"github.com/google/uuid"
)

type EntryStatus string

const (
	Pending   EntryStatus = "PENDING"
	Committed EntryStatus = "COMMITTED"
	Failed    EntryStatus = "FAILED"
)

type WriteAheadEntry struct {
	EntryID   string
	NodeID    int       // node where data will be saved
	Timestamp time.Time // entry time
	File      storage.WAL
	Data      []byte // The actual data being written
	Status    EntryStatus
}

func GetWAlStorage() storage.WAL {
	// TODO: intialize the storage based on config after integration
	return &storage.LocalWAL{}
}

func AddWalEntry() { // TODO look intpo reverting WAL entry incase write fails
	// TODO accept data to be logged as parameter
	walStorage := GetWAlStorage()

	writeAheadEntry := WriteAheadEntry{
		EntryID:   uuid.New().String(),
		NodeID:    0, // TODO: Retrieve actual node ID
		Timestamp: time.Now(),
		File:      walStorage,         // Initialized WAL storage
		Data:      []byte("log data"), // Replace with actual data
		Status:    Committed,
	}

	logData, err := json.Marshal(writeAheadEntry)
	if err != nil {
		panic("Failed to serialize the log entry")
	}

	if err := writeAheadEntry.File.WriteLog(logData); err != nil {
		panic("Failed to write to write ahead logs")
	}
}
