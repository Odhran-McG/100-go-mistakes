package main

import (
	"fmt"

	"github.com/Odhran-McG/100-go-mistakes/src/ch2/3-misusing-init-functions/redis"
)

// imported packages are initialised in order of import

// multiple init in same source file, executed in order of definition
func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init")
}

// main depends on redis, so redis package's init is initialised first so its init is executed
// next main init functions are executed, and finally the main() function is executed.

func main() {
	err := redis.Store("key", "value")
	_ = err
}
