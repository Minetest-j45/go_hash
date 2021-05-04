package main

import (
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "strings"
)

const multiplierInit = 716

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Not enough arguments!")
    }

    input := strings.Join(os.Args[1:], " ") 
    ba := []byte(input)

    fmt.Println(ba)
    
    before := int32(multiplierInit)
    
    for b := range ba {
        ba[b] = (ba[b] % (multiplierInit * before) )* before
        before = ba[b]
    }

    fmt.Println(ba)
    fmt.Println(hex.EncodeToString(ba))
}
