package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	tests := []struct {
		target string
	}{
		{"world"},
	}

	for _, tt := range tests {
		err := Hello(tt.target)
		assert.Nil(t, err, "Hello failed.")
	}

}

func TestMakeHelloMessage(t *testing.T) {
	tests := []struct {
		target string
		want   string
	}{
		{"world", "Hello World!"},
	}

	for _, tt := range tests {
		got, err := makeHelloMessage(tt.target)
		assert.Nil(t, err, "makeHelloMessage failed.")
		assert.Equal(t, got, tt.want, "they should be equal.")
	}
}
