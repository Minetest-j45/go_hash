package main

import (
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/tncardoso/gocurses"
)

// a rune is actually just an int32 but we need to change the variable type in order to be able to pass it to functions etc.

func main() {
    if len(os.Args) < 2 {
        log.Fatal("Not enough arguments!")
    }

    input := strings.Join(os.Args[1:], " ") 
    ra := []rune(input)        // convert it to a rune array
    var ia []int32           // empty int32 array to store the result in
    for _, v := range ra {   // loop over every element of the rune array
        ia = append(ia, int32(v)) // convert the rune into int32 (just a type update, the data doesn't change, rune is just int32 with some methods) and append it to the result array
    }

    fmt.Println(ia) // prints [115 116 114 32 116 111 32 105 110 116 32 99 111 110 118] you can iterate this and perform the % operation
    
    for i := range ia {
        ia[i] %= 7
    }
    
    fmt.Println(ia)
    
    var before
    
    for i := range ia {
        if not before {
            before = 716//random number
        }
        ia[i] = ia[i]*before
        before = ia[i]
    }
}
