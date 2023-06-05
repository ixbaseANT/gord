package main

import (
//	"context"
	"fmt"
	"io/ioutil"
	"log"
//	"strings"
	"net/http"
//	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/client"
//	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/pb"
)

func newAddress(conf *newAddressConfig) error {
	resp, err := http.Get("https://taxi-x.org/v.php?ix=gor-new-address")
	if err != nil {
	  fmt.Println(err)
	}
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
   sb := string(body)
	fmt.Printf("New address:\n%s\n", sb)
	return nil
}
