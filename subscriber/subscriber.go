package subscriber

import (
	"context"
	"log"
)

type Subscriber struct {
	subId string
}

func NewSubscriber(subId string) *Subscriber {
	return &Subscriber{subId: subId}
}
func (s *Subscriber) React(ctx context.Context, body float32, err chan error) {
	log.Printf("Subscriber with id '%s' get information about territory's changes: %4.2f", s.subId, body)
}
func (s *Subscriber) Id() string {
	return s.subId
}
