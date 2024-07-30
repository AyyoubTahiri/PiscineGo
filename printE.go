package main

import "fmt"

func QuadE(x, y int) {

    if x <= 0 || y <= 0 {
        return
    }

    fmt.Print("A")
    for i := 1; i < x-1; i++ {
        fmt.Print("B")
    }

    if x > 1 {
        fmt.Print("C")
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
            fmt.Print("A")
        }
        fmt.Println()
    }
}

func main() {
QuadE(5,3)
QuadE(5,1)
QuadE(1,1)
QuadE(1,5)
}



