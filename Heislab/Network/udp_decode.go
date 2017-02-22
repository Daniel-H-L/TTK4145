package Network

import (
	"encoding/json"
)

func Udp_struct_to_json(struct_object StandardData) []byte {
	json_object, _ := json.Marshal(struct_object)

	return json_object
}

func Udp_json_to_struct(json_object []byte) StandardData {
	struct_object := StandardData{}
	json.Unmarshal(json_object, &struct_object)

	return struct_object
}
