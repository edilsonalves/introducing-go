package main

import "fmt"

func main() {
    fmt.Print("Length (ft): ")
    var feet float64
    fmt.Scanf("%f", &feet)
    meter := feet * 0.3048
    fmt.Printf("%.2f m\n", meter)
}
