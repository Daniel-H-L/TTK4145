package Network

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var localIP string
var masterIP

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

func udp_get_master_ip(chan_is_master chan bool) (string, error) {
	is_master := <- chan_is_master

	if is_master {
		masterIP = udp_get_local_ip()
	}

	return masterIP
}

//Only master
func Udp_broadcast_ip(msg_id string, elevator_nr int, chan_udp_bcast chan string) {
	send_object := StandardData{}
	send_object.IP, _ = udp_get_local_ip()
	send_object.msg_ID = msg_id
	send_object.elevator_nr = elevator_nr

	send := Udp_struct_to_json(send_object)

	Udp_interface_bcast(chan_udp_bcast)
	chan_udp_bcast <- string(send)
}

func udp_send_is_alive(destination_ip string, chan_is_alive chan string) {
	alive := StandardData{}
	alive.IP, _ = udp_get_local_ip()
	send := Udp_struct_to_json(alive)

	for {
		Udp_interface_send(destination_ip, chan_is_alive) //burde for-løkken heller implementeres i laget over?
		chan_is_alive <- StandardData.IP
		time.Sleep(time.Second)
	}
}

func udp_receive_is_alive(chan_received_msg chan []byte, chan_is_alive chan string) {
	go Udp_interface_receive(chan_received_msg)

	received := <- chan_received_msg

	data := Udp_json_to_struct(received, 1024)
	chan_is_alive <- data.IP
}

//Only slaves
func udp_send_order_executed(order_nr int, chan_order_executed chan int) {
	order := StandardData{}
	order.order_executed = order_nr

	chan_order_executed <- Udp_struct_to_json(order)

	Udp_interface_send(udp_get_master_ip(), chan_order_executed)
}

//Only master
func udp_receive_order_executed(chan_order_executed chan int, chan_received_msg chan []byte) {
	go Udp_interface_receive(chan_received_msg)

	received := <- chan_received_msg
	data := Udp_json_to_struct(received, 1024)

	chan_order_executed <- data.order_executed
}

//Only master
func udp_send_descendant_nr(chan_descendant_nr chan int, descendant_nr int, dest_ip string) {
	number := StandardData{}
	number.descendant_nr = descendant_nr
	chan_descendant_nr <- Udp_struct_to_json(number)
	Udp_interface_send(dest_ip, chan_descendant_nr)
}

//Only slave
func udp_receive_descendant_nr(chan_descendant_nr chan int, chan_received_msg chan []byte) {
	go Udp_interface_receive(chan_received_msg)

	received := <- chan_received_msg
	data := Udp_json_to_struct(received, 1024)

	chan_descendant_nr <- data.descendant_nr
}

//Only master
func udp_send_new_order(new_order NewOrder, dest_ip string, chan_new_order chan NewOrder) {
	order := StandardData{}
	order.new_order = new_order
	chan_new_order <- Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, chan_new_order)
}

//Only slave
func udp_receive_new_order(chan_new_order chan NewOrder, chan_received_msg chan []byte) {
	go Udp_interface_receive(chan_received_msg)
	received := <- chan_new_order
	data := Udp_json_to_struct(received, 1024)
	chan_new_order <- data.new_order
}


func udp_send_local_order(local_order LocalOrder, dest_ip string, chan_local_order chan LocalOrder) {
	order := StandardData{}
	order.local_order = local_order
	chan_local_order <- Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, chan_local_order)
}

//Only master
func udp_receive_local_order(chan_local_order chan LocalOrder, chan_received_msg chan []byte) {
	go Udp_interface_receive(chan_received_msg)
	received := <- chan_local_order
	data := Udp_json_to_struct(received, 1024)
	chan_local_order <- data.local_order
}



