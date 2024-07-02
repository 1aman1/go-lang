package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
	permission      chan int
	waiting         chan int
	done            chan int
	wg              *sync.WaitGroup
}

func (p Philo) Eat() {
	for i := 0; i < 3; i++ {
		p.waiting <- p.number - 1
		<-p.permission
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		fmt.Printf("finishing eating %d\n", p.number)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
		p.done <- 1
	}
	p.wg.Done()
}

func host(philos []*Philo, waiting chan int, done chan int) {
	firstEater := <-waiting
	secondEater := <-waiting

	philos[firstEater].permission <- 1
	philos[secondEater].permission <- 1

	for {
		<-done
		next := <-waiting
		philos[next].permission <- 1
	}
}

func main() {
	numPhilos := 5

	chops := make([]*ChopS, numPhilos)
	for i := 0; i < numPhilos; i++ {
		chops[i] = new(ChopS)
	}

	waiting := make(chan int, numPhilos)
	done := make(chan int, 2)
	var wg sync.WaitGroup
	philos := make([]*Philo, numPhilos)
	for i := 0; i < numPhilos; i++ {
		philos[i] = &Philo{chops[i], chops[(i+1)%numPhilos], i + 1, make(chan int), waiting, done, &wg}
	}

	wg.Add(numPhilos)
	for i := 0; i < numPhilos; i++ {
		go philos[i].Eat()
	}
	go host(philos, waiting, done)
	wg.Wait()
}
