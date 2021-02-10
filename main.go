package main

var TestJson TestJsonT

func main() {
	var path string
	path = "config.json"

	err := LoadConfig(path, &TestJson)

	if err != nil {
		//
	}
}