package storage

type Storage interface {
	Connect(config map[string]interface{}) error // connect to storage
	Close() error                                // disconnect
	GetSize() (int, error)                       // get in memory size
}

type WAL interface {
	Storage
	WriteLog(data []byte) error           // write WAL logs
	ReadLog(offset int64) ([]byte, error) // read WAL Logs
	TruncateLog(offset int64) error       // save logs to memory and clear memory
}

type KVStore interface {
	Storage
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type LogStore interface {
	Storage
	WriteLog(app string, data []byte) error
	ReadLog(app string, offset int64) ([]byte, error)
	DeleteLogs(app string) error
}
