package Network

import (
	"encoding/json"
	"fmt"
)

func Udp_struct_to_json(struct_object StandardData) []byte {

	fmt.Println("Before json: ", struct_object)
	json_object, _ := json.Marshal(struct_object)
	fmt.Println("After json: ", string(json_object))

	return json_object
}

func Udp_json_to_struct(json_object []byte) StandardData {

	//var t json.RawMessage
	//var o json.RawMessage
	struct_object := StandardData{}

	fmt.Println(string(json_object))

	json.Unmarshal(json_object, &struct_object)
	/*
		var Order NewOrder
		json.Unmarshal(o, &Order)

		fmt.Println("Order: ", struct_object)

		var T int
		json.Unmarshal(t, &T)
	*/
	return struct_object
}
