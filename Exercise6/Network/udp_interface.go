package Network

import (
	"fmt"
	"net"
	"time"
)

var port string = ":30018"
var bcastPort string = ":40018"

func udp_interface_check_error(err error) {
	if err != nil {
		fmt.Println("... Error: ", err)
	}
}

func Udp_interface_init(portNr string) *net.UDPConn {
	//socket
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udp_interface_check_error(err)

	return conn
}

func Udp_interface_send(destinationIP string, msg chan []byte) { //removed string chan
	localAddr, err := net.ResolveUDPAddr("udp", port)
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	udp_interface_check_error(err)
	defer conn.Close()

	data := <-msg

	if len(data) > 0 {
		conn.Write([]byte(data))
	}

	time.Sleep(200 * time.Millisecond)
}

func Udp_interface_receive(msg chan []byte, portNr string, error_chan chan error, state int) {
	connection := Udp_interface_init(portNr)
	connection.SetReadDeadline(time.Now().Add(time.Second))
	buffer := make([]byte, 1024)

	for {
		defer connection.Close()
		n, _, err := connection.ReadFromUDP(buffer) //senderIP

		if err != nil {
			go func() { error_chan <- err }()
			udp_interface_check_error(err)
		}
		msg <- buffer[0:n] //removed string()
	}
}

func Udp_interface_bcast(msg chan []byte) {
	localAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:40018")
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	udp_interface_check_error(err)
	defer conn.Close()

	data := <-msg

	if len(data) > 0 {
		conn.Write([]byte(data))
	}
	time.Sleep(200 * time.Millisecond)

}
