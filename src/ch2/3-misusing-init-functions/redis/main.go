package redis

import "fmt"

// imports

func init() {
	fmt.Print("init redis package")
}

func Store(key, value string) error {
	// store something
	return nil
}
