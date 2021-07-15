package flat

import "strings"

const (
	dot = "."
)

func StereoMap(m map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{}, len(m))
	for k, v := range m {
		set(res, k, v)
	}
	return res
}

func FlatMap(m map[string]interface{}, pfx string) map[string]interface{} {
	res := make(map[string]interface{}, len(m))
	for key, value := range m {
		if pfx != "" {
			key = pfx + "." + key
		}
		switch v := value.(type) {
		case map[string]interface{}:
			for k1, v1 := range FlatMap(v, key) {
				res[k1] = v1
			}
		default:
			res[key] = value
		}
	}
	return res
}

func set(store map[string]interface{}, key string, value interface{}) {
	idx := strings.Index(key, dot)
	if idx == -1 {
		store[key] = value
		return
	}

	subKey := key[:idx]
	if _, ok := store[subKey]; !ok {
		store[subKey] = map[string]interface{}{}
	}

	if subMap, ok := store[subKey].(map[string]interface{}); ok {
		if idx+1 < len(key) {
			set(subMap, key[idx+1:], value)
		} else {
			set(subMap, "", value)
		}
	}
}
