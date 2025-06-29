/*
Package goperf is a highly concurrent website load tester with clean architecture.

This is the new main.go that replaces the monolithic 130+ line version.
It uses dependency injection, interfaces, and proper separation of concerns.

Usage examples:

Load testing:

	./goperf -url https://httpbin.org/get -sec 5 -users 5

Web server mode:

	./goperf -web -port 8080

Single fetch:

	./goperf -url https://httpbin.org/get -fetch

Fetch with all assets:

	./goperf -url https://httpbin.org/get -fetchall -printjson
*/
package main

import (
	"log"

	"github.com/Gosayram/goperf/core"
)

func main() {
	// Create application instance using clean architecture
	app, err := core.NewApp()
	if err != nil {
		log.Fatalf("Failed to create application: %v", err)
	}

	// Run the application
	if err := app.Run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}
