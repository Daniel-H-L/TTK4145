package Network

type NewOrder struct {
	Floor        int //0-3
	Direction    int //0-2
	Is_new_order int
	Is_executed  int
	In_progress   int
}

type Queue struct{  
	Orders    [3][4]int
	Direction int
	Floor     int
}

type Backup struct {
	MainQueue map[Queue]string
}

type StandardData struct {
	IP            string
	Is_alive      bool
	Descendant_nr int
	Order         NewOrder
	Last_floor    int
	Dir           int
	Main_queue    Backup
}
