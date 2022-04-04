/*
Lena Horsley

16 July 2021

Implement the dining philosopher’s problem with the following constraints/modifications.

    - There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent
		pair of philosophers.
	- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in
		lecture).
    -In order to eat, a philosopher must get permission from a host which executes in its own
		goroutine.
    - The host allows no more than 2 philosophers to eat concurrently.
    - Each philosopher is numbered, 1 through 5.
    - When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to
		eat <number>” on a line by itself, where <number> is the number of the philosopher.

    - When a philosopher finishes eating (before it has released its locks) it prints “finishing
		eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	MAX_VALUE int = 5
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	philospherId   int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func main() {

	var myWaitGroup sync.WaitGroup

	// STEP 1: Create 5 chopsticks
	theChopsticks := make([]*Chopstick, MAX_VALUE)
	for i := 0; i < MAX_VALUE; i++ {
		theChopsticks[i] = new(Chopstick)
	}

	// STEP 2. Create the 5 philosophers
	thePhilosophers := make([]*Philosopher, MAX_VALUE)

	//STEP 3. Now EAT!
	for i := 0; i < MAX_VALUE; i++ {
		thePhilosophers[i] = &Philosopher{(i + 1), theChopsticks[i], theChopsticks[(i+1)%MAX_VALUE]}
		myWaitGroup.Add(1)
		go thePhilosophers[i].eat(&myWaitGroup)
	}

	myWaitGroup.Wait()

}

func (myPhilosopher Philosopher) eat(myWaitGroup *sync.WaitGroup) {
	// deadlock occurs without the sleeps
	defer myWaitGroup.Done()

	// the philosopher can eat three times max
	for maxTimesToEat := 0; maxTimesToEat < 3; maxTimesToEat++ {
		myPhilosopher.leftChopstick.Lock()
		myPhilosopher.rightChopstick.Lock()

		fmt.Printf("Philosopher #%d has grabbed the chopsticks and is now eating\n", myPhilosopher.philospherId)
		time.Sleep(time.Second)

		myPhilosopher.rightChopstick.Unlock()
		myPhilosopher.leftChopstick.Unlock()

		fmt.Printf("Philosopher #%d is no longer eating and has put the chopsticks down\n", myPhilosopher.philospherId)
		time.Sleep(time.Second)
	}
}
