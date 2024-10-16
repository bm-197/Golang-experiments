package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// AggregateMetrics stores the aggregated results.
type AggregateMetrics struct {
	Count             int
	TotalResponseTime float64
}

// ProcessLogs reads a log file and sends aggregated metrics to the channel.
func ProcessLogs(filename string, wg *sync.WaitGroup, ch chan<- map[string]*AggregateMetrics) {
	defer wg.Done()

	metrics := make(map[string]*AggregateMetrics)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) < 3 {
			continue // Skip invalid lines
		}

		endpoint := parts[1]
		responseTime := parseResponseTime(parts[2]) // Implement this function to extract response time

		if _, exists := metrics[endpoint]; !exists {
			metrics[endpoint] = &AggregateMetrics{}
		}

		metrics[endpoint].Count++
		metrics[endpoint].TotalResponseTime += responseTime
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	ch <- metrics
}

func parseResponseTime(part string) float64 {
	var responseTime float64
	fmt.Sscanf(part, "%f", &responseTime)
	return responseTime
}

func main() {
	// Check if log files are provided as command-line arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <log_file1> <log_file2> ... <log_fileN>")
		return
	}

	// Collect log files from command-line arguments
	logFiles := os.Args[1:]

	var wg sync.WaitGroup
	ch := make(chan map[string]*AggregateMetrics)

	// Launch goroutines to process logs concurrently
	for _, logFile := range logFiles {
		wg.Add(1)
		go ProcessLogs(logFile, &wg, ch)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Aggregate results
	finalMetrics := make(map[string]*AggregateMetrics)
	for metrics := range ch {
		for endpoint, metric := range metrics {
			if _, exists := finalMetrics[endpoint]; !exists {
				finalMetrics[endpoint] = &AggregateMetrics{}
			}
			finalMetrics[endpoint].Count += metric.Count
			finalMetrics[endpoint].TotalResponseTime += metric.TotalResponseTime
		}
	}

	// Print summary of metrics
	fmt.Println("Summary of Log Metrics:")
	for endpoint, metric := range finalMetrics {
		avgResponseTime := metric.TotalResponseTime / float64(metric.Count)
		fmt.Printf("Endpoint: %s, Requests: %d, Avg. Response Time: %.2f ms\n",
			endpoint, metric.Count, avgResponseTime)
	}
}

