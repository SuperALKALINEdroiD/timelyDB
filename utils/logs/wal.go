package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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

func WriteAheadMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: Initialize WAL based on configuration
		walStorage := &storage.LocalWAL{}

		writeAheadEntry := WriteAheadEntry{
			EntryID:   uuid.New().String(),
			NodeID:    0, // TODO: Retrieve actual node ID
			Timestamp: time.Now(),
			File:      walStorage,         // Initialized WAL storage
			Data:      []byte("log data"), // Replace with actual data
			Status:    Pending,
		}

		logData, err := json.Marshal(writeAheadEntry)
		if err != nil {
			http.Error(w, "Failed to serialize log entry", http.StatusInternalServerError)
			return
		}

		if err := writeAheadEntry.File.WriteLog(logData); err != nil {
			http.Error(w, "Failed to write to WAL", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "WALEntryId", writeAheadEntry.EntryID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WriteAheadStatusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entryID, ok := r.Context().Value("WALEntryId").(string)
		if !ok {
			http.Error(w, "Missing WAL Entry ID in context", http.StatusInternalServerError)
			return
		}

		// TODO update wal status based on entryId

		fmt.Println(entryID)

		next.ServeHTTP(w, r)
	})
}
