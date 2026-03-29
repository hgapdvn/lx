package lxtypes_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hgapdvn/lx/types"
)

// Example of basic async operation with FutureDo
func ExampleFutureDo() {
	future := lxtypes.FutureDo(func() (int, error) {
		// Simulate async work
		time.Sleep(100 * time.Millisecond)
		return 42, nil
	})

	result, err := future.Get(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
	// Output: 42
}

// Example of sequential operations with type transformation
func ExampleFutureThen() {
	// Start with user ID
	userIdFuture := lxtypes.FutureDo(func() (int, error) {
		return 123, nil
	})

	// Transform to user name
	userNameFuture := lxtypes.FutureThen(userIdFuture, func(userId int) (string, error) {
		return fmt.Sprintf("User_%d", userId), nil
	})

	// Transform to greeting
	greetingFuture := lxtypes.FutureThen(userNameFuture, func(name string) (string, error) {
		return fmt.Sprintf("Hello, %s!", name), nil
	})

	greeting, _ := greetingFuture.Get(context.Background())
	fmt.Println(greeting)
	// Output: Hello, User_123!
}

// Example of parallel execution with FutureAll
func ExampleFutureAll() {
	// Start 3 operations in parallel
	f1 := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 1, nil
	})
	f2 := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 2, nil
	})
	f3 := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(50 * time.Millisecond)
		return 3, nil
	})

	// Wait for all to complete
	allFuture := lxtypes.FutureAll(f1, f2, f3)
	results, err := allFuture.Get(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(results)
	// Output: [1 2 3]
}

// Example of joining different types with FutureJoin2
func ExampleFutureJoin2() {
	// Fetch user and config in parallel
	userFuture := lxtypes.FutureDo(func() (string, error) {
		return "Alice", nil
	})

	configFuture := lxtypes.FutureDo(func() (int, error) {
		return 100, nil
	})

	// Join them
	joined := lxtypes.FutureJoin2(userFuture, configFuture)
	result, err := joined.Get(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("User: %s, Config: %d\n", result.First, result.Second)
	// Output: User: Alice, Config: 100
}

// Example of context cancellation
func ExampleFuture_contextCancellation() {
	future := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(5 * time.Second)
		return "completed", nil
	})

	// Create context with 100ms timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := future.Get(ctx)
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Timeout!")
	}
	// Output: Timeout!
}

// Example of error handling
func ExampleFuture_errorHandling() {
	future := lxtypes.FutureDo(func() (int, error) {
		return 0, errors.New("something went wrong")
	})

	_, err := future.Get(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Output: Error: something went wrong
}

// Example of mixed sequential and parallel patterns
func ExampleFuture_mixedPattern() {
	// Step 1: Get base value
	step1 := lxtypes.FutureDo(func() (int, error) {
		return 10, nil
	})

	// Step 2: Use that value to start 3 parallel operations
	step2 := lxtypes.FutureThen(step1, func(n int) ([]int, error) {
		f1 := lxtypes.FutureDo(func() (int, error) { return n * 1, nil })
		f2 := lxtypes.FutureDo(func() (int, error) { return n * 2, nil })
		f3 := lxtypes.FutureDo(func() (int, error) { return n * 3, nil })

		allFuture := lxtypes.FutureAll(f1, f2, f3)
		return allFuture.Get(context.Background())
	})

	results, _ := step2.Get(context.Background())
	fmt.Println(results)
	// Output: [10 20 30]
}

// Example of combining same-type and different-type futures
func ExampleFuture_mixedTypes() {
	// Fetch 3 data sources of the same type in parallel
	data1 := lxtypes.FutureDo(func() (int, error) {
		return 100, nil
	})
	data2 := lxtypes.FutureDo(func() (int, error) {
		return 200, nil
	})
	data3 := lxtypes.FutureDo(func() (int, error) {
		return 300, nil
	})

	// Combine same-type data
	allData := lxtypes.FutureAll(data1, data2, data3) // Future[[]int]

	// Also fetch config (different type)
	config := lxtypes.FutureDo(func() (string, error) {
		return "production", nil
	})

	// Combine both: same-type collection + different type
	combined := lxtypes.FutureJoin2(allData, config) // Future[Pair[[]int, string]]

	// Transform the combined result
	result := lxtypes.FutureThen(combined, func(pair lxtypes.Pair[[]int, string]) (string, error) {
		sum := 0
		for _, v := range pair.First {
			sum += v
		}
		return fmt.Sprintf("Total: %d (env: %s)", sum, pair.Second), nil
	})

	output, _ := result.Get(context.Background())
	fmt.Println(output)
	// Output: Total: 600 (env: production)
}

// Example of fetching from multiple services (5+) and combining into response
func ExampleFuture_multipleServices() {
	// Real-world scenario: Build a dashboard by fetching from 5 different services

	// Start all service calls in parallel
	userFuture := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(20 * time.Millisecond)
		return "Alice", nil
	})

	ordersFuture := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(20 * time.Millisecond)
		return 3, nil // number of orders
	})

	paymentFuture := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(20 * time.Millisecond)
		return "paid", nil
	})

	inventoryFuture := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(20 * time.Millisecond)
		return 50, nil // items in stock
	})

	recommendationsFuture := lxtypes.FutureDo(func() (int, error) {
		time.Sleep(20 * time.Millisecond)
		return 5, nil // recommended products count
	})

	// Combine all 5 services using FutureJoin5
	allData := lxtypes.FutureJoin5(userFuture, ordersFuture, paymentFuture, inventoryFuture, recommendationsFuture)

	// Transform into final response
	response := lxtypes.FutureThen(allData,
		func(data lxtypes.Tuple5[string, int, string, int, int]) (string, error) {
			return fmt.Sprintf("Dashboard: %s | Orders: %d | Payment: %s | Stock: %d | Recommendations: %d",
				data.V1, data.V2, data.V3, data.V4, data.V5), nil
		})

	result, _ := response.Get(context.Background())
	fmt.Println(result)
	// Output: Dashboard: Alice | Orders: 3 | Payment: paid | Stock: 50 | Recommendations: 5
}

// Example of FutureJoin5 with different types
func ExampleFutureJoin5() {
	f1 := lxtypes.FutureDo(func() (int, error) {
		return 1, nil
	})
	f2 := lxtypes.FutureDo(func() (string, error) {
		return "two", nil
	})
	f3 := lxtypes.FutureDo(func() (bool, error) {
		return true, nil
	})
	f4 := lxtypes.FutureDo(func() (float64, error) {
		return 4.0, nil
	})
	f5 := lxtypes.FutureDo(func() ([]int, error) {
		return []int{5, 6}, nil
	})

	// Combine all 5 futures
	combined := lxtypes.FutureJoin5(f1, f2, f3, f4, f5)
	result, _ := combined.Get(context.Background())

	fmt.Printf("V1=%d, V2=%s, V3=%t, V4=%.1f, V5=%v\n",
		result.V1, result.V2, result.V3, result.V4, result.V5)
	// Output: V1=1, V2=two, V3=true, V4=4.0, V5=[5 6]
}

// Example of FutureJoin6 with six different types
func ExampleFutureJoin6() {
	f1 := lxtypes.FutureDo(func() (int, error) { return 1, nil })
	f2 := lxtypes.FutureDo(func() (string, error) { return "two", nil })
	f3 := lxtypes.FutureDo(func() (bool, error) { return true, nil })
	f4 := lxtypes.FutureDo(func() (float64, error) { return 4.0, nil })
	f5 := lxtypes.FutureDo(func() ([]int, error) { return []int{5}, nil })
	f6 := lxtypes.FutureDo(func() (rune, error) { return 'a', nil })

	combined := lxtypes.FutureJoin6(f1, f2, f3, f4, f5, f6)
	result, _ := combined.Get(context.Background())

	fmt.Printf("Values: %d, %s, %t, %.1f, %v, %c\n",
		result.V1, result.V2, result.V3, result.V4, result.V5, result.V6)
	// Output: Values: 1, two, true, 4.0, [5], a
}

// Example of FutureJoin7 with seven different types
func ExampleFutureJoin7() {
	f1 := lxtypes.FutureDo(func() (int, error) { return 1, nil })
	f2 := lxtypes.FutureDo(func() (string, error) { return "two", nil })
	f3 := lxtypes.FutureDo(func() (bool, error) { return true, nil })
	f4 := lxtypes.FutureDo(func() (float64, error) { return 4.0, nil })
	f5 := lxtypes.FutureDo(func() ([]int, error) { return []int{5}, nil })
	f6 := lxtypes.FutureDo(func() (rune, error) { return 'a', nil })
	f7 := lxtypes.FutureDo(func() (byte, error) { return 7, nil })

	combined := lxtypes.FutureJoin7(f1, f2, f3, f4, f5, f6, f7)
	result, _ := combined.Get(context.Background())

	fmt.Printf("7 values: %d, %s, %t, %.1f, %v, %c, %d\n",
		result.V1, result.V2, result.V3, result.V4, result.V5, result.V6, result.V7)
	// Output: 7 values: 1, two, true, 4.0, [5], a, 7
}

// Example of FutureJoin8 with eight different types
func ExampleFutureJoin8() {
	f1 := lxtypes.FutureDo(func() (int, error) { return 1, nil })
	f2 := lxtypes.FutureDo(func() (string, error) { return "two", nil })
	f3 := lxtypes.FutureDo(func() (bool, error) { return true, nil })
	f4 := lxtypes.FutureDo(func() (float64, error) { return 4.0, nil })
	f5 := lxtypes.FutureDo(func() ([]int, error) { return []int{5}, nil })
	f6 := lxtypes.FutureDo(func() (rune, error) { return 'a', nil })
	f7 := lxtypes.FutureDo(func() (byte, error) { return 7, nil })
	f8 := lxtypes.FutureDo(func() (uint, error) { return 8, nil })

	combined := lxtypes.FutureJoin8(f1, f2, f3, f4, f5, f6, f7, f8)
	result, _ := combined.Get(context.Background())

	fmt.Printf("8 values: %d, %s, %t, %.1f, %v, %c, %d, %d\n",
		result.V1, result.V2, result.V3, result.V4, result.V5, result.V6, result.V7, result.V8)
	// Output: 8 values: 1, two, true, 4.0, [5], a, 7, 8
}

// Example showing parallel execution performance with FutureJoin5
func ExampleFutureJoin5_parallel() {
	// Each service takes 100ms
	// Sequential would take 500ms, parallel takes ~100ms

	f1 := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "service1", nil
	})
	f2 := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "service2", nil
	})
	f3 := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "service3", nil
	})
	f4 := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "service4", nil
	})
	f5 := lxtypes.FutureDo(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "service5", nil
	})

	start := time.Now()
	combined := lxtypes.FutureJoin5(f1, f2, f3, f4, f5)
	result, _ := combined.Get(context.Background())
	duration := time.Since(start)

	fmt.Printf("Services: %s, %s, %s, %s, %s\n",
		result.V1, result.V2, result.V3, result.V4, result.V5)
	fmt.Printf("Completed in: ~%dms (parallel)\n", duration.Milliseconds()/100*100)
	// Output:
	// Services: service1, service2, service3, service4, service5
	// Completed in: ~100ms (parallel)
}
