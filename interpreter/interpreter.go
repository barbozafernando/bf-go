package interpreter

import ("fmt")

type Interpreter struct {
    filename string
    tape []byte
}

func Run(filename string, tape []byte) {
    fmt.Printf("Running interpreter")
}
