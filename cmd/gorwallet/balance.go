package main

import (
	"context"
	"fmt"
	"strings"
	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/client"
	"github.com/ixbaseANT/gord/cmd/gorwallet/daemon/pb"
	"github.com/ixbaseANT/gord/cmd/gorwallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
	p1:=strings.Trim(addressBalance.Address, " ")
	p2:=strings.Trim(utils.FormatKas(addressBalance.Available), " ")
	p3:=strings.Trim(utils.FormatKas(addressBalance.Pending), " ")
			fmt.Printf("%s %s %s\n",p1,p2,p3)
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, GOR %s %s%s\n", utils.FormatKas(response.Available), utils.FormatKas(response.Pending), pendingSuffix)

	return nil
}
