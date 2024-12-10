package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	piChan := make(chan float64)

	for i := range 3 {
		wg.Add(1)
		go calcPi(i, &wg, piChan)
	}

	wg.Wait()

	//close(piChan)

}

func calcPi(s int, wg *sync.WaitGroup, partPi chan float64) {
	defer wg.Done()
	log.Printf("%d test\n", s)
	log.Printf("%d test2\n", s)
	log.Printf("%d test3\n", s)
	var pi float64 = 0.0
	pi = pi + 1
	pi = pi + 1
	pi = pi + 1
	if s == 1 {
		//partPi <- pi
	}
}

func calcPiOrig() {
	const numerator = 4.0
	var pi float64 = 0
	isNegative := false
	for i := 1; i <= 11000000000; i = i + 2 {
		if isNegative {
			pi = pi - numerator/float64(i)
			isNegative = false
		} else {
			pi = pi + numerator/float64(i)
			isNegative = true
		}
	}
	log.Println(pi)
	fmt.Printf("Result pi=, %d!\n", pi)
}
