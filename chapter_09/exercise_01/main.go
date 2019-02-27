func average(numbers []float64) float64 {
    if len(numbers) == 0 {
        return 0
    }

    sum := float64(0)

    for _, number := range numbers {
        sum += number
    }

    return sum / float64(len(numbers))
}
