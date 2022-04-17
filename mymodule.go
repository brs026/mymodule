package mymodule

import "fmt"

const version = "1.1.1"

func Version() {
	fmt.Println("mymodule version =", version)
}
