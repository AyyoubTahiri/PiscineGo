package main

import "fmt"

func QuadC(x, y int) {

    if x <= 0 || y <= 0 {
        return
    }

    fmt.Print("A")
    for i := 1; i < x-1; i++ {
        fmt.Print("B")
    }

    if x > 1 {
        fmt.Print("A")
    }
    fmt.Println()

    for j := 1; j < y-1; j++ {
        fmt.Print("B")
        for i := 1; i < x-1; i++ {
            fmt.Print(" ")
        }
        if x > 1 {
            fmt.Print("B")
        }
        fmt.Println()
    }

    if y > 1 {
        fmt.Print("C")
        for i := 1; i < x-1; i++ {
            fmt.Print("B")
        }
        if x > 1 {
            fmt.Print("C")
        }
        fmt.Println()
    }
}

func main() {
QuadC(5,3)
QuadC(5,1)
QuadC(1,1)
QuadC(1,5)
}



