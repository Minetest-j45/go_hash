package main

import (
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "strings"
)

const multiplierInit = 62

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Not enough arguments!")
    }

    input := strings.Join(os.Args[1:], " ") 
    ba := []byte(input)
    
    before := uint8(multiplierInit)
    
    for b := range ba {
        ba[b] = (ba[b] % (multiplierInit * before) )* before
        if ba[b] <= 0 {
           ba[b] = before
        }
        before = ba[b]
    }
    fmt.Println(hex.EncodeToString(ba))
}
