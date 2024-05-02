package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	tests := []struct {
		target  string
		exclNum uint8
	}{
		{"World", 1},
	}

	for _, tt := range tests {
		err := Hello(tt.target, tt.exclNum)
		assert.Nil(t, err, "Hello failed.")
	}

}

func TestMakeHelloMessage(t *testing.T) {
	tests := []struct {
		target  string
		exclNum uint8
		want    string
	}{
		{"World", 3, "Hello World!!!"},
	}

	for _, tt := range tests {
		got, err := makeHelloMessage(tt.target, tt.exclNum)
		assert.Nil(t, err, "makeHelloMessage failed.")
		assert.Equal(t, got, tt.want, "they should be equal.")
	}
}
