package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	if err := Hello(); err != nil {
		t.Error(err)
	}
}

func TestGetHelloMessage(t *testing.T) {
	tests := []struct {
		want string
	}{
		{"Hello World!"},
	}

	for _, tt := range tests {
		got, err := getHelloMessage()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, got, tt.want, "they should be equal")
	}
}
