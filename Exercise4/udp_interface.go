package udp_interface

import (
	"fmt"
	"net"
	"time"
)

var portNr string = ":30018"

func udp_interface_check_error(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
	}
}

func udp_interface_init() net.Listener {
	//socket
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udp_interface_check_error(err)

	return conn
}

func udp_interface_send(destinationIP string, msg chan string) {
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", localAddr)
	udp_interface_check_error(err)

	data := <-msg

	if len(data) > 0 {
		conn.Write([]byte(data))
	}

	time.Sleep(200 * time.Millisecond)
}

func udp_interface_receive(msg chan string) {
	connection := udp_interface_init()

	buffer := make([]byte, 1024)

	for {
		n, _, err := connection.ReadFromUDP(buffer) //senderIP

		if err != nil {
			udp_interface_check_error(err)
		}
		msg <- (string(buffer[0:n]))
	}
}

func udp_interface_bcast(msg chan string) {
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := conn.DialBroadcastUDP(portNr)
	udp_interface_check_error(err)

	data := <-msg

	if len(data) > 0 {
		conn.Write([]byte(data))
	}
	time.Sleep(200 * time.Millisecond)

}
