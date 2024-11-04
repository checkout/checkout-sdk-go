package common

import (
	"strconv"
	"sync"
	"testing"
)

const maxCountInTelemetryQueue = 10

func TestTelemetryQueue_Concurrency(t *testing.T) {
	tq := NewTelemetryQueue()

	var wg sync.WaitGroup
	numWorkers := 20 
	numEnqueues := 5 

	enqueueWorker := func(id int) {
		defer wg.Done()
		for i := 0; i < numEnqueues; i++ {
			metrics := RequestMetrics{
				PrevRequestId:       "",
				RequestId:           strconv.Itoa(id) + "_" + strconv.Itoa(i),
				PrevRequestDuration: i,
			}
			tq.Enqueue(metrics)
		}
	}

	dequeueWorker := func() {
		defer wg.Done()
		for i := 0; i < numEnqueues; i++ {
			_, _ = tq.Dequeue()
		}
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go enqueueWorker(i)
	}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go dequeueWorker()
	}

	// Wait for all workers to finish
	wg.Wait()

	if len(tq.queue) > maxCountInTelemetryQueue {
		t.Errorf("Expected queue to have max %d items, but found %d items", maxCountInTelemetryQueue, len(tq.queue))
	}
}
