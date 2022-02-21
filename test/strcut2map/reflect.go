package strcut2map

import "reflect"

func reflectTrans(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{}, 0)

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return m
}
