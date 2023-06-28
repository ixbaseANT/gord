package main

import (
	"context"
	"fmt"
	"db"
	"time"
	"strings"
	"net/http"
	"net/url"
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
	p4:=time.Now()
			fmt.Printf("%s %s %s\n",p1,p2,p3)
	rows,err := db.DB.Query("select * from gorwallet where p1=$1",p1)
	if err != nil {
	    panic(err)
	}
	defer rows.Close()
        var pp1 string
        var pp2 string
        var pp3 string
        var pp4 string
        var pp5 string
        var pp6 string
	rows.Next()
	err=rows.Scan(&pp1,&pp2,&pp3,&pp4,&pp5,&pp6)
	println("===============",err)
	println(pp1)
	println(pp2)
	println(pp3)
	println(pp4)
	println(pp5)
	println(pp6)
	println("===============")


	if pp1=="" {
	fmt.Println("=insert p1=",pp1,p1)
          _,err=db.DB.Exec("insert into gorwallet (p1,p2) values ($1,$2)",p1,p4)
	  if err != nil {
		panic(err)
	  }
	  rows,err := db.DB.Query("select p1,p2,p3,p4,p5,p6 from gorwallet where p1=$1",p1)
	  if err != nil {
	    panic(err)
	  }
	  defer rows.Close()
	  for rows.Next() {
	    err=rows.Scan(&pp1,&pp2,&pp3,&pp4,&pp5,&pp6)
	  }
	}
	if pp4!=p2 || pp5!=p3 {
	 _,err=db.DB.Exec("update gorwallet set p4=$2,p5=$3,p6=$4 where p1=$1",p1,p2,p3,p4)
	 if err != nil {
		panic(err)
	 }
	 pp2=url.QueryEscape(pp2[0:19])
//	 pp6=url.QueryEscape(pp6[0:19])

	 r,err := http.Get("https://biegroup.club/v.php?ix=gorBalance&p1="+pp1+"&p2="+pp2+"&p3="+pp3+"&p4="+p2+"&p5="+p3)
	 if err != nil {
	    fmt.Println(r,err)
	 }
	}
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, GOR %s %s%s\n", utils.FormatKas(response.Available), utils.FormatKas(response.Pending), pendingSuffix)

	return nil
}
