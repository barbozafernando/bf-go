package main

import (
	"bf-go/interpreter"
)

func main() {
    tape := new([]byte)
    interpreter.Run("hello.bf", *tape)
}
