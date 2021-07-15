package flat

import (
	"reflect"
	"testing"
)

func TestFlatMap(t *testing.T) {
	in := map[string]interface{}{
		"a1": map[string]interface{}{
			"a2": map[string]interface{}{
				"a31": "31",
				"a32": "32",
			},
		},
		"b1": map[string]interface{}{
			"b21": "b21",
			"b22": "b22",
		},
	}

	res := FlatMap(in, "")

	if res["a1.a2.a31"] != "31" || res["a1.a2.a32"] != "32" {
		t.Fatal("FlatMap error")
	}
	if res["b1.b21"] != "b21" || res["b1.b22"] != "b22" {
		t.Fatal("FlatMap error")
	}

	out := StereoMap(res)
	if !reflect.DeepEqual(in, out) {
		t.Fatal("StereoMap error")
	}
}
