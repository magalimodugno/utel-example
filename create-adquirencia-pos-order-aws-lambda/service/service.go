package service

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

type (
	Clients struct {
	}

	Service struct {
		clients *Clients
	}
)

func New(clients *Clients) *Service {
	return &Service{clients}
}

func (s *Service) CreateOrder(_ context.Context, _ *CreateOrderRequest) error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 0 {
		return errors.New("something went wrong")
	}
	return nil
}
