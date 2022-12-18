package main

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/jsonq"
)

func main() {

	data := []byte(`
	{
		"foo": 1,
		"bar": 2,
		"test": "Hello, world!",
		"baz": 123.1,
		"array": [
			{"foo": 1},
			{"bar": 2},
			{"baz": 3}
		],
		"subobj": {
			"foo": 1,
			"subarray": [1,2,3],
			"subsubobj": {
				"bar": 2,
				"baz": 3,
				"array": ["hello", "world"]
			}
		},
		"bool": true
	}
	`)

	mi := map[string]any{}
	err := json.Unmarshal(data, &mi)
	if err != nil {
		panic(err.Error())
	}

	jq := jsonq.NewQuery(mi)
	mprint(jq.Int("foo"))
	mprint(jq.Int("subobj", "subarray", "1"))

}

func mprint(arg ...any) {
	fmt.Println(arg...)
}
