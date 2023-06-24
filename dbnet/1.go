package main

import (
	"fmt"
	"path/filepath"
	"github.com/syndtr/goleveldb/leveldb"
//	"github.com/syndtr/goleveldb/leveldb/util"
	"log"
	"strings"
	"encoding/hex"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
    dbb, err := sql.Open("postgres", "host=localhost port=5432 user=gorbaniov password=1 dbname=gor sslmode=disable")
    if err != nil {
	fmt.Println("Ошибка при подключении к базе данных:", err)
	return
    }
    defer dbb.Close()
	dbPath := filepath.Join("/home/ixb/.gord/gor-mainnet/datadir2", "log")
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatal("Yikes!")
	}
	defer db.Close()

	//	err = db.Put([]byte("fizz"), []byte("buzz"), nil)
	//	err = db.Put([]byte("fizz2"), []byte("buzz2"), nil)
	//	err = db.Put([]byte("fizz3"), []byte("buzz3"), nil)

	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	i:=0
	for iter.Next() {
		key := iter.Key()
//		value := iter.Value()
		if strings.Contains(string(key), "chain-block-index-by-hash") {
		i=i+1
//			fmt.Println("Подстрока найдена",i)
// 			fmt.Printf("key: %s | value: %s\n", key[0:20], value)
//			fmt.Println(string(key[0:28]))
//			fmt.Println(key[28:])
//			fmt.Println(value)
	encodedHash := hex.EncodeToString(key[28:])
	fmt.Println("Encoded Hash:", encodedHash)
	_,err=dbb.Exec("insert into netblocks (poolid,hash)values($1,$2)","GOR", encodedHash)
	if err != nil {
	    panic(err)
	}

		} else {
//			fmt.Println("Подстрока не найдена")
		}
	}

	fmt.Println("nnn===",i)
/*
	for ok := iter.Seek([]byte("block")); ok; ok = iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, value)
	}

	fmt.Println("222 \n")

	for ok := iter.First(); ok; ok = iter.Next() {
		key := iter.Key()
//		value := iter.Value()
		fmt.Printf("key: %s | value: %s\n", key, "value")
	}
*/
	iter.Release()
	err = iter.Error()
}
