package implementations

import (
	"fmt"
	"time"

	"github.com/Gosayram/goperf/interfaces"
)

// MockMetricsCollector is a simple implementation for testing
type MockMetricsCollector struct {
	sessions map[string]*interfaces.TestSession
}

// NewMockMetricsCollector creates a new mock metrics collector
func NewMockMetricsCollector() *MockMetricsCollector {
	return &MockMetricsCollector{
		sessions: make(map[string]*interfaces.TestSession),
	}
}

// StartTest implements interfaces.MetricsCollector
func (m *MockMetricsCollector) StartTest(config *interfaces.TestConfig) (*interfaces.TestSession, error) {
	sessionID := fmt.Sprintf("test-%d", time.Now().Unix())

	session := &interfaces.TestSession{
		ID:      sessionID,
		Config:  config,
		Started: time.Now(),
		Status:  interfaces.SessionStatusRunning,
	}

	m.sessions[sessionID] = session

	return session, nil
}

// RecordRequest implements interfaces.MetricsCollector
func (m *MockMetricsCollector) RecordRequest(session *interfaces.TestSession, _ *interfaces.RequestResult) error {
	// In a real implementation, this would store the request result
	// For mock, we just update the session status
	if session.Status == interfaces.SessionStatusCreated {
		session.Status = interfaces.SessionStatusRunning
	}

	return nil
}

// GetStats implements interfaces.MetricsCollector
func (m *MockMetricsCollector) GetStats(_ *interfaces.TestSession) (*interfaces.Statistics, error) {
	// Mock statistics based on config
	return &interfaces.Statistics{
		TotalRequests:   MockTotalRequests,
		SuccessRequests: MockSuccessRequests,
		FailedRequests:  0,
		AvgLatency:      MockAvgLatency * time.Millisecond,
		MinLatency:      MockMinLatency * time.Millisecond,
		MaxLatency:      MockMaxLatency * time.Millisecond,
		Throughput:      MockThroughput, // requests per second
		TotalBytes:      MockTotalBytes,
	}, nil
}

// FinishTest implements interfaces.MetricsCollector
func (m *MockMetricsCollector) FinishTest(session *interfaces.TestSession) (*interfaces.TestReport, error) {
	session.Status = interfaces.SessionStatusCompleted

	stats, err := m.GetStats(session)
	if err != nil {
		return nil, err
	}

	finished := time.Now()

	report := &interfaces.TestReport{
		Session:     session,
		Stats:       stats,
		AssetStats:  []*interfaces.AssetStats{},
		Started:     session.Started,
		Finished:    finished,
		ElapsedTime: finished.Sub(session.Started),
	}

	return report, nil
}

// Reset implements interfaces.MetricsCollector
func (m *MockMetricsCollector) Reset(session *interfaces.TestSession) error {
	session.Status = interfaces.SessionStatusCreated
	return nil
}
