package main

import (
	"fmt"

	"poc-replication/process"

	"github.com/georgexx009/mock-network-golang/network"
)

func main() {
	fmt.Println("hello world")
	network.Init()
	process.New()
}
