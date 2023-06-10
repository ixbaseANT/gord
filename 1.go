package main

import (
	"fmt"
	"file"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	dbPath := filepath.Join("/home/qq/.gord/gor-mainnet/datadir2", "*.log")
	db, err := leveldb.OpenFile(dbPath, nil)
//	db, err := leveldb.OpenFile("/home/qq/.gord/gor-mainnet/datadir2", nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()

	//	err = db.Put([]byte("fizz"), []byte("buzz"), nil)
	//	err = db.Put([]byte("fizz2"), []byte("buzz2"), nil)
	//	err = db.Put([]byte("fizz3"), []byte("buzz3"), nil)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.Seek([]byte("fizz2")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("\n")

	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	iter.Release()
	err = iter.Error()
}
