package main

type Adventure struct {
	Title string `json:"title" context:"global"`
	Autor string `json:"autor"`
	Year  uint8 `json:"year"`
}

type Romance struct {
	Title string `json:"title"`
	Autor string `json:"autor"`
	Year  uint8 `json:"year"`
}

type Book struct {
	Adventure Adventure `json:"adventure"`
	Romance Romance `json:"romance"`
}

type Food struct {
	Like string `json:"like" context:"env"`
	Dislike string `json:"dislike"`
}

type TestJsonT struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
	Food Food `json:"food"`
	Book Book `json:"book"`
}
