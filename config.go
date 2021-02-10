package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func LoadConfig(path string, c interface{}) (err error) {
	//validação para ver se o path está vazio
	if path == "" {
		fmt.Println("O caminho do arquivo passado está vazio.")
	}

	bytesFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("O caminho do arquivo passado está vazio.")
	}

	// chamada do unmarshal com o bytes do arquivo e o endereço de memória da estrutura
	err = json.Unmarshal(bytesFile, c)
	t := reflect.ValueOf(c)
	parserConfig(&t, 0)

	return
}

func parserConfig(t *reflect.Value, depth int)  {
	// switch do kind para saber tipo
	switch t.Kind() {
	case reflect.Ptr:
		//se é um elemento do tipo ponteiro, chama a parseConfig novamente, descendo um nível
		temp := t.Elem()
		parserConfig(&temp, depth+1)
	case reflect.Struct:
		// Itera sobre a árvore json
		jsonTreeIterate(t)
	}
}

func jsonTreeIterate(r *reflect.Value) {
	fmt.Println(r)
	t := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		if f.Kind().String() == "struct" {
			node := reflect.Indirect(f)
			jsonTreeIterate(&node)
		}

		field := t.Field(i)
		if content, ok := field.Tag.Lookup("context"); ok {
			switch content {
			case "global":
				// fmt.Println("Global: Faço algo, criar função")
			case "env":
				// pega o nome do campo e o valor dele
				value := r.FieldByName(t.Field(i).Name)

				// pegar o valor do campo da env
				env := os.Getenv(value.String())

				// se estiver vazia diz o campo que tentou buscar, para não ter log de dados que possam ser sensíveis
				if env == "" {
					fmt.Println("Valor da env, encontra-se vazio! Valor do atributo para conferência:",
						value.String())
				}
				r.FieldByName(t.Field(i).Name).SetString(env)
			}
		}
	}
}



