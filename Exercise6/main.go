package main

import (
	"./Exercise6"
	"fmt"
	"net"
)

func main() {
	aliveChan := make(chan string)
	aliveChan <- "Alive"

	go Primary_bcast(aliveChan)

}
