func min(numbers []float64) float64 {
    var smaller float64

    for i, number := range numbers {
        if i == 0 || number < smaller {
            smaller = number
        }
    }

    return smaller
}

func max(numbers []float64) float64 {
    var bigger float64

    for i, number := range numbers {
        if i == 0 || number > bigger {
            bigger = number
        }
    }

    return bigger
}
