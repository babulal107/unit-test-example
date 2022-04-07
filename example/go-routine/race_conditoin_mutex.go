package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var counter int

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d go routine : ", i)
		go IncrementCounter(str, &wg, &mutex)
	}
	wg.Wait()
}

func IncrementCounter(s string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()
	defer timeTrack(time.Now(), "IncrementCounter")
	for i := 0; i < 20; i++ {
		mutex.Lock()
		time.Sleep(5 * time.Millisecond)
		counter++
		fmt.Println(s, counter)
		mutex.Unlock()
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}