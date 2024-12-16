package common

import "sync"

const MaxCountInTelemetryQueue = 10

type RequestMetrics struct {
	PrevRequestId       string `json:"prev_request_id"`
	RequestId           string `json:"request_id"`
	PrevRequestDuration int    `json:"prev_request_duration"`
}

type TelemetryQueue struct {
	queue []RequestMetrics
	mutex sync.Mutex
}

func NewTelemetryQueue() *TelemetryQueue {
	return &TelemetryQueue{
		queue: make([]RequestMetrics, 0),
	}
}

func (tq *TelemetryQueue) Enqueue(metrics RequestMetrics) {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()

	if len(tq.queue) < MaxCountInTelemetryQueue {
		tq.queue = append(tq.queue, metrics)
	}
}

func (tq *TelemetryQueue) Dequeue() (RequestMetrics, bool) {
	tq.mutex.Lock()
	defer tq.mutex.Unlock()

	if len(tq.queue) == 0 {
		return RequestMetrics{}, false
	}

	metrics := tq.queue[0]
	tq.queue = tq.queue[1:]
	return metrics, true
}
