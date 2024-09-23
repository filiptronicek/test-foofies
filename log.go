package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting extended log streaming tests...")

	// // Test 1: Rapid logging
	// fmt.Println("\nTest 1: Rapid logging")
	// rapidLogging(10000, 10*time.Millisecond)

	// Test 2: Updating same line
	fmt.Println("\nTest 2: Updating same line")
	updateSameLine(100)

	// // Test 3: Long-running logs
	// fmt.Println("\nTest 3: Long-running logs")
	// longRunningLogs(5 * time.Second)

	// // Test 4: Variable log levels
	// fmt.Println("\nTest 4: Variable log levels")
	// variableLogLevels(100)

	// // Test 5: Concurrent logging
	// fmt.Println("\nTest 5: Concurrent logging")
	// concurrentLogging(5, 100)

	// // Test 6: Large log entries
	// fmt.Println("\nTest 6: Large log entries")
	// largeLogEntries(10)

	// // Test 7: Intermittent logging
	// fmt.Println("\nTest 7: Intermittent logging")
	// intermittentLogging(15 * time.Second)

	fmt.Println("\nExtended log streaming tests completed.")
}

func rapidLogging(count int, duration time.Duration) {
	start := time.Now()
	for i := 0; i < count; i++ {
		fmt.Printf("Rapid log entry %d: %s\n", i, randomString(20))
		time.Sleep(duration / time.Duration(count))
	}
	elapsed := time.Since(start)
	fmt.Printf("Logged %d entries in %v\n", count, elapsed)
}

func updateSameLine(count int) {
	for i := 0; i < count; i++ {
		length := rand.Intn(50) + 10 // Random length between 10 and 59
		fmt.Printf("\rUpdating line %d: %s", i, randomString(length))
		os.Stdout.Sync()
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println() // Add a newline at the end
}

func longRunningLogs(duration time.Duration) {
	start := time.Now()
	count := 0
	for time.Since(start) < duration {
		fmt.Printf("Long-running log %d: %s\n", count, randomString(30))
		count++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Logged %d entries over %v\n", count, duration)
}

func variableLogLevels(count int) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	for i := 0; i < count; i++ {
		level := levels[rand.Intn(len(levels))]
		fmt.Printf("[%s] Log entry %d: %s\n", level, i, randomString(25))
		time.Sleep(50 * time.Millisecond)
	}
}

func concurrentLogging(sources, entriesPerSource int) {
	var wg sync.WaitGroup
	for i := 0; i < sources; i++ {
		wg.Add(1)
		go func(sourceID int) {
			defer wg.Done()
			for j := 0; j < entriesPerSource; j++ {
				fmt.Printf("Source %d - Log %d: %s\n", sourceID, j, randomString(20))
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
}

func largeLogEntries(count int) {
	for i := 0; i < count; i++ {
		size := rand.Intn(9000) + 1000 // Random size between 1KB and 10KB
		fmt.Printf("Large log entry %d (%d bytes): %s\n", i, size, randomString(size))
		time.Sleep(500 * time.Millisecond)
	}
}

func intermittentLogging(duration time.Duration) {
	start := time.Now()
	for time.Since(start) < duration {
		burstCount := rand.Intn(20) + 1
		for i := 0; i < burstCount; i++ {
			fmt.Printf("Intermittent log burst %d - Entry %d: %s\n", i, i, randomString(30))
		}
		quietPeriod := time.Duration(rand.Intn(5000)+1000) * time.Millisecond
		time.Sleep(quietPeriod)
	}
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}