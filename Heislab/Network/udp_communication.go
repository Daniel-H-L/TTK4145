package Network

import (
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

func udp_get_master_ip(chan_is_master chan bool) string {
	is_master := <-chan_is_master

	if is_master {
		masterIP, _ = udp_get_local_ip()
	}
	return masterIP
}

//Only master
func Udp_broadcast(msg_id string) {
	send_object := StandardData{}
	send_object.IP, _ = udp_get_local_ip()
	send_object.msg_ID = msg_id

	send := Udp_struct_to_json(send_object)

	Udp_interface_bcast(send)
	time.Sleep(200 * time.Millisecond)
}

func Udp_send_is_alive(destination_ip string) {
	alive := StandardData{}
	alive.IP, _ = udp_get_local_ip()
	send := Udp_struct_to_json(alive)

	for {
		Udp_interface_send(destination_ip, send)
		time.Sleep(200 * time.Millisecond)
	}
}

func Udp_receive_is_alive(chan_received_msg chan []byte, chan_is_alive chan string, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_error)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)
			chan_is_alive <- data.IP //Trenger vi dette???
		case err := <-chan_local_err:
			chan_error <- err
			return
		}

	}
}

//Only slaves
func udp_send_order_executed(order_nr int) {
	order := StandardData{}
	order.order_executed = order_nr

	send := Udp_struct_to_json(order)

	Udp_interface_send(udp_get_master_ip(), send)
}

//Only master
func udp_receive_order_executed(chan_order_executed chan int, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)
			//chan_order_executed <- data.Order_executed

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

//Only master
func udp_send_descendant_nr(chan_descendant_nr chan int, descendant_nr int, dest_ip string) {
	number := StandardData{}
	number.descendant_nr = descendant_nr
	send := Udp_struct_to_json(number)
	Udp_interface_send(dest_ip, send)
}

//Only slave
func udp_receive_descendant_nr(chan_descendant_nr chan int, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

//Only master
func udp_send_new_order(new_order NewOrder, dest_ip string, chan_new_order chan NewOrder) {
	order := StandardData{}
	order.new_order = new_order
	send := Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, send)
}

//Only slave
func udp_receive_new_order(chan_new_order chan NewOrder, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

func udp_send_local_order(local_order LocalOrder, dest_ip string, chan_local_order chan LocalOrder) {
	order := StandardData{}
	order.local_order = local_order
	send := Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, chan_local_order)
}

//Only master
func udp_receive_local_order(chan_local_order chan LocalOrder, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}
