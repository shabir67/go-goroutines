package goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TesGettGomaxprocs(t *testing.T) {
	totalCPU := runtime.NumCPU()
	fmt.Println("CPU:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine:", totalGoroutine)

}
