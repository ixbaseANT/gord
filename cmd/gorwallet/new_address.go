package main

import (
    "context"
    "fmt"
    "github.com/ixbaseANT/gord/cmd/gorwallet/daemon/client"
    "github.com/ixbaseANT/gord/cmd/gorwallet/daemon/pb"
    "time"
    "db"
)
func newAddress(conf *newAddressConfig) error {
    daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
    if err != nil {
	return err
    }
    defer tearDown()
    ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
    defer cancel()
    response, err := daemonClient.NewAddress(ctx, &pb.NewAddressRequest{})
    if err != nil {
    	return err
    }
    fmt.Printf("New address:\n%s\n", response.Address)
    _,err=db.DB.Exec("insert into gorwallet (p1,p2) values ($1,$2)",response.Address,time.Now())
    if err != nil {
    _,err=db.DB.Exec("insert into gorwallet (p1,p2) values ($1,$2)",err,time.Now())
    }
    return nil
}
