package main

import (
	"bytes"
	"testing"
)

func TestHex2bytes(t *testing.T) {
	var tests = []struct {
		input string
		want  []byte
	}{
		{"FF", []byte{255}},
		{"FFFF", []byte{255, 255}},
		{"EEEEEE", []byte{238, 238, 238}},
	}

	for _, test := range tests {
		want := test.want
		actual := hex2bytes(test.input)
		if !bytes.Equal(actual, test.want) {
			t.Errorf("want %v, but get %v", want, actual)
		}
	}
}
