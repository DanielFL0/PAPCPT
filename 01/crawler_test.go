func TestFetchURLReturnsContent(t *testing.T) {
    // Test that FetchURL returns non-empty content or error for a valid URL
}

func TestConcurrentFetchLimits(t *testing.T) {
    // Test that the crawler respects concurrency limits (e.g., max goroutines)
}

func TestChannelFanInFanOut(t *testing.T) {
    // Verify fan-out sends URLs to workers, fan-in collects results correctly
}

func TestNoDeadlockOnChannelClose(t *testing.T) {
    // Ensure channels are closed properly, no goroutine leaks or deadlocks
}