package main

import (
    "fmt"
)


func stringsam(s string) string {
    
    a := len(s) / 2

    return s[:a]
}

func main() {
    fmt.Println(stringsam("Bonjour Ayyoub")) 
    
}
