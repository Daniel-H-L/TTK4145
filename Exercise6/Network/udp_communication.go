package Network

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var localIP string
var masterIP string

func udp_get_local_ip() (string, error) {
	if localIP == "" {
		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: []byte{8, 8, 8, 8}, Port: 30018})
		if err != nil {
			return "", err
		}
		defer conn.Close()
		localIP = strings.Split(conn.LocalAddr().String(), ":")[0]
	}
	return localIP, nil
}

func udp_get_master_ip(is_master bool) string {
	if is_master {
		masterIP, _ = udp_get_local_ip()
	}

	return masterIP
}

//Only master
func Udp_broadcast(msg_id string, elevator_nr int) {
	fmt.Println("Udp_bcast connection 3")
	fmt.Println("Communic: ", elevator_nr)
	send_object := StandardData{}
	send_object.IP, _ = udp_get_local_ip()
	send_object.Msg_ID = msg_id
	send_object.Order_executed = elevator_nr
	fmt.Println("Elev_nr: ", send_object.Order_executed)

	send := Udp_struct_to_json(send_object)

	Udp_interface_bcast(send)

	fmt.Println("Udp_bcast connection 1")
	time.Sleep(200 * time.Millisecond)
}

func udp_send_is_alive(destination_ip string) {
	alive := StandardData{}
	chan_is_alive := make(chan []byte)
	alive.IP, _ = udp_get_local_ip()
	send := Udp_struct_to_json(alive)

	for {
		Udp_interface_send(destination_ip, chan_is_alive) //burde for-løkken heller implementeres i laget over?
		chan_is_alive <- send
		time.Sleep(time.Second)
	}
}

// func udp_receive_is_alive(chan_received_msg chan []byte, chan_is_alive chan string, portNr string) {
// 	go Udp_interface_receive(chan_received_msg, portNr)

// 	received := <-chan_received_msg

// 	data := Udp_json_to_struct(received, 1024)
// 	chan_is_alive <- data.IP
// }
//Only slaves
func udp_send_order_executed(order_nr int, is_master bool) {
	order := StandardData{}
	chan_order_executed := make(chan []byte)
	order.Order_executed = order_nr

	chan_order_executed <- Udp_struct_to_json(order)

	Udp_interface_send(udp_get_master_ip(is_master), chan_order_executed)
}

//Only master
func Udp_receive_order_executed(chan_order_executed chan int, chan_received_msg chan []byte, portNr string, chan_error chan error, state int) {
	err_chan2 := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, err_chan2, state)

	for {
		select {
		case received := <-chan_received_msg:

			data := Udp_json_to_struct(received, 1024)

			//chan_order_executed <- data.Order_executed
			fmt.Println("Received: ", data.Order_executed)

		case err := <-err_chan2:
			fmt.Println("Error detected\n")
			chan_error <- err
			return
		}
	}

}

//Only master
func udp_send_descendant_nr(chan_descendant_nr chan []byte, descendant_nr int, dest_ip string) {
	number := StandardData{}
	number.Descendant_nr = descendant_nr
	chan_descendant_nr <- Udp_struct_to_json(number)
	Udp_interface_send(dest_ip, chan_descendant_nr)
}

//Only slave
// func udp_receive_descendant_nr(chan_descendant_nr chan int, chan_received_msg chan []byte, portNr string) {
// 	go Udp_interface_receive(chan_received_msg, portNr)

// 	received := <-chan_received_msg
// 	data := Udp_json_to_struct(received, 1024)

// 	chan_descendant_nr <- data.Descendant_nr
// }

//Only master
func udp_send_new_order(new_order NewOrder, dest_ip string, chan_new_order chan []byte) {
	order := StandardData{}
	order.New_order = new_order
	chan_new_order <- Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, chan_new_order)
}

// //Only slave
// func udp_receive_new_order(chan_new_order chan NewOrder, chan_received_msg chan []byte, portNr string) {
// 	for {
// 		Udp_interface_receive(chan_received_msg, portNr)
// 		received := <-chan_received_msg
// 		data := Udp_json_to_struct(received, 1024)
// 		chan_new_order <- data.New_order
// 	}
// }

func udp_send_local_order(local_order LocalOrder, dest_ip string, chan_local_order chan []byte) {
	order := StandardData{}
	order.Local_order = local_order
	chan_local_order <- Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, chan_local_order)
}

//Only master
// func udp_receive_local_order(chan_local_order chan LocalOrder, chan_received_msg chan []byte, portNr string) {
// 	for {
// 		Udp_interface_receive(chan_received_msg, portNr)
// 		received := <-chan_received_msg
// 		data := Udp_json_to_struct(received, 1024)
// 		chan_local_order <- data.Local_order
// 	}
// }
