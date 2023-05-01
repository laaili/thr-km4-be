package main

import (
	"fmt"
	"strings"
)

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	keys := strings.Split(key, ",")
	for _, n := range keys {
		delete(dataMap, n)
	}
	return dataMap
}

func main() {
	data := map[string]any{
		"run"		:	"lari",
		"jump"		:	"loncat",
		"swim"		:	"berenang",
	}
	m := removeUnrelated(data, "jump")
	fmt.Println(m)
}
