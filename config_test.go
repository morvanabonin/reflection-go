package main

import (
	"testing"
)

func Test_Config(t *testing.T) {
	tests := []struct {
		name    string
		inspect func(t *testing.T)
	}{
		{
			name: "GetElement",
			inspect: func(t *testing.T) {
				path := "config.json"
				structure := new(TestJsonT)
				err := LoadConfig(path, structure)

				if err != nil {
					t.Errorf("Erro ao fazer load do arquivo")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.inspect != nil {
				tt.inspect(t)
			}
		})
	}
}
