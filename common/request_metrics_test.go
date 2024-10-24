package common

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

// Test for concurrent Enqueue and Dequeue operations with a MaxCount of 10
func TestTelemetryQueue_Concurrency(t *testing.T) {
	tq := NewTelemetryQueue()

	var wg sync.WaitGroup
	numWorkers := 20 // Reduced number of workers to match smaller queue size
	numEnqueues := 5 // Adjusted number of enqueues to work with MaxCount of 10

	// Function to enqueue metrics
	enqueueWorker := func(id int) {
		defer wg.Done()
		for i := 0; i < numEnqueues; i++ {
			metrics := RequestMetrics{
				PrevRequestId:       "",
				RequestId:           strconv.Itoa(id) + "_" + strconv.Itoa(i),
				PrevRequestDuration: i,
			}
			tq.Enqueue(metrics)
			time.Sleep(1 * time.Millisecond) // Slight delay to simulate real-world scenario
		}
	}

	// Function to dequeue metrics
	dequeueWorker := func() {
		defer wg.Done()
		for i := 0; i < numEnqueues; i++ {
			_, _ = tq.Dequeue()
			time.Sleep(1 * time.Millisecond) // Slight delay to simulate real-world scenario
		}
	}

	// Start multiple workers to enqueue
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go enqueueWorker(i)
	}

	// Start multiple workers to dequeue
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go dequeueWorker()
	}

	// Wait for all workers to finish
	wg.Wait()

	// Final check: Queue should be empty or within max limit
	if len(tq.queue) > MaxCountInTelemetryQueue {
		t.Errorf("Expected queue to have max %d items, but found %d items", MaxCountInTelemetryQueue, len(tq.queue))
	}
}
