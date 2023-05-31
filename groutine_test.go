package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutinne(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {

	fmt.Println("Display", number)

}

func TestMultipleGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go DisplayNumber(i)
	}

}
