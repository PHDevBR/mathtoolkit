package main

import (
    "fmt"
    "os"
    "extras"
)

func main() {
    args := os.Args 
    if (args[1] == "calculator") {
        extras.returnme()
    }
}