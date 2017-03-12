package Network

import (
	"encoding/json"
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

	fmt.Println("Init connection: ", portNr)

	localAddr, err := net.ResolveUDPAddr("udp", portNr)
	udp_interface_check_error(err)

	conn, err := net.ListenUDP("udp", localAddr)
	udp_interface_check_error(err)

	return conn
}

func Udp_interface_send(destinationIP string, standarddata StandardData) {
	fmt.Println("Sending... to ", destinationIP)
	fmt.Println("Sending... ", standarddata)
	json_object, _ := json.Marshal(standarddata)
	fmt.Println("buffer  ", string(json_object))

	localAddr, err := net.ResolveUDPAddr("udp", port)
	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, localAddr)
	defer conn.Close()
	//fmt.Println("conn OK!")
	udp_interface_check_error(err)

	if len(json_object) > 0 {

		conn.Write(json_object)
		fmt.Println("Data sent")
	}

	time.Sleep(200 * time.Millisecond)

}

func Udp_interface_receive(msg chan StandardData, portNr string, chan_kill chan bool) {
	connection := udp_interface_init(portNr)
	defer connection.Close()
	buffer := make([]byte, 1024)

	fmt.Println("In received...")

	for {

		select {
		case <-chan_kill:
			fmt.Println("Killing interface", portNr)
			return
		default:
			connection.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

			n, _, err := connection.ReadFromUDP(buffer)
			/*if err.Timeout() {
				continue
			} else*/if err != nil {
				//fmt.Println("Received msg...")
				//udp_interface_check_error(err)

			} else if n > 0 {
				//fmt.Println("Add buffer to chan...")
				var struct_object StandardData

				//fmt.Println(string(buffer))

				json.Unmarshal(buffer[0:n], &struct_object)
				msg <- struct_object
				time.Sleep(2 * time.Millisecond)
			}
		}
	}

}

func Udp_interface_bcast(data []byte) {
	bcastAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:40018")

	//localAddr, err := net.ResolveUDPAddr("udp", ":0")

	udp_interface_check_error(err)

	conn, err := net.DialUDP("udp", nil, bcastAddr)
	defer conn.Close()
	udp_interface_check_error(err)

	if len(data) > 0 {
		conn.Write([]byte(data))
	}
	time.Sleep(500 * time.Millisecond)

}
