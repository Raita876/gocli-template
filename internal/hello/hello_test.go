package hello

import "testing"

func TestHello(t *testing.T) {
	if err := Hello(); err != nil {
		t.Error(err)
	}
}
