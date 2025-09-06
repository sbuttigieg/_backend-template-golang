package service

import "github.com/sbuttigieg/_backend-template-golang/foundation/config"

type Service struct {
	build    string
	deferred []func() error
}

func New(_ *config.Config, build string) (*Service, error) {
	s := &Service{
		build: build,
	}

	return s, nil
}
