package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
//	"strings"
	"net/http"
	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/client"
	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/pb"
)

func newAddress(conf *newAddressConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
	    return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
//	fmt.Print("ctx = %s",ctx)
	response, err := daemonClient.NewAddress(ctx, &pb.NewAddressRequest{})
	if err != nil {
	    return err
	}

	resp, err := http.Get("https://taxi-x.org/v.php?ix=gor-new-address")
	if err != nil {
	  fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
    	log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf("HTML address:\n%s\n", sb)
	fmt.Printf("New address:\n%s\n", response.Address)

	return nil
}
