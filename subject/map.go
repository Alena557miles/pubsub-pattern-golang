package subject

import (
	"context"
	"log"
	"pubsub-pattern-golang/publisher"
)

type Area struct {
	name string
	sq   float32
	*publisher.Publisher
}

func OccupiedTerritory(name string, sq float32, ctx context.Context) *Area {
	log.Printf("Country: %s, occupied territory: %4.2f", name, sq)
	sub := publisher.NewPublisher(ctx)
	return &Area{
		name,
		sq,
		sub,
	}
}

func (a *Area) FreeTerritory(sq float32) {
	if (a.sq >= 0) && (a.sq-sq > 0) {
		a.sq = a.sq - sq
		log.Printf("For now liberated territory is: %4.2f", a.sq)
		a.Publish(context.Background(), a.sq, nil)
	} else if (a.sq > 0) && (a.sq-sq < 0) {
		a.sq = sq - a.sq
		log.Printf("For now country %s are expanded its territory for: %4.2f", a.name, a.sq)
		a.Publish(context.Background(), a.sq, nil)
	} else {
		a.sq = a.sq + sq
		log.Printf("For now country %s are expanded its territory for: %4.2f", a.name, a.sq)
		a.Publish(context.Background(), a.sq, nil)
	}
}
