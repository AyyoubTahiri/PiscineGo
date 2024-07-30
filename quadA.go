package main

import "fmt"

func QuadA(x, y int) {

    if x <= 0 || y <= 0 {
        return
    }

    fmt.Print("o")
    for i := 1; i < x-1; i++ {
        fmt.Print("-")
    }

    if x > 1 {
        fmt.Print("o")
    }
    fmt.Println()

    for j := 1; j < y-1; j++ {
        fmt.Print("|")
        for i := 1; i < x-1; i++ {
            fmt.Print(" ")
        }
        if x > 1 {
            fmt.Print("|")
        }
        fmt.Println()
    }

    if y > 1 {
        fmt.Print("o")
        for i := 1; i < x-1; i++ {
            fmt.Print("-")
        }
        if x > 1 {
            fmt.Print("o")
        }
        fmt.Println()
    }
}

func main() {
    QuadA(5,3)
QuadA(5,1)
QuadA(1,1)
QuadA(1,5)
}





package piscine

// Made by: mzakri,helbadao, at
import "github.com/01-edu/z01"

func QuadC(x, y int) {
    if x >= 0 && y >= 0 {
        for i := 1; i <= y; i++ {
            for j := 1; j <= x; j++ {
                if (i == 1) && (j == 1 || j == x) {
                    z01.PrintRune('A')
                } else if (i == y) && (j == 1 || j == x) {
                    z01.PrintRune('C')
                } else if i == 1 || i == y {
                    z01.PrintRune('B')
                } else if j == 1 || j == x {
                    z01.PrintRune('B')
                } else {
                    z01.PrintRune(' ')
                }
            }
            z01.PrintRune('\n')
        }
        z01.PrintRune('\n')
    } else {
        return
    }
}
package piscine

// Made by: mzakri,helbadao, at
import "github.com/01-edu/z01"

func QuadC(x, y int) {
    if x > 0 && y > 0 {
        for i := 1; i <= y; i++ {
            for j := 1; j <= x; j++ {
                if (i == 1) && (j == 1 || j == x) {
                    z01.PrintRune('A')
                } else if (i == y) && (j == 1 || j == x) {
                    z01.PrintRune('C')
                } else if i == 1 || i == y {
                    z01.PrintRune('B')
                } else if j == 1 || j == x {
                    z01.PrintRune('B')
                } else {
                    z01.PrintRune(' ')
                }
            }
            z01.PrintRune('\n')
        }
        z01.PrintRune('\n')
    } else {
        return
    }
}




func QuadA(x, y int) {
    if x > 0 && y > 0 {
        for i := 1; i <= y; i++ {
            for j := 1; j <= x; j++ {
                if (j == 1 || j == x) && (i == 1 || i == y) {
                    z01.PrintRune('o')
                    if j == x {
                        z01.PrintRune('\n')
                    }
                } else if (i == 1 && j > i) ||  (i == y && j < x) {
                    z01.PrintRune('-')
                } else {
                    if j == 1 || j == x {
                        z01.PrintRune('|')
                        if j == x {
                            z01.PrintRune('\n')
                        }
                    } else {
                        z01.PrintRune(' ')
                    }
                }
            }
        }
        z01.PrintRune('\n')
    } else {
        return
    }
}