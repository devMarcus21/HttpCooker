package invokers

import (
	"sync"
)

func InvokerAsync(callback func() int, calls int) map[int]int {
	var wg sync.WaitGroup
	wg.Add(calls)

	channel := make(chan int)

	for call := 0; call < calls; call++ {
		go func() {
			defer wg.Done()
			channel <- callback()
		}()
	}

	results := map[int]int{}

	// wg.Wait()
	for call := 0; call < calls; call++ {
		select {
		case result := <-channel:
			results[result] = results[result] + 1
		}
	}

	return results
}
