package main

import (
	"bf-go/interpreter"
)

func main() {
    t := new([]byte)
    interpreter.Run("test.txt", *t)
}
