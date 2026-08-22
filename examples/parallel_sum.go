package main

import (
	"fmt"
	"sync"
	"time"
)

// =============================================================================
// PARALLEL SUM - Learning Go Concurrency
// =============================================================================
//
// This file demonstrates how to use goroutines and channels to compute
// the sum of numbers in parallel, which is faster for large datasets.
//
// KEY CONCEPTS:
// - goroutine: A lightweight thread managed by Go runtime
// - channel:   A pipe for goroutines to communicate
// - sync.WaitGroup: Waits for a collection of goroutines to finish
// =============================================================================

// sequentialSum adds numbers one by one (slow for large arrays)
func sequentialSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

// parallelSumWithChannels splits work across goroutines using channels
//
// CHANNEL EXPLANATION:
// - make(chan int, numWorkers) creates a buffered channel
// - ch <- value    sends a value into the channel
// - value := <-ch  receives a value from the channel
func parallelSumWithChannels(numbers []int, numWorkers int) int {
	if len(numbers) == 0 {
		return 0
	}

	// Create a channel to receive partial sums from each worker
	// Buffered channel: can hold numWorkers values without blocking
	results := make(chan int, numWorkers)

	// Calculate chunk size for each worker
	chunkSize := (len(numbers) + numWorkers - 1) / numWorkers

	// Launch worker goroutines
	for i := 0; i < numWorkers; i++ {
		// Calculate start and end indices for this worker
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		if start >= len(numbers) {
			break
		}

		// go keyword launches a goroutine (runs concurrently)
		// This is like spawning a new thread, but much lighter
		go func(chunk []int) {
			sum := 0
			for _, n := range chunk {
				sum += n
			}
			results <- sum // Send partial sum to channel
		}(numbers[start:end])
	}

	// Collect results from all workers
	totalSum := 0
	workersLaunched := (len(numbers) + chunkSize - 1) / chunkSize
	for i := 0; i < workersLaunched; i++ {
		totalSum += <-results // Receive from channel (blocks until value available)
	}

	return totalSum
}

// parallelSumWithWaitGroup uses sync.WaitGroup for synchronization
//
// WAITGROUP EXPLANATION:
// - wg.Add(n)  increments the counter by n
// - wg.Done()  decrements the counter by 1
// - wg.Wait()  blocks until counter is 0
//
// MUTEX EXPLANATION:
// - mu.Lock()   acquires exclusive access
// - mu.Unlock() releases the lock
// - Prevents race conditions when multiple goroutines access shared data
func parallelSumWithWaitGroup(numbers []int, numWorkers int) int {
	if len(numbers) == 0 {
		return 0
	}

	var (
		totalSum int
		mu       sync.Mutex    // Protects totalSum from race conditions
		wg       sync.WaitGroup // Waits for all goroutines to complete
	)

	chunkSize := (len(numbers) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		if start >= len(numbers) {
			break
		}

		wg.Add(1) // Increment counter before launching goroutine

		go func(chunk []int) {
			defer wg.Done() // Decrement counter when goroutine completes

			// Calculate partial sum
			partialSum := 0
			for _, n := range chunk {
				partialSum += n
			}

			// Safely add to total (only one goroutine can do this at a time)
			mu.Lock()
			totalSum += partialSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait() // Block until all goroutines call Done()
	return totalSum
}

func main() {
	// Create a large slice of numbers to sum
	size := 100_000_000
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = i + 1
	}

	fmt.Printf("Summing %d numbers...\n\n", size)

	// Sequential sum
	start := time.Now()
	result1 := sequentialSum(numbers)
	duration1 := time.Since(start)
	fmt.Printf("Sequential Sum:    %d (took %v)\n", result1, duration1)

	// Parallel sum with channels (4 workers)
	start = time.Now()
	result2 := parallelSumWithChannels(numbers, 4)
	duration2 := time.Since(start)
	fmt.Printf("Parallel (Chan):   %d (took %v)\n", result2, duration2)

	// Parallel sum with WaitGroup (4 workers)
	start = time.Now()
	result3 := parallelSumWithWaitGroup(numbers, 4)
	duration3 := time.Since(start)
	fmt.Printf("Parallel (WG):     %d (took %v)\n", result3, duration3)

	// Parallel with more workers (8 workers)
	start = time.Now()
	result4 := parallelSumWithChannels(numbers, 8)
	duration4 := time.Since(start)
	fmt.Printf("Parallel (8 workers): %d (took %v)\n", result4, duration4)

	// Speedup calculation
	fmt.Printf("\n--- Performance ---\n")
	fmt.Printf("Speedup (4 workers): %.2fx faster\n", float64(duration1)/float64(duration2))
	fmt.Printf("Speedup (8 workers): %.2fx faster\n", float64(duration1)/float64(duration4))
}
