package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) > 1 {
        args := os.Args[1:]
        text := strings.Join(args, " ")
        fmt.Println(text, "Rocks!!")
    } else {
        fmt.Println("Full Cycle Rocks!!")
    }

}