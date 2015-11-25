package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func writeNBytes(bdb *bolt.DB, k []byte, N int) error {
	buf := make([]byte, N)
	return bdb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("predicate"))
		if err != nil {
			return err
		}
		return bucket.Put(k, buf)
	})
}

func readValue(bdb *bolt.DB, k []byte) (N int) {
	bdb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("predicate"))
		if bucket == nil {
			return errors.New("Bucket not found")
		}
		val := bucket.Get(k)
		m := make([]byte, len(val))
		copy(m, val)
		N = len(m)
		return nil
	})
	return
}

func main() {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	k := []byte("key")
	N := 512
	if err := writeNBytes(db, k, N); err != nil {
		log.Fatal(err)
	}
	fmt.Println(readValue(db, k))
}
