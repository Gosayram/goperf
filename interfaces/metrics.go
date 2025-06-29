package interfaces

import (
	"time"
)

// MetricsCollector defines the contract for collecting and analyzing performance metrics
// This replaces the current perf package structure
type MetricsCollector interface {
	// StartTest initializes a new test session
	StartTest(config *TestConfig) (*TestSession, error)

	// RecordRequest records the result of a single request
	RecordRequest(session *TestSession, result *RequestResult) error

	// GetStats returns current statistics for a test session
	GetStats(session *TestSession) (*Statistics, error)

	// FinishTest completes the test session and generates final report
	FinishTest(session *TestSession) (*TestReport, error)

	// Reset clears all collected metrics
	Reset(session *TestSession) error
}

// TestConfig represents configuration for a load test
// This replaces perf.Init struct
type TestConfig struct {
	Target      *Request      `json:"target"`
	Users       int           `json:"users"`
	Duration    time.Duration `json:"duration"`
	Iterations  int           `json:"iterations"`
	OutputLevel string        `json:"output_level"`
}

// TestSession represents an active test session
type TestSession struct {
	ID      string        `json:"id"`
	Config  *TestConfig   `json:"config"`
	Started time.Time     `json:"started"`
	Status  SessionStatus `json:"status"`
}

// SessionStatus represents the status of a user session during testing
type SessionStatus int

const (
	// SessionStatusCreated indicates that a session has been created but not started
	SessionStatusCreated SessionStatus = iota
	// SessionStatusRunning indicates that a session is currently active
	SessionStatusRunning
	// SessionStatusCompleted indicates that a session has finished successfully
	SessionStatusCompleted
	// SessionStatusFailed indicates that a session has failed or encountered errors
	SessionStatusFailed
)

// RequestResult represents the result of a single request
type RequestResult struct {
	URL          string        `json:"url"`
	StatusCode   int           `json:"status_code"`
	Duration     time.Duration `json:"duration"`
	Size         int           `json:"size"`
	Success      bool          `json:"success"`
	ErrorMessage string        `json:"error_message,omitempty"`
	Timestamp    time.Time     `json:"timestamp"`
}

// Statistics represents real-time test statistics
type Statistics struct {
	TotalRequests   int           `json:"total_requests"`
	SuccessRequests int           `json:"success_requests"`
	FailedRequests  int           `json:"failed_requests"`
	AvgLatency      time.Duration `json:"avg_latency"`
	MinLatency      time.Duration `json:"min_latency"`
	MaxLatency      time.Duration `json:"max_latency"`
	Throughput      float64       `json:"throughput"` // requests per second
	TotalBytes      int           `json:"total_bytes"`
}

// TestReport represents the final test report
// This replaces request.IterateReqRespAll and perf.Output structs
type TestReport struct {
	Session     *TestSession  `json:"session"`
	Stats       *Statistics   `json:"stats"`
	AssetStats  []*AssetStats `json:"asset_stats"`
	Started     time.Time     `json:"started"`
	Finished    time.Time     `json:"finished"`
	ElapsedTime time.Duration `json:"elapsed_time"`
}

// AssetStats represents statistics for a specific asset type
type AssetStats struct {
	URL         string        `json:"url"`
	Type        string        `json:"type"` // "js", "css", "img"
	Count       int           `json:"count"`
	AvgLatency  time.Duration `json:"avg_latency"`
	SuccessRate float64       `json:"success_rate"`
}

// String returns a string representation of session status
func (s SessionStatus) String() string {
	switch s {
	case SessionStatusCreated:
		return "created"
	case SessionStatusRunning:
		return "running"
	case SessionStatusCompleted:
		return "completed"
	case SessionStatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}
