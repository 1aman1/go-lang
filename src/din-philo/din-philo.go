package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)

type ChopS struct{ sync.Mutex }

type Philo struct {
	no      int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) eat() {

	for i := 0; i < 3; i++ {
		host <- true

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("Starting to eat: ", p.no+1)
		fmt.Println("finishing eating: ", p.no+1)

		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-host
	}

	wg.Done()
}

func main() {

	chopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = &ChopS{}
	}

	philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philo{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	for i := range philosophers {
		wg.Add(1)
		go philosophers[i].eat()
	}

	wg.Wait()

}
