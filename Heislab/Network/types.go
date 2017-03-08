package Network

type NewOrder struct {
	floor        int //0-3
	direction    int //0-2
	is_inside    int
	is_new_order bool
	is_executed  bool
	in_progess   bool
}

type Queue struct {
	Orders    [3][4]int
	Direction int
	Floor     int
}

type Backup struct {
	MainQueue [Queue]string
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
