package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philoNum, maxEat, eatPerPhilo = 5, 2, 3
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftChop, rightChop *ChopS
}

func (p Philo) eat(idx int, permission chan bool, wg *sync.WaitGroup) {

	for i := 0; i < eatPerPhilo; i++ {
		permission <- true

		p.leftChop.Lock()
		p.rightChop.Lock()

		fmt.Printf("Start to eat %d (%d EAT)\n", idx, i+1)
		time.Sleep(100 * time.Millisecond) // Eating time
		p.leftChop.Unlock()
		p.rightChop.Unlock()

		fmt.Printf("Finishing eating %d (%d EAT)\n", idx, i+1)

		// <-permission
	}
	wg.Done()

}

func host(permission chan bool) {
	for {
		// fmt.Println("----- Eat -----")
		<-permission
	}
}

func main() {
	var wg sync.WaitGroup
	permission := make(chan bool, maxEat)

	fmt.Printf("Start eating!!\n\n")

	chopSts := make([]*ChopS, 5)
	for i := 0; i < philoNum; i++ {
		chopSts[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < philoNum; i++ {
		philos[i] = &Philo{chopSts[i], chopSts[(i+1)%5]}
	}

	go host(permission)

	for i := 0; i < philoNum; i++ {
		wg.Add(1)
		go philos[i].eat(i+1, permission, &wg)
	}
	wg.Wait()

	fmt.Println("\nFinish eating!!")
}
