package Network

import (
	"encoding/json"
)

func Udp_struct_to_json(struct_object StandardData) []byte {
	json_object, _ := json.Marshal(struct_object)

	return json_object
}

func Udp_json_to_struct(json_object []byte, n int) StandardData {
	struct_object := StandardData{}
	json.Unmarshal(json_object[0:n], &struct_object)

	return struct_object
}
