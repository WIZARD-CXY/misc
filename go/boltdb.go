package main

import (
	"log"
	"os"

	bolt "github.com/coreos/bbolt"
)

var (
	keyBucketName = []byte("key")
)

func main() {
	db, err := bolt.Open(os.Args[1], 0600, &bolt.Options{ReadOnly: true})

	if err != nil {
		log.Fatalln("open db failed", err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		stats := tx.Bucket(keyBucketName).Stats()

		log.Printf("db static: %#v\n", stats)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
