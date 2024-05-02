package hello

import (
	"fmt"
	"log/slog"
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
	slog.Debug(fmt.Sprintf("target=%s", target))
	slog.Debug(fmt.Sprintf("exclNum=%d", exclNum))

	exclMsg := strings.Repeat("!", int(exclNum))
	return fmt.Sprintf("Hello %s%s", target, exclMsg), nil
}
