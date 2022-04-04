/*
Lena Horsley

12 July 2021

Write two goroutines which have a race condition when executed concurrently. Explain what the race
condition is and how it can occur.

*/
package main

import (
	"fmt"
	"time"
)

var cookieCount int

func main() {
	/*
		The race condition occurs because the goroutines (giveOneCookie and takeOneCookie) are
		simultaneously changing the value of the global variable cookieCount.
			giveOneCookie: increments cookieCount by 1. Prints "Here you go:" and the value of
				cookieCount
			takeOneCookie: decrements cookieCount by 1. Prints "Give me my cookie:" and the value of
				cookieCount

		To verify this, run the following command in the terminal:
		go run -race race.go

		Reference: https://golang.org/doc/articles/race_detector

		Output (you shoud see something similar):
			Here you go:  1
			==================
			WARNING: DATA RACE
			Read at 0x00000122dcb0 by goroutine 8:
			main.takeOneCookie()
				/Users/lenahorsley/Desktop/concurrency/week002/race.go:23 +0x3e

			Previous write at 0x00000122dcb0 by goroutine 7:
			main.giveOneCookie()
				/Users/lenahorsley/Desktop/concurrency/week002/race.go:18 +0x5a

			Goroutine 8 (running) created at:
			main.main()
				/Users/lenahorsley/Desktop/concurrency/week002/race.go:34 +0x5c

			Goroutine 7 (finished) created at:
			main.main()
				/Users/lenahorsley/Desktop/concurrency/week002/race.go:33 +0x44
			==================
			Give me my cookie:  2
			Found 1 data race(s)
			exit status 66

		NOTE: If you run the program with the command:
		go run race.go

		The race condition IS NOT obvious and the output is:
		Here you go:  1
		Give me my cookie:  2
	*/

	go giveOneCookie()
	go takeOneCookie()
	time.Sleep(1 * time.Second)
}

func giveOneCookie() {
	cookieCount++
	fmt.Println("Here you go: ", cookieCount)
}

func takeOneCookie() {
	cookieCount++
	fmt.Println("Give me my cookie: ", cookieCount)
}
