package main

import "fmt"

func QuadB(x, y int) {

    if x <= 0 || y <= 0 {
        return
    }

    fmt.Print("/")
    for i := 1; i < x-1; i++ {
        fmt.Print("*")
    }

    if x > 1 {
        fmt.Print("\\")
    }
    fmt.Println()

    for j := 1; j < y-1; j++ {
        fmt.Print("*")
        for i := 1; i < x-1; i++ {
            fmt.Print(" ")
        }
        if x > 1 {
            fmt.Print("*")
        }
        fmt.Println()
    }

    if y > 1 {
        fmt.Print("\\")
        for i := 1; i < x-1; i++ {
            fmt.Print("*")
        }
        if x > 1 {
            fmt.Print("/")
        }
        fmt.Println()
    }
}

func main() {
QuadB(5,3)
QuadB(5,1)
QuadB(1,1)
QuadB(1,5)
}



