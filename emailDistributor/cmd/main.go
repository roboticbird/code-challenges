package main

import (
	"bufio"
	"code-challenges/emailDistributor/internal/emaildistributor"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal(fmt.Sprintf("Expected: %s <input file> <number of workers>", os.Args[0]))
	}

	// number of workers
	workers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if workers <= 0 {
		log.Fatal(fmt.Sprintf("Number of workers needs to be greater than 0, Received %d", workers))
	}

	// file path
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		// if we were actually sending emails we might want to validate them at some point
		lines = append(lines, scanner.Text())
	}

	// execute email distributor
	start := time.Now()
	succeeded, failed := emaildistributor.DistributeEmails(lines, workers)
	elapsed := time.Since(start)

	fmt.Printf("-----Finished sending emails.-----\n")
	fmt.Printf("Successfully sent: %d\n", len(succeeded))
	fmt.Printf("Failed to send: %d\n", len(failed))
	fmt.Printf("Number of workers: %d\n", workers)
	fmt.Printf("Execution time: %s\n", elapsed)
}
