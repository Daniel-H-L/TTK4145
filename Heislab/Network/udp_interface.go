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

func udp_interface_init(portNr string) *net.UDPConn {
	//socket
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udp_interface_check_error(err)

	return conn
}

func Udp_interface_send(destinationIP string, data []byte) {
	localAddr, err := net.ResolveUDPAddr("udp", port)
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	udp_interface_check_error(err)

	if len(data) > 0 {
		conn.Write(data)
	}

	time.Sleep(200 * time.Millisecond)
}

func Udp_interface_receive(msg chan string, portNr string, chan_error chan error) {
	connection := udp_interface_init(portNr)
	connection.SetReadDeadline(time.Now().Add(time.Second))
	buffer := make([]byte, 1024)

	for {
		defer connection.Close()
		n, _, err := connection.ReadFromUDP(buffer)

		if err != nil {
			error_chan <- err
			udp_interface_check_error(err)
			return
		}
		msg <- buffer[0:n]
		connection.SetReadDeadline(time.Now().Add(time.Second))
	}
}

func Udp_interface_bcast(data []byte) {
	localAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:40018")
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	udp_interface_check_error(err)

	if len(data) > 0 {
		conn.Write([]byte(data))
	}
	time.Sleep(200 * time.Millisecond)

}
