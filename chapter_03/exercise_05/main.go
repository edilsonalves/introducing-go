package main

import "fmt"

func main() {
    fmt.Print("Temperature (°F): ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)
    celsius := (fahrenheit - 32) * 5 / 9
    fmt.Printf("%.2f °C\n", celsius)
}
