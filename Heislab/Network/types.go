package Network

type NewOrder struct {
	floor        int //0-3
	direction    int //0-2
	is_inside    bool
	is_new_order bool
	is_executed  bool
	in_progess   bool
}

type StandardData struct {
	IP            string
	Msg_ID        string
	Is_alive      bool
	Descendant_nr int
	Order         NewOrder
	Last_floor    int
	//dir
	//backup ???
}
