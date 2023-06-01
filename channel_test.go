package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Muhammad Shobir Abdussyakur"
		fmt.Println("Sending data to channel complete")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Prabowo SUbianto"

}

func TestChannnelAsParamater(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	defer close(channel)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Megawati Soekarnoputri"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Muhammad"
		channel <- "Shobir"
		channel <- "Abdussyakur"
	}()

	go func() {

		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Received data", data)
	}

	fmt.Println("Finish")
}
