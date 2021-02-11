package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type (
	// Config faz a load do arquivo json fazendo um unmarshall
	// Também faz o gerenciamento de acordo as tags passadas
	Config interface {

		// LoadConfig faz o load e o unmarshall do arquivo e coloca-o na estrutura passada
		//
		// LoadConfig("config.json", &Struct)
		//
		LoadConfig(path string, typ interface{}) error
	}
)

func LoadConfig(path string, c interface{}) (err error) {
	// validação para ver se o path está vazio
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
	parserConfig(&t, c, 0)

	return
}

func parserConfig(t *reflect.Value, c interface{}, depth int)  {
	// switch do kind para saber tipo
	switch t.Kind() {
	case reflect.Ptr:
		// se é um elemento do tipo ponteiro, chama a parseConfig novamente, descendo um nível
		temp := t.Elem()
		parserConfig(&temp, c, depth+1)
	case reflect.Struct:
		// Itera sobre a árvore json
		s := t
		jsonTreeIterate(t, c, s)
	}
}

func jsonTreeIterate(r *reflect.Value, c interface{}, s *reflect.Value) {
	t := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		if f.Kind().String() == "struct" {
			node := reflect.Indirect(f)
			jsonTreeIterate(&node, c, s)
		}

		field := t.Field(i)
		if content, ok := field.Tag.Lookup("context"); ok {
			switch content {
			case "global":
				
			case "env":
				value := r.FieldByName(t.Field(i).Name)

				// pegar o valor do campo da env
				env := os.Getenv(value.String())

				// se estiver vazia diz o campo que tentou buscar, para não ter log de dados que possam ser sensíveis
				if env == "" {
					fmt.Println("Valor da env, encontra-se vazio! Valor do atributo para conferência: ",
						value.String())
				}

				if !r.CanSet() {
					fmt.Printf("Não é possível setar valores na struct")
				}

				r.FieldByName(t.Field(i).Name).SetString(env)
			}
		}
	}
}



