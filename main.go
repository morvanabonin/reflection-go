package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

var TestJson TestJsonT

func main() {
	var path string
	path = "main.json"

	err := LoadConfig(path, &TestJson)

	if err != nil {
		//
	}
}

func LoadConfig(path string, c interface{}) (err error) {
	//validação para ver se o path está vazio
	if path == "" {
		fmt.Println("O caminho do arquivo passado está vazio.")
	}

	bytesFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("O caminho do arquivo passado está vazio.")
	}

	err = json.Unmarshal(bytesFile, &c)
	t := reflect.ValueOf(c)
	parserConfig(t, 0)

	return
}

func parserConfig(t reflect.Value, depth int)  {
	// switch do kind para saber tipo
	switch t.Kind() {
	case reflect.Ptr:
		//se é um elemento do tipo ponteiro, chama a parseConfig novamente, descendo um nível
		parserConfig(t.Elem(), depth+1)
	case reflect.Struct:
		// Itera sobre a árvore json
		jsonTreeIterate(t)
	}
}

func jsonTreeIterate(r reflect.Value) {
	t := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		if f.Kind().String() == "struct" {
			node := reflect.ValueOf(f.Interface())
			jsonTreeIterate(node)
		}

		field := t.Field(i)
		if content, ok := field.Tag.Lookup("context"); ok {
			switch content {
			case "global":
				// fmt.Println("Global: Faço algo, criar função")
			case "env":
				// pega o nome do campo e o valor dele
				nameEl := r.FieldByName(t.Field(i).Name)
				// pegar o valor do campo da env
				env := os.Getenv(nameEl.String())

				// se estiver vazia diz o campo que tentou buscar, para não ter log de dados que possam ser sensíveis
				if env == "" {
					fmt.Println("Valor da env, encontra-se vazio! Valor do atributo para conferência:",
						nameEl.String())
				}
			}
		}
	}
}

