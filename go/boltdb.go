package main

import (
	"log"
	"os"

	bolt "go.etcd.io/bbolt"
)

var (
	keyBucketName = []byte("lease")
)

func main() {
	db, err := bolt.Open(os.Args[1], 0600, &bolt.Options{ReadOnly: true})

	if err != nil {
		log.Fatalln("open db failed", err)
	}

	err = db.View(func(tx *bolt.Tx) error {
		//stats := tx.Bucket(keyBucketName).Stats()
                 b := tx.Bucket(keyBucketName)
                 c := b.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		log.Printf("key=%s, value=%s\n", k, v)
	}

	return nil
		//log.Printf("db static: %#v\n", stats)
		//return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
