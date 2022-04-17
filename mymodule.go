package mymodule

import "fmt"

const version = "2.0.0"

func Version() {
	fmt.Println("mymodule version =", version)
}
