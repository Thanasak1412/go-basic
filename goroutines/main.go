package main

import (
	"errors"
	"fmt"
	"sync"
)

func processItem(item int, ch chan<- string, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	if item%2 == 0 {
		errCh <- errors.New(fmt.Sprintf("Error processing item: %d", item))
	} else {
		ch <- fmt.Sprintf("Successfully processed item: %d", item)
	}
}

func main() {
	//	Collecting Results and Error from Goroutines Using Channels
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup
	resultCh := make(chan string, len(data))
	errorCh := make(chan error, len(data))

	for _, d := range data {
		wg.Add(1)
		go processItem(d, resultCh, errorCh, &wg)
	}

	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}

	for err := range errorCh {
		fmt.Println("Received error:", err)
	}
}
