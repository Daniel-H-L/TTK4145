package Network

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const PORTNR = ":40018"

const (
	IsAlive    = 0
	Order      = 1
	Descendant = 2
	Lights     = 3
)

func udpInterfaceCheckError(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
	}
}

func UDPInit() *net.UDPConn {
	fmt.Println("Init connection: ", PORTNR)

	localAddr, err := net.ResolveUDPAddr("udp", PORTNR)
	udpInterfaceCheckError(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udpInterfaceCheckError(err)

	return conn
}

func UDPSend(dataCh chan StandardData) {

	localAddr, err := net.ResolveUDPAddr("udp", PORTNR)
	udpInterfaceCheckError(err)

	sendConn, err := net.DialUDP("udp", nil, localAddr)
	defer sendConn.Close()

	udpInterfaceCheckError(err)

	dataCh := make(chan StandardData, 100)

	if len(jsonObject) > 0 {
		select {
		case Alive := <-dataCh:
			data := StandardData{}
			data.IsAlive = Alive
			data.IP = Udp_get_local_ip()
			data.Type = 0
			dataCh <- data

		case Order := <-<-dataCh:
			data := StandardData{}
			data.Backup = Order
			data.Type = 1
			dataCh <- data

		case Descendant := <-dataCh:
			data := StandardData{}
			data.DescendantNr = Descendant
			data.IP = Udp_get_local_ip()
			data.Type = 2
			dataCh <- data

		case Lights := <-dataCh:
			data := StandardData{}
			data.SetLights = Lights
			data.IP = Udp_get_local_ip()
			data.Type = 3
			dataCh <- data

		case <-dataCh:
			json_object, _ := json.Marshal(data)
			conn.Write(json_object)
		}
	}

}

func UDPReceive(descendantCh chan int, backupCh chan Backup, masterIPCh chan string, senderIPCh string) {
	recConn := UDPInit()
	defer recConn.Close()

	buffer := make([]byte, 1024)

	recConn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

	n, _, err := recConn.ReadFromUDP(buffer)

	if n > 0 {
		var data StandardData
		json.Unmarshal(buffer[0:n], &data)

		switch data.Type {

		case IsAlive:
			senderIPCh <- data.IP
			masterIPCh <- data.IsAlive

		case Backup:
			backupCh <- data.Backup
			senderIPCh <- data.IP

		case DescendantNr:
			descendantCh <- data.DescendantNr
			senderIPCh <- data.IP
		}
	}
}
