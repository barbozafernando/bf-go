package interpreter

import (
	"fmt"
	"log"
	"os"
)

type Interpreter struct {
    filename string
    tape []byte
}

func Run(filename string, tape []byte) {
    fmt.Printf("Running interpreter")
    f, err := os.OpenFile(filename, os.O_RDONLY, 0644)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error when opening the file")
    }

    b := make([]byte, 1024)
    for {
        n, err := f.Read(b)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error when reading the file")
        }

        if n == 0
            break

        if _, err := fo.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
}
