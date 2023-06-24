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


	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	i:=0
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

//		if strings.Contains(string(value), "311") {
//		    fmt.Println("key: ",string(key))
//		    fmt.Println("+++",string(value))
//		}
//		if strings.Contains(string(key), "multi") {
		if strings.Contains(string(value), "311") {

		i=i+1
//			fmt.Println("Подстрока найдена",)
// 			fmt.Printf("key: %s | value: %s\n", key[0:20], value)
//			fmt.Println(string(key[0:28]))
//			fmt.Println(key[12:])
//			fmt.Println(value)
	encodedHash := hex.EncodeToString(key[12:])
	encodedValue := hex.EncodeToString(value)
	fmt.Println("Подстрока найдена",encodedHash)
	fmt.Println("Value:", string(value))
	_,err=dbb.Exec("insert into multisets (p1,p2,p3,p4)values($1,$2,$3,$4)","string(key)",encodedHash,encodedValue,"value")
	if err != nil {
	    panic(err)
	}

		} else {
//			fmt.Println("Подстрока не найдена")
		}
	}

	fmt.Println("nnn===",i)
	iter.Release()
	err = iter.Error()
}
