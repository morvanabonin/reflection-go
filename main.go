package main

import (
	"fmt"
	"log"
)

var TestJson TestJsonT

func main() {

	path := "config.json"

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered by panic in", r)
		}
	}()

	err := LoadConfig(path, &TestJson)

	if err != nil {
		log.Printf("Erro ao fazer load da config %s", err)
	}
}