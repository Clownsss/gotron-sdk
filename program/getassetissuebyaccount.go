package main

import (
	"flag"
	"fmt"
	"github.com/tronprotocol/go-client-api/service"
	"log"
	"strings"
)

func main() {
	grpcAddress := flag.String("grpcAddress", "",
		"gRPC address: localhost:50051")

	address := flag.String("address", "",
		"address: TH4MqfHpmKFpdtvPaeYufXHiPyxLvmJWP6")

	flag.Parse()

	if strings.EqualFold("", *address) && len(*address) == 0 {
		log.Fatalln("./get-asset-issue-by-account -grpcAddress localhost" +
			":50051 -address TH4MqfHpmKFpdtvPaeYufXHiPyxLvmJWP6")
	}

	client := service.NewGrpcClient(*grpcAddress)
	client.Start()
	defer client.Conn.Close()

	assetIssueList := client.GetAssetIssueByAccount(*address)

	fmt.Printf("asset issue list: %v\n",
		assetIssueList)
}