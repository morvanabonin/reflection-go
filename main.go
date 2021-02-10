package main

import (
	. "github.com/morvanabonin/reflection-go/config"
)

var TestJson TestJsonT

func main() {
	var path string
	path = "config.json"

	err := LoadConfig(path, &TestJson)

	if err != nil {
		//
	}
}