package EventManager

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var conn net.Conn
var mutex = &sync.Mutex{}

const (
	SIM_SERV_ADDR   = "127.0.0.1:15657"
	USING_SIMULATOR = false
)

func Elevator_init() int {
	if USING_SIMULATOR {

		fmt.Println("Mode: USING_SIMULATOR")

		tcpAddr, err := net.ResolveTCPAddr("tcp", SIM_SERV_ADDR)
		if err != nil {
			fmt.Println("ResolveTCPAddr failed:", err.Error())
			log.Fatal(err)
		}
		fmt.Println("ResolveTCPAddr success")

		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Println("Dial failed:", err.Error())
			log.Fatal(err)
		}
		fmt.Println("Dial success")
		return 1

	}

	if !USING_SIMULATOR {
		return IoInit()
	}

	return 0
}

// ----------------------------------------------------------
// -------------------- Inputs ------------------------------
// ----------------------------------------------------------

func Elevator_get_floor_sensor_signal() int {
	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{7, byte(0), byte(0), byte(0)})
		mutex.Unlock()
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}

		buffer := make([]byte, 4)
		mutex.Lock()
		conn.Read(buffer)
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if buffer[1] == 1 {
			return int(buffer[2])
		} else {
			return -1
		}
	}

	if !USING_SIMULATOR {
		if IoReadBit(SENSOR_FLOOR1) {
			return 0
		} else if IoReadBit(SENSOR_FLOOR2) {
			return 1
		} else if IoReadBit(SENSOR_FLOOR3) {
			return 2
		} else if IoReadBit(SENSOR_FLOOR4) {
			return 3
		} else {
			return -1
		}
	}
	return -1
}

func Elevator_get_button_signal(button int, floor int) int {
	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{6, byte(button), byte(floor), byte(0)})
		mutex.Unlock()
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}

		buffer := make([]byte, 4)
		mutex.Lock()
		conn.Read(buffer)
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if buffer[1] == 1 {
			return 1
		} else {
			return 0
		}
	}

	if !USING_SIMULATOR {
		if floor < 0 || floor >= 4 {
			log.Printf("Error: Floor %d out of range!\n", floor)
			return 0
		}
		if button < 0 || button >= 3 {
			log.Printf("Error: Button %d out of range!\n", button)
			return 0
		}
		if button == 0 && floor == 3 {
			log.Println("Button up from top floor does not exist!")
			return 0
		}
		if button == 1 && floor == 0 {
			log.Println("Button down from ground floor does not exist!")
			return 0
		}

		if IoReadBit(buttonChannelMatrix[floor][button]) {
			return 1
		} else {
			return 0
		}
	}
	return 0
}

// ----------------------------------------------------------
// -------------------- Outputs -----------------------------
// ----------------------------------------------------------

func Elevator_set_motor_dir(dir int) {
	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{1, byte(dir), byte(0), byte(0)})
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}
	}

	if !USING_SIMULATOR {
		if dir == 0 {
			IoWriteAnalog(MOTOR, 0)
		} else if dir < 0 {
			IoSetBit(MOTORDIR)
			IoWriteAnalog(MOTOR, 2800)
		} else if dir > 0 {
			IoClearBit(MOTORDIR)
			IoWriteAnalog(MOTOR, 2800)
		}
	}
}

func Elevator_set_floor_indicator(floor int) {
	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{3, byte(floor), byte(0), byte(0)})
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}
	}

	if !USING_SIMULATOR {
		if floor < 0 || floor > 3 {
			IoClearBit(LIGHT_FLOOR_IND1)
			IoClearBit(LIGHT_FLOOR_IND2)
		} else if floor == 0 {
			IoClearBit(LIGHT_FLOOR_IND1)
			IoClearBit(LIGHT_FLOOR_IND2)
		} else if floor == 2 {
			IoSetBit(LIGHT_FLOOR_IND1)
			IoClearBit(LIGHT_FLOOR_IND2)
		} else if floor == 1 {
			IoClearBit(LIGHT_FLOOR_IND1)
			IoSetBit(LIGHT_FLOOR_IND2)
		} else {
			IoSetBit(LIGHT_FLOOR_IND1)
			IoSetBit(LIGHT_FLOOR_IND2)
		}
	}

}

func Elevator_set_button_lamp(b int, f int, val int) {

	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{2, byte(b), byte(f), byte(val)})
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}
	}
	if !USING_SIMULATOR {
		if val == 1 {
			IoSetBit(lightChannelMatrix[f][b])
		} else {
			IoClearBit(lightChannelMatrix[f][b])
		}
	}

}

func Elevator_set_door_open_lamp(val int) {

	if USING_SIMULATOR {
		mutex.Lock()
		_, err := conn.Write([]byte{4, byte(val), byte(0), byte(0)})
		mutex.Unlock()
		time.Sleep(5 * time.Millisecond)
		if err != nil {
			fmt.Println("Write to server failed:", err.Error())
			log.Fatal(err)
		}
	}
	if !USING_SIMULATOR {
		if val == 1 {
			IoSetBit(LIGHT_DOOR_OPEN)
		} else {
			IoClearBit(LIGHT_DOOR_OPEN)
		}
	}

}
