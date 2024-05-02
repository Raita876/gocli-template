package hello

import (
	"fmt"
	"strings"
)

func Hello(target string, exclNum uint8) error {
	s, err := makeHelloMessage(target, exclNum)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", s)

	return nil
}

func makeHelloMessage(target string, exclNum uint8) (string, error) {
	exclMsg := strings.Repeat("!", int(exclNum))
	return fmt.Sprintf("Hello %s%s", target, exclMsg), nil
}
