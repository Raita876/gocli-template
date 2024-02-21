package hello

import "fmt"

func Hello() error {
	s, err := getHelloMessage()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", s)

	return nil
}

func getHelloMessage() (string, error) {
	return "Hello World!", nil
}
