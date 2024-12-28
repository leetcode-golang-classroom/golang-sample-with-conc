# aggregate-sample

```golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sourcegraph/conc/pool"
)

// simulateAPI simulates an API call to fetech product prices
func simulateAPI(apiID int) int {
	// Simulate varying response times
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random duration
	randomDuration := time.Duration(rng.Int63n(int64(2 * time.Second)))

	time.Sleep(randomDuration)

	// Return a mock price
	return apiID*10 + rand.Intn(10)
}

func main() {
	// Create a result pool
	pool := pool.NewWithResults[int]().WithMaxGoroutines(5)

	// Simulate fetching prices
	for apiID := 1; apiID <= 10; apiID++ {
		apiID := apiID // Capture variable for goroutine
		pool.Go(func() int {
			fmt.Printf("Fetching price from API %d...\n", apiID)
			return simulateAPI(apiID)
		})
	}

	// Wait for all API calls to finish
	prices := pool.Wait()
	// Display prices in the console
	fmt.Printf("Collected Prices: %v\n", prices)
}
```