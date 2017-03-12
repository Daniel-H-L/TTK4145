package Network

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("Bcast 1... ", IP)

	//send := Udp_struct_to_json(send_object)

	json_object, _ := json.Marshal(send_object)

	for {
		Udp_interface_bcast(json_object)
		fmt.Println("Broadcast 2...")
		time.Sleep(500 * time.Millisecond)
	}
}

func Udp_send_is_alive(destination_ip string) {
	alive := StandardData{}
	alive.IP, _ = Udp_get_local_ip()
	alive.Type = 0

	fmt.Println("for loop alive...")
	Udp_interface_send(destination_ip, alive)
	time.Sleep(200 * time.Millisecond)

}

func Udp_receive_is_alive(chan_received_msg chan StandardData, chan_is_alive chan string, portNr string, chan_kill chan bool) {
	chan_kill_receive := make(chan bool)
	fmt.Println("Communication.. alive")
	go Udp_interface_receive(chan_received_msg, ":40018", chan_kill_receive)

	for {
		select {
		case data := <-chan_received_msg:

			chan_is_alive <- data.IP
			time.Sleep(500 * time.Millisecond)
		case kill := <-chan_kill:
			fmt.Println("Udp _receive_is_alive received kill")
			chan_kill_receive <- kill
			return
		}

	}
}

//Only slaves
// func Udp_send_order_status(status NewOrder, masterIP string) {
// 	order := StandardData{}
// 	order.Order = status
// 	order.IP, _ = Udp_get_local_ip()
// 	order.Type = 3

// 	Udp_interface_send(masterIP, order)
// }

func Udp_receive_standard_data(portNr string, chan_source_ip chan string, chan_descendant_nr chan int64, chan_order_dir chan int64, chan_order_floor chan int64, chan_kill chan bool) {
	chan_kill_receive := make(chan bool)

	chan_received_msg := make(chan StandardData, 100)
	fmt.Println("Received sd...", portNr)
	go Udp_interface_receive(chan_received_msg, ":30018", chan_kill_receive)

	for {
		select {
		case data := <-chan_received_msg:
			switch data.Type {
			case 0:
				chan_source_ip <- data.IP
				fmt.Println("Type er null", data.IP)
			case 1:
				chan_descendant_nr <- data.DescendantNr
				fmt.Println("Descendant ")
			case 2:
				chan_order_dir <- data.OrderDir
				chan_order_floor <- data.OrderFloor
			}

		case <-chan_kill:
			fmt.Println("Killing slave network")
			chan_kill_receive <- true
			return
		}
	}
}

//Only master
func Udp_send_descendant_nr(descendant_nr int64, dest_ip string) {
	number := StandardData{}
	number.DescendantNr = descendant_nr
	number.IP, _ = Udp_get_local_ip()
	number.Type = 1
	//send := Udp_struct_to_json(number)
	Udp_interface_send(dest_ip, number)
}

//Only master
func Udp_send_new_order(dir int64, floor int64, dest_ip string) {
	order := StandardData{}
	order.OrderDir = dir
	order.OrderFloor = floor
	order.IP = "123.3.3.1"
	order.Type = 2
	//send := Udp_struct_to_json(order)

	Udp_interface_send(dest_ip, order)
}

// func Udp_send_state(status int, value int, dest_ip string) {
// 	state := StandardData{}

// 	if status == 1 {
// 		state.Status.Direction = value
// 	} else if status == 2 {
// 		state.Status.Floor = value
// 	} else if status == 3 {
// 		state.Status.State = value
// 	}

// 	state.IP, _ = Udp_get_local_ip()
// 	state.Type = 5
// 	//send := Udp_struct_to_json(state)
// 	Udp_interface_send(dest_ip, state)
// }

// func Udp_send_set_lights(set_lights [3][4]int, dest_ip string) {
// 	lights := StandardData{}
// 	lights.IP, _ = Udp_get_local_ip()
// 	lights.Lights = set_lights
// 	lights.Type = 6

// 	//send := Udp_struct_to_json(lights)
// 	Udp_interface_send(dest_ip, lights)
// }
