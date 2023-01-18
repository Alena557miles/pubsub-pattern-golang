package publisher

import (
	"context"
	"log"
)

type Publisher struct {
	subscribers       map[string]Subscriber
	in                chan inBody
	addSubscriberChan chan Subscriber
	stop              chan chan struct{}
}

type inBody struct {
	ctx       context.Context
	body      float32
	returnErr chan error
}

func NewPublisher(ctx context.Context) *Publisher {
	p := &Publisher{
		subscribers:       make(map[string]Subscriber),
		in:                make(chan inBody),
		addSubscriberChan: make(chan Subscriber),
		stop:              make(chan chan struct{}),
	}
	go p.Start(ctx)
	return p
}
func (p *Publisher) AddSubscriber(subscriber Subscriber) {
	p.addSubscriberChan <- subscriber
}

func (p *Publisher) Publish(ctx context.Context, body float32, err chan error) {
	in := inBody{
		ctx,
		body,
		err,
	}

	p.in <- in
}

func (p *Publisher) Start(_ context.Context) {
	defer log.Println("Publisher finish listening for messages")
	log.Println("Publisher start listening for messages ")

	for {
		select {
		case body := <-p.in:
			for _, subs := range p.subscribers {
				subs.React(body.ctx, body.body, body.returnErr)
			}
		case s := <-p.addSubscriberChan:
			p.subscribers[s.Id()] = s
		case stop := <-p.stop:
			stop <- struct{}{}
			return
		}
	}
}

func (p *Publisher) Stop(ctx context.Context) error {
	log.Println(ctx, "Publisher stopping")
	stopAck := make(chan struct{})
	p.stop <- stopAck
	<-stopAck
	log.Println(ctx, "Publisher stopped")
	return nil
}

type Subscriber interface {
	React(ctx context.Context, body float32, err chan error)
	Id() string
}
