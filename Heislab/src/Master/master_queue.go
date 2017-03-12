package Master

import (
	//"../DriveElevator"
	"../Network"
	"math"
)

func Costfunction(button int, floor_order int, backup map[string]*Network.Backup, elevatorIP string) int {
	cost := 0

	state = backup[elevatorIP].State
	dir = backup[elevatorIP].Direction
	floor = backup[elevatorIP].Floor
	internalQueue = backup[elevatorIP].Orders

	for ip := range backup {

		if state == 0 && ((dir == -1 && button == 1) || (dir == 1 && button == 0)) && floor == floor_order {
			return cost
		}

		cost_punish_wrong_dir := 3

		switch dir {
		case 1:
			if button == 1 {
				cost += cost_punish_wrong_dir
			}

		case -1:
			if button == 0 {
				cost += cost_punish_wrong_dir
			}
		}

		// for q := range backup.MainQueue {
		// 	if q == IP {
		// 		internal_queue = q
		// 	}
		// }

		target_direction := floor_order - floor

		if target_direction > 0 && dir == 1 || dir == 0 {
			for f := floor; f < floor_order || f == 3; f++ {
				if internalQueue[button][f] == 1 || internalQueue[2][f] == 1 {
					cost++
				}
				//cost++
			}
		}

		if target_direction < 0 && dir == -1 || dir == 0 {
			for f := floor; f > floor || f == 0; f-- {
				if internalQueue[button][f] == 1 || internalQueue[2][f] == 1 {
					cost++
				}
				//cost++
			}
		}
	}

	cost_punish_floor_difference := 3

	cost += cost_punish_floor_difference * (int(math.Abs(float64(floor - floor_order))))
	return cost
}

func desideElevator(button int, floor int, backup map[string]*Network.Backup) string {
	minElevCost := -1
	var elevator string

	for ip := range backup {
		if backup[ip].State != -1 {
			currentMinCost := Costfunction(button, floor, backup, ip)

			if minElevCost == -1 || minElevCost > currentMinCost {
				minElevCost = currentMinCost
				elevator = ip
			}
		}
	}
	return elevator
}

//func Delegate_order(order Network.NewOrder, chan_backup Network.Backup) string {
//button := order.Button
//floor := order.Floor

// lowest_slave_cost := -1 //variabelnavn?
// for s := range master.Slaves {
// 	current_slave_cost := Costfunction(order_direction, floor_order, s.Last_floor, s.Direction, s.IP, backup)
// 	if lowest_slave_cost == -1 || lowest_slave_cost > current_slave_cost {
// 		lowest_slave_cost := current_slave_cost
// 		slave_cost_IP := s.IP
// 	}
// }

// master_cost := Costfunction(order_direction, floor_order, master.Last_floor, master.Direction, master.IP, backup)
// if master_cost < lowest_slave_cost {
// 	return master.IP
// } else {
// 	return slave_cost_IP
// }
//return "129.241.187.141"
