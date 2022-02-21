package strcut2map

import "encoding/json"

func jsonTrans(obj interface{}) (m map[string]interface{}) {
	str, _ := json.Marshal(obj)
	json.Unmarshal(str, &m)
	return
}
