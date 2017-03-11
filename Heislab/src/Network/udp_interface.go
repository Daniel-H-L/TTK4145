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
	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udp_interface_check_error(err)

	return conn
}

func Udp_interface_send(destinationIP string, data []byte) {
	fmt.Println("Sending... to ", destinationIP)
	localAddr, err := net.ResolveUDPAddr("udp", port)
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	fmt.Println("conn OK!")
	udp_interface_check_error(err)

	if len(data) > 0 {
		conn.Write(data)
		fmt.Println("Data sent")
	}

	time.Sleep(200 * time.Millisecond)
	defer conn.Close()
}

func Udp_interface_receive(msg chan []byte, portNr string, chan_kill chan bool) {
	fmt.Println("In received...")
	connection := Udp_interface_init(portNr)
	buffer := make([]byte, 1024)

	for {

		select {
		case <-chan_kill:
			fmt.Println("Killing interface")
			connection.Close()
			return
		}

		connection.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

		n, _, err := connection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Received msg...")
			udp_interface_check_error(err)

		}
		fmt.Println("Add buffer to chan...")
		msg <- buffer[0:n]
	}
	defer connection.Close()
}

func Udp_interface_bcast(data []byte) {
	bcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:40018")

	localAddr, err := net.ResolveUDPAddr("udp", ":0")

	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", localAddr, bcastAddr)
	udp_interface_check_error(err)
	defer conn.Close()

	if len(data) > 0 {
		conn.Write([]byte(data))
	}
	time.Sleep(200 * time.Millisecond)

}
