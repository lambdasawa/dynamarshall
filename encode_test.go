package main_test

import (
	"bytes"
	main "github.com/lambdasawa/dynamarshall"
	"testing"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		output  string
		wantErr bool
	}{
		{
			name:    "json",
			input:   `{"foo": "bar"}`,
			output:  `{"foo":{"B":null,"BOOL":null,"BS":null,"L":null,"M":null,"N":null,"NS":null,"NULL":null,"S":"bar","SS":null}}` + "\n",
			wantErr: false,
		},
		{
			name:  "json lines",
			input: `{"foo": "bar"}` + "\n" + `{"hoge": "fuga"}`,
			output: `{"foo":{"B":null,"BOOL":null,"BS":null,"L":null,"M":null,"N":null,"NS":null,"NULL":null,"S":"bar","SS":null}}` + "\n" +
				`{"hoge":{"B":null,"BOOL":null,"BS":null,"L":null,"M":null,"N":null,"NS":null,"NULL":null,"S":"fuga","SS":null}}` + "\n",
			wantErr: false,
		},
		{
			name:    "number",
			input:   `1`,
			output:  `{"B":null,"BOOL":null,"BS":null,"L":null,"M":null,"N":"1","NS":null,"NULL":null,"S":null,"SS":null}` + "\n",
			wantErr: false,
		},
		{
			name:    "string",
			input:   `"foo"`,
			output:  `{"B":null,"BOOL":null,"BS":null,"L":null,"M":null,"N":null,"NS":null,"NULL":null,"S":"foo","SS":null}` + "\n",
			wantErr: false,
		},
		{
			name:    "bool",
			input:   `true`,
			output:  `{"B":null,"BOOL":true,"BS":null,"L":null,"M":null,"N":null,"NS":null,"NULL":null,"S":null,"SS":null}` + "\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			reader := bytes.NewBufferString(tt.input)
			writer := bytes.NewBufferString("")

			err := main.Encode(reader, writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := writer.String(); got != tt.output {
				t.Errorf("Encode() \ngot = %v\nwant = %v", got, tt.output)
			}
		})
	}
}
