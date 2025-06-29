// Package core provides the main application logic and dependency injection container
// for the GoPerf load testing tool. It implements clean architecture principles
// with proper separation of concerns and interface-based design.
package core

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Gosayram/goperf/interfaces"
)

// App represents the main application
// This replaces the monolithic main function
type App struct {
	container *Container
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewApp creates a new application instance
func NewApp() (*App, error) {
	// Load configuration
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	// Create DI container
	container := NewContainer(config)

	// Create context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	return &App{
		container: container,
		ctx:       ctx,
		cancel:    cancel,
	}, nil
}

// Run starts the application based on configuration
func (a *App) Run() error {
	// Setup graceful shutdown
	a.setupShutdown()

	config := a.container.Config()

	// Determine run mode based on configuration
	if config.Web.Enabled {
		return a.runWebServer()
	}

	// Check for special commands
	// TODO: Add support for -fetch, -fetchall flags

	// Default to load testing mode
	return a.runLoadTest()
}

// runWebServer starts the HTTP web server
func (a *App) runWebServer() error {
	config := a.container.Config()

	fmt.Printf("Starting web server on port %d...\n", config.Web.Port)

	// TODO: Implement web server using the container services
	// This will replace the current web server in main.go

	<-a.ctx.Done()
	return nil
}

// runLoadTest performs a load test
func (a *App) runLoadTest() error {
	config := a.container.Config()

	// Create test configuration
	testConfig := &interfaces.TestConfig{
		Target: &interfaces.Request{
			URL:           config.Test.DefaultURL,
			Method:        "GET",
			Headers:       make(map[string]string),
			UserAgent:     config.HTTP.UserAgent,
			Timeout:       config.HTTP.Timeout,
			ReturnContent: false,
		},
		Users:       config.Test.DefaultUsers,
		Duration:    config.Test.DefaultDuration,
		Iterations:  config.Test.Iterations,
		OutputLevel: config.Log.Level,
	}

	// Get services from container
	metrics := a.container.MetricsCollector()
	formatter := a.container.OutputFormatter()

	// Start test session
	session, err := metrics.StartTest(testConfig)
	if err != nil {
		return fmt.Errorf("failed to start test: %w", err)
	}

	fmt.Printf("Starting load test: %d users for %v\n",
		testConfig.Users, testConfig.Duration)

	// TODO: Implement actual load testing logic using container services
	// This will replace the current perf package logic

	// For now, simulate a test
	time.Sleep(testConfig.Duration)

	// Finish test and get report
	report, err := metrics.FinishTest(session)
	if err != nil {
		return fmt.Errorf("failed to finish test: %w", err)
	}

	// Format and output results
	output, err := formatter.FormatText(report)
	if err != nil {
		return fmt.Errorf("failed to format output: %w", err)
	}

	fmt.Println(output)

	return nil
}

// setupShutdown configures graceful shutdown handling
func (a *App) setupShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received shutdown signal, gracefully shutting down...")

		// Cancel context to signal all components to stop
		a.cancel()

		// Shutdown container services
		if err := a.container.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}

		os.Exit(0)
	}()
}

// Shutdown gracefully shuts down the application
func (a *App) Shutdown() error {
	a.cancel()
	return a.container.Shutdown()
}
