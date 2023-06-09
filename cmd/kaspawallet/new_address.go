package main

import (
	"context"
	"fmt"
	"github.com/ixbaseANT/gord/cmd/kaspawallet/daemon/client"
	"github.com/ixbaseANT/gord/cmd/kaspawallet/daemon/pb"
)

func newAddress(conf *newAddressConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	fmt.Print("ctx = %s", ctx)
	response, err := daemonClient.NewAddress(ctx, &pb.NewAddressRequest{})
	if err != nil {
		return err
	}

	fmt.Printf("New address:\n%s\n", response.Address)
	return nil
}
