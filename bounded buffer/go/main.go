// package main

// import (
// 	"fmt"
// 	"time"
// )

// func producer(ch chan int){

//     for i := 0; i < 10; i++ {
//         time.Sleep(100 * time.Millisecond)
//         fmt.Printf("[producer]: pushing %d\n", i)
//         ch <- i
//     }

// }

// func consumer(ch chan int){

//     time.Sleep(1 * time.Second)
//     for {
//         i := <- ch
//         fmt.Printf("[consumer]: %d\n", i)
//         time.Sleep(50 * time.Millisecond)
//     }

// }

// func main(){

//     // TODO: make a bounded buffer

//     buf := make(chan int, 5)

//     go consumer(buf)
//     go producer(buf)

//     select {}
// }

package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[producer]: pushing %d\n", i)
		ch <- i
	}
	close(ch) // Close the channel to signal that the producer is done
}

func consumer(ch chan int) {
	time.Sleep(1 * time.Second)
	for val := range ch { // Use range to read until the channel is closed
		fmt.Printf("[consumer]: %d\n", val)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("[consumer]: Finished consuming")
}

func main() {
	// Create a bounded buffer
	buf := make(chan int, 5)

	go consumer(buf)
	go producer(buf)

	time.Sleep(2 * time.Second) // Allow time for producer and consumer to finish
	fmt.Println("[main]: Exiting")
}
