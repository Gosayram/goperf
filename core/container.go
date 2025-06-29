package core

import (
	"context"
	"fmt"
	"time"

	"github.com/Gosayram/goperf/interfaces"
)

// Container manages all application dependencies
// This replaces scattered direct instantiations throughout the codebase
type Container struct {
	httpClient  interfaces.HTTPClient
	assetParser interfaces.AssetParser
	metrics     interfaces.MetricsCollector
	formatter   interfaces.OutputFormatter
	config      *Config
}

// NewContainer creates a new dependency injection container
func NewContainer(config *Config) *Container {
	container := &Container{
		config: config,
	}

	// Wire up dependencies
	container.initServices()

	return container
}

// initServices initializes all services with their dependencies
func (c *Container) initServices() {
	// Initialize with mock implementations for demonstration
	// TODO: These will be replaced with actual implementations later

	// Initialize HTTP client
	c.httpClient = newMockHTTPClient()

	// Initialize asset parser
	c.assetParser = newMockAssetParser()

	// Initialize metrics collector
	c.metrics = newMockMetricsCollector()

	// Initialize output formatter
	c.formatter = newMockOutputFormatter()
}

// Mock constructors using implementations package
func newMockHTTPClient() interfaces.HTTPClient {
	return &mockHTTPClient{}
}

func newMockAssetParser() interfaces.AssetParser {
	return &mockAssetParser{}
}

func newMockMetricsCollector() interfaces.MetricsCollector {
	return &mockMetricsCollector{}
}

func newMockOutputFormatter() interfaces.OutputFormatter {
	return &mockOutputFormatter{}
}

// Simple inline mocks for core functionality
type mockHTTPClient struct{}

func (c *mockHTTPClient) Fetch(_ context.Context, req *interfaces.Request) (*interfaces.Response, error) {
	return &interfaces.Response{
		URL: req.URL, StatusCode: MockHTTPStatusOK, Body: "Mock response",
		Size: MockResponseSize, Duration: MockLatency * time.Millisecond,
	}, nil
}

func (c *mockHTTPClient) FetchBatch(ctx context.Context, req *interfaces.Request) (*interfaces.BatchResponse, error) {
	baseResp, _ := c.Fetch(ctx, req)
	return &interfaces.BatchResponse{
		BaseResponse: baseResp,
		Assets:       []*interfaces.Response{},
		TotalTime:    TotalTimeInMilliseconds * time.Millisecond,
		TotalSize:    MockResponseSize,
	}, nil
}

func (c *mockHTTPClient) SetTimeout(_ time.Duration) {}
func (c *mockHTTPClient) SetUserAgent(_ string)      {}
func (c *mockHTTPClient) SetMaxConnections(_ int)    {}

type mockAssetParser struct{}

func (p *mockAssetParser) ParseAssets(_ string) (*interfaces.Assets, error) {
	return &interfaces.Assets{Total: 0}, nil
}
func (p *mockAssetParser) ParseJS(_ string) ([]string, error)          { return []string{}, nil }
func (p *mockAssetParser) ParseCSS(_ string) ([]string, error)         { return []string{}, nil }
func (p *mockAssetParser) ParseImages(_ string) ([]string, error)      { return []string{}, nil }
func (p *mockAssetParser) SetParsingMethod(_ interfaces.ParsingMethod) {}

type mockMetricsCollector struct{}

func (m *mockMetricsCollector) StartTest(config *interfaces.TestConfig) (*interfaces.TestSession, error) {
	return &interfaces.TestSession{
		ID:      "mock-session",
		Config:  config,
		Started: time.Now(),
		Status:  interfaces.SessionStatusRunning,
	}, nil
}

func (m *mockMetricsCollector) RecordRequest(_ *interfaces.TestSession, _ *interfaces.RequestResult) error {
	return nil
}

func (m *mockMetricsCollector) GetStats(_ *interfaces.TestSession) (*interfaces.Statistics, error) {
	return &interfaces.Statistics{
		TotalRequests:   MockTotalRequests,
		SuccessRequests: MockSuccessRequests,
		AvgLatency:      MockAvgLatency * time.Millisecond,
		Throughput:      MockThroughput,
		TotalBytes:      MockTotalBytes,
	}, nil
}

func (m *mockMetricsCollector) FinishTest(session *interfaces.TestSession) (*interfaces.TestReport, error) {
	stats, _ := m.GetStats(session)
	return &interfaces.TestReport{
		Session:     session,
		Stats:       stats,
		Started:     session.Started,
		Finished:    time.Now(),
		ElapsedTime: time.Since(session.Started),
	}, nil
}

func (m *mockMetricsCollector) Reset(_ *interfaces.TestSession) error { return nil }

type mockOutputFormatter struct{}

func (f *mockOutputFormatter) FormatJSON(_ interface{}) ([]byte, error) {
	return []byte(`{"mock": "json"}`), nil
}

func (f *mockOutputFormatter) FormatText(data interface{}) (string, error) {
	if report, ok := data.(*interfaces.TestReport); ok {
		return fmt.Sprintf(`
🚀 GoPerf Load Test Results
═══════════════════════════

📊 Test Summary:
  URL: %s
  Users: %d
  Duration: %v

📈 Performance Metrics:
  Total Requests: %d
  Success Rate: 100%%
  Avg Latency: %v
  Throughput: %.1f req/sec

✅ Test completed successfully!
`,
			report.Session.Config.Target.URL,
			report.Session.Config.Users,
			report.ElapsedTime,
			report.Stats.TotalRequests,
			report.Stats.AvgLatency,
			report.Stats.Throughput), nil
	}
	return "Mock text output", nil
}

func (f *mockOutputFormatter) FormatCSV(_ interface{}) ([]byte, error) {
	return []byte("Mock,CSV,Data"), nil
}
func (f *mockOutputFormatter) FormatHTML(_ interface{}) ([]byte, error) {
	return []byte("<html><body>Mock HTML</body></html>"), nil
}
func (f *mockOutputFormatter) SetIndentation(_ string) {}
func (f *mockOutputFormatter) SetColors(_ bool)        {}

// HTTPClient returns the configured HTTP client
func (c *Container) HTTPClient() interfaces.HTTPClient {
	return c.httpClient
}

// AssetParser returns the configured asset parser
func (c *Container) AssetParser() interfaces.AssetParser {
	return c.assetParser
}

// MetricsCollector returns the configured metrics collector
func (c *Container) MetricsCollector() interfaces.MetricsCollector {
	return c.metrics
}

// OutputFormatter returns the configured output formatter
func (c *Container) OutputFormatter() interfaces.OutputFormatter {
	return c.formatter
}

// Config returns the application configuration
func (c *Container) Config() *Config {
	return c.config
}

// Shutdown gracefully shuts down all services
func (c *Container) Shutdown() error {
	// TODO: Implement graceful shutdown
	// - Close HTTP connections
	// - Flush metrics
	// - Save any pending data
	return nil
}
