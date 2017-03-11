package Network

type NewOrder struct {
	Floor        int //0-3
	Button       int //0-2
	Is_new_order int
	Is_executed  int
	In_progress  int
}

type Queue struct {
	Orders    [3][4]int
	Direction int
	Floor     int
	State     int
}

type Backup struct {
	MainQueue map[string]*Queue
}

type ElevState struct {
	Direction int
	Floor     int
	State     int
}

type StandardData struct {
	Type int
	IP   string
	//Is_alive      bool
	Descendant_nr int
	Order         NewOrder
	Main_queue    Backup
	Status        ElevState
	Lights        [3][4]int
}

// type StandardData struct {
// 	Type int
// 	IP   string
// 	//Is_alive      bool
// 	Descendant_nr int
// 	Data          interface{}
// }
