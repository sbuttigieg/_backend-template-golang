package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sbuttigieg/_backend-template-golang/api/cmd/backendtemplate/service"
	"github.com/sbuttigieg/_backend-template-golang/foundation/config"
)

// The following variable is used at build. We pass a revision hash via -ldflags, the revision hash is the output
// of command `git rev-parse HEAD` which can be found at github actions environment variable GITHUB_SHA
var build string

// @title Backend Template API
// @version 1.0.0
// @description This is a sample server for a backend template API.
// @host TBD
// @BasePath /api/v1
// @schemes http https
func main() {
	// Determine build version
	var buildflag bool

	flag.BoolVar(&buildflag, "build", false, "if true, print the build version and exit")
	flag.Parse()

	if buildflag {
		if build == "" {
			fmt.Println("Build version is empty")
			os.Exit(1)
		}

		fmt.Println("Build version: ", build)
		os.Exit(0)
	}

	// Create a new configuration
	cfg := config.New("BKT")
	err := cfg.Parse()
	if err != nil {
		log.Print("Failed to parse configuration:", err)
	}

	if err := run(cfg); err != nil {
		fmt.Println("Main function encountered an error:", err)
	}
}

func run(cfg *config.Config) error {
	log.Print("backendtemplate API started")

	// Handle termination signals
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// Initialise the service
	s, err := service.New(cfg, build)
	if err != nil {
		return fmt.Errorf("failed to initialise service: %w", err)
	}

	// Defer service destruction
	defer func() {
		if err := s.Destroy(cfg.App.DestroyTimeout); err != nil {
			log.Printf("Error during service destruction: %v", err)
		}
	}()

	<-ctx.Done()
	log.Print("backendtemplate received a termination signal")

	return nil
}
