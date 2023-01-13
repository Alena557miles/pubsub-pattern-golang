package main

import (
	"context"
	"pubsub-pattern-golang/subject"
	"pubsub-pattern-golang/subscriber"
)

func main() {
	ctx := context.Background()
	a := subject.OccupiedTerritory("Ukraine", 120.7, ctx)

	s1 := subscriber.NewSubscriber("Ivan")
	s2 := subscriber.NewSubscriber("Svetlana")
	a.AddSubscriber(s1)
	a.AddSubscriber(s2)

	a.FreeTerritory(68.4)

}
