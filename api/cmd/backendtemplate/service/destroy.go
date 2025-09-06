package service

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Destroy destroys the service.
func (s *Service) Destroy(timeout time.Duration) error {
	fmt.Print("Destroying services")

	errChan := make(chan error)

	go s.runDeferred(errChan)

	select {
	case err := <-errChan:
		if err != nil {
			return fmt.Errorf("failed to destroy services: %w", err)
		}
	case <-time.After(time.Second * timeout):
		return fmt.Errorf("failed to destroy services: timeout of %s exceeded", timeout)
	}

	return nil
}

// runDeferred runs deferred functions
func (s *Service) runDeferred(errChan chan<- error) {
	defErrChan := make(chan error)

	for i := range s.deferred {
		go func(i int, defErrChan chan<- error) {
			defErrChan <- s.deferred[i]()
		}(i, defErrChan)
	}

	var messages []string

	for i := 0; i < len(s.deferred); i++ {
		if err := <-defErrChan; err != nil {
			messages = append(messages, err.Error())
		}
	}

	if len(messages) > 0 {
		errChan <- errors.New(strings.Join(messages, ", "))
		return
	}

	close(errChan)
}
