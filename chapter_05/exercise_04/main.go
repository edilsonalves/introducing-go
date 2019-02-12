package main

import "fmt"

func main() {
    var smaller int
    x := []int{
        48, 96, 86, 68,
        57, 82, 63, 70,
        37, 34, 83, 27,
        19, 97, 9, 17,
    }

    for i, number := range x {
        if i == 0 || number < smaller {
            smaller = number
        }
    }

    fmt.Println(smaller)
}
