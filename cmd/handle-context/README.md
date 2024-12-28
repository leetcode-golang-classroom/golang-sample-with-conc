# handle-context

```golang
package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sourcegraph/conc/pool"
)

// validateFile simulates file validation. It returns an error for invalid files
func validateFile(fileName string) error {
	if strings.HasPrefix(fileName, "invalid") {
		return fmt.Errorf("validation failed for file: %s", fileName)
	}
	fmt.Printf("File %s validated successfully\n", fileName)
	return nil
}
func main() {
	// A list of files
	files := []string{
		"file1.csv",
		"invalid_file2.csv",
		"file3.csv",
		"file4.csv",
		"invalid_file5.csv",
		"file6.csv",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a new pool with error handling
	pool := pool.New().WithErrors().WithContext(ctx).WithMaxGoroutines(2)

	// Add validation tasks for each file
	for _, file := range files {
		pool.Go(func(poolCtx context.Context) error {
			twoSecondTimerCh := time.After(2 * time.Second)
			select {
			case <-poolCtx.Done():
				fmt.Printf("Task cancelled\n")
				return poolCtx.Err()
			case <-twoSecondTimerCh:
				err := validateFile(file)
				return err
			}
		})
	}

	// Wait for all tasks to finish
	if err := pool.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("All files validated successfully")
	}
}

```