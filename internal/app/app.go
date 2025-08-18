package app

import (
	"os"
	"path/filepath"
	"time"

	bolt "go.etcd.io/bbolt"
)

const (
	DBFileName        = "sshmama.db"
	EntriesBucket     = "entries"
	AttachmentsBucket = "attachments"
	DefaultDirName    = "sshmama"
)

type App struct {
	DataDir string
	DB      *bolt.DB
}

func DefaultDataDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		home, e2 := os.UserHomeDir()
		if e2 != nil {
			return "", e2
		}
		return filepath.Join(home, ".config", DefaultDirName), nil
	}
	return filepath.Join(base, DefaultDirName), nil
}

func New(dataDir string) (*App, error) {
	if err := os.MkdirAll(dataDir, 0o700); err != nil {
		return nil, err
	}
	dbPath := filepath.Join(dataDir, DBFileName)
	db, err := bolt.Open(dbPath, 0o600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		if _, e := tx.CreateBucketIfNotExists([]byte(EntriesBucket)); e != nil {
			return e
		}
		if _, e := tx.CreateBucketIfNotExists([]byte(AttachmentsBucket)); e != nil {
			return e
		}
		return nil
	}); err != nil {
		_ = db.Close()
		return nil, err
	}
	return &App{DataDir: dataDir, DB: db}, nil
}

func (a *App) Close() {
	if a.DB != nil {
		_ = a.DB.Close()
	}
}
