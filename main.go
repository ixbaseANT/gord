// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main
import (
    _ "net/http/pprof"
    "os"
    "fmt"
    "github.com/ixbaseANT/gord/app"
)
func main() {
    fmt.Println("Connected to the database!")
    if err := app.StartApp(); err != nil {
	os.Exit(1)
    }
}
