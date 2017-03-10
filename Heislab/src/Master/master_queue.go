package Master

import (
	//"../DriveElevator"
	"../Network"
	"math"
)

func Costfunction(button int, floor_order int, current_floor int, direction int, source_IP string, backup Network.Backup, state int) int {
	cost := 0
	internal_queue := [3][4]int{}
	for ip, order := range backup.MainQueue{
		if ip==source_IP{
			internal_queue=order.Orders
		}
	}

	if state == 0 && ((direction == -1 && button == 1) || (direction == 1 && button == 0)) && current_floor == floor_order {
		return cost
	}

	cost_punish_wrong_dir := 3

	switch direction {
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

	target_direction := floor_order - current_floor

	if target_direction > 0 && direction == 1 || direction == 0 {
		for floor := current_floor; floor < floor_order || floor == 3; floor++ {
			if internal_queue[button][floor] == 1 || internal_queue[2][floor] == 1 {
				cost++
			}
			//cost++
		}
	}

	if target_direction < 0 && direction == -1 || direction == 0 {
		for floor := current_floor; floor > current_floor || floor == 0; floor-- {
			if internal_queue[button][floor] == 1 || internal_queue[2][floor] == 1 {
				cost++
			}
			//cost++
		}
	}
	cost_punish_floor_difference := 3

	cost += cost_punish_floor_difference * (int(math.Abs(float64(current_floor - floor_order))))
	return cost
}

func Delegate_order(order Network.NewOrder, chan_backup Network.Backup) string {
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
	return "129.241.187.141"
}
