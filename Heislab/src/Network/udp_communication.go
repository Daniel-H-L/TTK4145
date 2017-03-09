package Network

import (
	"net"
	"strings"
	"time"
)

var localIP string

func Udp_get_local_ip() (string, error) {
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

//Only master
func Udp_broadcast(IP string) {
	send_object := StandardData{}
	send_object.IP = IP

	send := Udp_struct_to_json(send_object)

	Udp_interface_bcast(send)
	time.Sleep(200 * time.Millisecond)
}

func Udp_send_is_alive(destination_ip string) {
	alive := StandardData{}
	alive.IP, _ = Udp_get_local_ip()
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
			if received != nil {
				data := Udp_json_to_struct(received)
				chan_is_alive <- data.IP
			} else {
				chan_is_alive <- ""
			}
		case err := <-chan_local_err:
			chan_error <- err
			return
		}

	}
}

//Only slaves
func Udp_send_order_status(status NewOrder, masterIP string) {
	order := StandardData{}
	order.Order = status

	send := Udp_struct_to_json(order)

	Udp_interface_send(masterIP, send)
}

//Only master
func Udp_receive_order_status(chan_order_status chan NewOrder, chan_received_msg chan []byte, portNr string, chan_error chan error, chan_source_ip chan string) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			if received != nil {
				data := Udp_json_to_struct(received)
				chan_order_status <- data.Order
				chan_source_ip <- data.IP
			}
		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

//Only master
func Udp_send_descendant_nr(descendant_nr int, dest_ip string) {
	number := StandardData{}
	number.Descendant_nr = descendant_nr
	send := Udp_struct_to_json(number)
	Udp_interface_send(dest_ip, send)
}

//Only slave
func Udp_receive_descendant_nr(chan_descendant_nr chan int, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)
			chan_descendant_nr <- data.Descendant_nr

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

//Only master
func Udp_send_new_order(new_order NewOrder, dest_ip string) {
	order := StandardData{}
	order.Order = new_order
	order.IP,_ = Udp_get_local_ip()
	send := Udp_struct_to_json(order)
	Udp_interface_send(dest_ip, send)
}

//Only slave
func Udp_receive_new_order(chan_new_order chan NewOrder, chan_received_msg chan []byte, portNr string, chan_error chan error, chan_source_ip chan string) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)
			chan_new_order <- data.Order
			chan_source_ip <- data.IP

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}

func Udp_send_state(dir int, floor int, dest_ip string) {
	state := StandardData{}
	state.Dir = dir
	state.Last_floor = floor
	send := Udp_struct_to_json(state)
	Udp_interface_send(dest_ip, send)
}

func Udp_receive_state(chan_elev_state chan NewOrder, chan_received_msg chan []byte, portNr string, chan_error chan error) {
	chan_local_err := make(chan error, 1)
	go Udp_interface_receive(chan_received_msg, portNr, chan_local_err)

	for {
		select {
		case received := <-chan_received_msg:
			data := Udp_json_to_struct(received)
			chan_elev_state <- data.Order

		case err := <-chan_local_err:
			chan_error <- err
			return
		}
	}
}
