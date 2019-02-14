func bigger(numbers ...int) int {
    var bigger int

    for i, number := range numbers {
        if i == 0 || number > bigger {
            bigger = number
        }
    }

    return bigger
}
