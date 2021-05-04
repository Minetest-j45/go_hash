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
    ra := []rune(input)
    var ia []int32
    for _, v := range ra {
        ia = append(ia, int32(v))
    }

    fmt.Println(ia)
    
    before := int32(multiplierInit)
    
    for i := range ia {
        ia[i] = (ia[i] % (multiplierInit * before) )* before
        before = ia[i]
    }

    fmt.Println(ia)
    fmt.Println(hex.EncodeToString(ia))
}
