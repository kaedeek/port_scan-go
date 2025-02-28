// Copyright (c) 2025 kaedeek
// This software is released under the MIT License with additional restrictions.
// See the LICENSE file for details.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	startPort = 1
	endPort   = 1024
	timeout   = time.Second
)

// ScanResult represents the result of a port scan
type ScanResult struct {
	Port    int
	IsOpen  bool
	Service string
}

// ScanPort performs a port scan on a specific port
func ScanPort(ip string, port int) ScanResult {
	address := fmt.Sprintf("%s:%d", ip, port)
	result := ScanResult{Port: port}

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		result.IsOpen = false
		return result
	}
	defer conn.Close()

	result.IsOpen = true
	result.Service = getServiceName(port)
	return result
}

// getServiceName returns the common service name for well-known ports
func getServiceName(port int) string {
	services := map[int]string{
		20:  "FTP-DATA",
		21:  "FTP",
		22:  "SSH",
		23:  "Telnet",
		25:  "SMTP",
		53:  "DNS",
		80:  "HTTP",
		443: "HTTPS",
	}

	if service, ok := services[port]; ok {
		return service
	}
	return "Unknown"
}

// setupLogging configures the logging to write to both file and stdout
func setupLogging() (*os.File, error) {
	file, err := os.OpenFile("scan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening log file: %v", err)
	}

	log.SetOutput(file)
	return file, nil
}

func main() {
	// Parse command line arguments
	ip := flag.String("ip", "", "Target IP address to scan")
	flag.Parse()

	if *ip == "" {
		fmt.Println("Error: IP address is required")
		fmt.Println("Usage: go run main.go -ip <target_ip>")
		os.Exit(1)
	}

	// Setup logging
	logFile, err := setupLogging()
	if err != nil {
		fmt.Printf("Failed to setup logging: %v\n", err)
		os.Exit(1)
	}
	defer logFile.Close()

	log.Printf("Starting port scan on %s (ports %d-%d)\n", *ip, startPort, endPort)

	// Create channels for results and synchronization
	results := make(chan ScanResult, endPort-startPort+1)
	var wg sync.WaitGroup

	// Launch port scanning goroutines
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			result := ScanPort(*ip, p)
			if result.IsOpen {
				results <- result
			}
		}(port)
	}

	// Wait for all scans to complete in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Process and log results
	for result := range results {
		log.Printf("Port %d is open (%s)\n", result.Port, result.Service)
	}

	log.Println("Scan completed")
}
