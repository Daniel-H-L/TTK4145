package Network

type NewOrder struct {
	floor     int
	direction int //endre til C-definert variabeltype
	//priority  int
	order_nr int
}

type LocalOrder struct {
	is_inside_order bool
	floor           int
	direction       int //endre til C-typen
}

type StandardData struct {
	IP             string
	Msg_ID         string
	Is_alive       bool
	Order_executed int
	Descendant_nr  int
	New_order      NewOrder
	Local_order    LocalOrder
	Last_floor     int
	//dir
	//backup ???
}
