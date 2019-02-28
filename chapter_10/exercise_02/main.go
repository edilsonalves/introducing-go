func sleep(duration time.Duration) {
    <-time.After(duration)
}
