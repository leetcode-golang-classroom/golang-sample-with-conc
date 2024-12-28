# simple-task

```golang
package main

import (
	"fmt"
	"time"

	"github.com/sourcegraph/conc"
)

func task1() {
	// Simulate a task
	fmt.Println("Task 1 stared")
}

func main() {
	var group conc.WaitGroup

	// Add task to the group
	group.Go(task1)

	group.Go(func() {
		// Simulate another task
		fmt.Println("Task 2 started")
		time.Sleep(time.Second)
	})

	// Wait for all tasks to complete
	group.Wait()
	fmt.Println("All tasks completed successfully")
}
```