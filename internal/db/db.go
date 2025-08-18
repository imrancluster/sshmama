package db

import (
	"encoding/json"
	"errors"

	"github.com/imrancluster/sshmama/internal/app"
	"github.com/imrancluster/sshmama/internal/model"
	bolt "go.etcd.io/bbolt"
)

var ErrNotFound = errors.New("not found")

func PutEntry(a *app.App, e *model.Entry) error {
	return a.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.EntriesBucket))
		data, err := json.Marshal(e)
		if err != nil {
			return err
		}
		return b.Put([]byte(e.Name), data)
	})
}

func GetEntry(a *app.App, name string) (*model.Entry, error) {
	var out model.Entry
	err := a.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.EntriesBucket))
		v := b.Get([]byte(name))
		if v == nil {
			return ErrNotFound
		}
		return json.Unmarshal(v, &out)
	})
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func ListEntries(a *app.App) ([]model.Entry, error) {
	var res []model.Entry
	err := a.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.EntriesBucket))
		return b.ForEach(func(k, v []byte) error {
			var e model.Entry
			if err := json.Unmarshal(v, &e); err != nil {
				return err
			}
			res = append(res, e)
			return nil
		})
	})
	return res, err
}

func DeleteEntry(a *app.App, name string) error {
	// Also remove attachment if present
	ent, err := GetEntry(a, name)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	}
	return a.DB.Update(func(tx *bolt.Tx) error {
		if ent != nil && ent.AttachmentID != "" {
			_ = tx.Bucket([]byte(app.AttachmentsBucket)).Delete([]byte(ent.AttachmentID))
		}
		return tx.Bucket([]byte(app.EntriesBucket)).Delete([]byte(name))
	})
}

func PutAttachment(a *app.App, id string, data []byte) error {
	return a.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte(app.AttachmentsBucket)).Put([]byte(id), data)
	})
}

func GetAttachment(a *app.App, id string) ([]byte, error) {
	var out []byte
	err := a.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(app.AttachmentsBucket))
		v := b.Get([]byte(id))
		if v == nil {
			return ErrNotFound
		}
		out = append([]byte(nil), v...)
		return nil
	})
	return out, err
}
