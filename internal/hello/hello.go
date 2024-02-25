package hello

import "fmt"

func Hello(target string) error {
	s, err := makeHelloMessage(target)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", s)

	return nil
}

func makeHelloMessage(target string) (string, error) {
	return fmt.Sprintf("Hello %s\n", target), nil
}
