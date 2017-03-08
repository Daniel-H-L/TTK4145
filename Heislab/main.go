package main

import (
	"./DriveElevator"
	"./Master"
	"./Network"
	"./Slave"
	"fmt"
	"time"
)

var backup = Network.Backup{}

func main() {

	Slave.Slave_run(*backup)


}


for{

	select{

		case <- channel: 



	}
}