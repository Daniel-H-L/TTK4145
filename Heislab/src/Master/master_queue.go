package Master

// import (
// 	//"../DriveElevator"
// 	"../Network"
// 	"math"
// )

// func Master_costfunction(button_type int, floor_order int, current_floor int, direction int, IP string, backup_p *Network.Backup) int {
// 	backup := *backup_p
// 	cost := 0
// 	if state == STOP && ((direction == -1 && button_type == 1) || (direction == 1 && button_type == 0)) && current_floor == floor_order {
// 		return cost
// 	}

// 	cost_punish_wrong_dir := 3

// 	switch direction {
// 	case 1:
// 		if button_type == 1 {
// 			cost += cost_punish_wrong_dir
// 		}

// 	case -1:
// 		if button_type == 0 {
// 			cost += cost_punish_wrong_dir
// 		}
// 	}

// 	for q := range backup.MainQueue {
// 		if backup.MainQueue[q] == IP {
// 			internal_queue = q
// 		}
// 	}

// 	target_direction := floor_order - current_floor

// 	if target_direction > 0 && direction == 1 || direction == 0 {
// 		for floor := current_floor; floor < floor_order || floor == 3; floor++ {
// 			if internal_queue[buttion_type][floor] == 1 || internal_queue[2][floor] == 1 {
// 				cost++
// 			}
// 			//cost++
// 		}
// 	}

// 	if target_direction < 0 && direction == -1 || dir == 0 {
// 		for floor := current_floor; floor > current_floor || floor == 0; floor-- {
// 			if internal_queue[buttion_type][floor] == 1 || internal_queue[2][floor] == 1 {
// 				cost++
// 			}
// 			//cost++
// 		}
// 	}
// 	cost_punish_floor_difference := 3

// 	cost += cost_punish_floor_difference * (int(math.Abs(float64(current_floor - floor_order))))
// 	return cost
// }

// func (master *Master) Master_queue_delegate_order(order Network.NewOrder, backup *Network.Backup) string {
// 	order_direction = order.direction
// 	floor_order = order.floor

// 	lowest_slave_cost := -1 //variabelnavn?
// 	for s := range master.Slaves {
// 		current_slave_cost := master_costfunction(order_direction, floor_order, s.Last_floor, s.Direction, s.IP, backup)
// 		if lowest_slave_cost == -1 || lowest_slave_cost > current_slave_cost {
// 			lowest_slave_cost := current_slave_cost
// 			slave_cost_IP := s.IP
// 		}
// 	}

// 	master_cost := master_costfunction(order_direction, floor_order, master.Last_floor, master.Direction, master.IP, backup)
// 	if master_cost < lowest_slave_cost {
// 		return master.IP
// 	} else {
// 		return slave_cost_IP
// 	}
// }
