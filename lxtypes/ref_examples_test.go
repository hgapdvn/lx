package lxtypes_test

import (
	"fmt"
	"sync"

	"github.com/hgapdvn/lx/lxtypes"
)

// ExampleNewRef demonstrates creating a Ref with an initial value.
func ExampleNewRef() {
	ref := lxtypes.NewRef(42)
	fmt.Println(ref.Get())

	// Output:
	// 42
}

// ExampleRef_Set demonstrates replacing the current value.
func ExampleRef_Set() {
	ref := lxtypes.NewRef("hello")

	ref.Set("world")
	fmt.Println(ref.Get())

	// Output:
	// world
}

// ExampleRef_Update demonstrates atomically transforming the current value.
func ExampleRef_Update() {
	counter := lxtypes.NewRef(0)

	counter.Update(func(v int) int { return v + 1 })
	counter.Update(func(v int) int { return v + 1 })
	counter.Update(func(v int) int { return v + 1 })

	fmt.Println(counter.Get())

	// Output:
	// 3
}

// ExampleRef_concurrent demonstrates safe concurrent use across goroutines.
func ExampleRef_concurrent() {
	counter := lxtypes.NewRef(0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Update(func(v int) int { return v + 1 })
		}()
	}
	wg.Wait()

	fmt.Println(counter.Get())

	// Output:
	// 10
}
