package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		fmt.Println("Main function encountered an error:", err)
	}
}

func run() error {
	log.Print("backendtemplate started")

	// Handle termination signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Simulate some processing
	fmt.Println("Running the main function...")

	// Simulate an error
	// return fmt.Errorf("an error occurred in run function")

	<-ctx.Done()
	log.Print("backendtemplate received a termination signal")

	return nil
}
