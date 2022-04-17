package mymodule

import "fmt"

const version = "1.0.0"

func Version() {
   fmt.Println("mymodule version =", version)
}
