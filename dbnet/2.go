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
	dbPath := filepath.Join("/home/ixb/.kaspad/kaspa-mainnet/datadir2", "")
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
//		value := iter.Value()
		if strings.Contains(string(key), "chain-block-index-by-hash") {
		i=i+1
	encodedHash := hex.EncodeToString(key[28:])
	fmt.Println("Encoded:", 1)

        var id int

	rows,err := dbb.Query("SELECT id FROM netkaspa WHERE hash = $1", encodedHash)
        if err != nil {
	    log.Fatal(err)
        }
        defer rows.Close()
        for rows.Next() {
	    err := rows.Scan(&id)
	    if err != nil {
	        log.Fatal(err)
	    }
	}
    if id<1 {
	_,err=dbb.Exec("insert into netkaspa (hash)values($1)", encodedHash)
	if err != nil {
	    panic(err)
	}
    }
		} else {
//			fmt.Println("Подстрока не найдена")
		}
	}

	fmt.Println("nnn===",i)
	iter.Release()
	err = iter.Error()
}
