package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestTimerAfter(t *testing.T) {
	chanel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-chanel
	fmt.Println(time)
}
