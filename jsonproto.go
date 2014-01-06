// A tool to generate Go struct by JSON example.
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

func camelCase(s string) (r string) {
	for _, part := range strings.Split(s, "_") {
		r += strings.Title(part)
	}
	return
}

func emit(s string) {
	fmt.Print(s)
}

func tab(depth int) string {
	return strings.Repeat("\t", depth)
}

func translate(v interface{}, depth int) {
	// don't know how to type switch nil yet.
	if v == nil {
		emit("interface{}")
		return
	}
	switch v := v.(type) {
	case bool:
		emit("bool")
	case float64:
		if math.Floor(v) == v {
			emit("int64")
		} else {
			emit("float64")
		}
	case string:
		emit("string")
	case []interface{}:
		if len(v) == 0 {
			emit("[]interface{}")
		} else {
			emit("[]")
			translate(v[0], depth)
		}
	case map[string]interface{}:
		emit("struct {\n")
		for k, v := range v {
			camel := camelCase(k)
			emit(tab(depth) + camel + " ")
			translate(v, depth+1)
			if strings.ToLower(camel) != strings.ToLower(k) {
				emit(" `json:\"" + k + "\"`")
			}
			emit("\n")
		}
		emit(tab(depth-1) + "}")
		if depth == 1 {
			emit("\n")
		}
	default:
		panic(fmt.Sprintf("unknown type: %+v", v))
	}
}

func main() {
	decoder := json.NewDecoder(os.Stdin)
	var v interface{}
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	translate(v, 1)
}
