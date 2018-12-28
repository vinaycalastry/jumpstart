package db

import (
	"time"

	"github.com/boltdb/bolt"
)

var linkBucket = []byte("links")
var db *bolt.DB

type Link struct {
	Key   string
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(linkBucket)
		return err
	})

}

func AddLink(title string, link string) error {

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(linkBucket)
		id, _ := b.NextSequence()
		_ = id
		return b.Put([]byte(title), []byte(link))
	})

	if err != nil {
		return err
	}
	return nil
}

func GetLinks() ([]Link, error) {
	var links []Link
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(linkBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			links = append(links, Link{
				Key:   string(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return links, nil
}

func FindLink(title string) ([]byte, error) {
	var url []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(linkBucket)
		v := b.Get([]byte(title))
		url = v
		return nil
	})

	if err != nil {
		return nil, err
	}

	return url, nil
}
