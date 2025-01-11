// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"sync"
)


func server(requestChanel chan string, responseChannel chan int)  {
    var i =0
    for {
    select {
    case req:= <-requestChanel:
        switch req {
        case "increment":
            i++
        case  "decrement":
            i--
        case "get":
            responseChannel <- i

        case "done":
            close(responseChannel)
            return
        
        default:
            Println("Unkown request", req)

        

            
        }
    }
}
    
}

func incrementing(ch chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {
        ch <- "increment"
    }
}

func decrementing(ch chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    //TODO: decrement i 1000000 times
    for j := 0; j < 999998; j++ {
        ch <- "decrement"
    }
}




func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1? It chooses how many OS-threads can run simoultanelesly. If you set it to one, only one process may run at the same time, so you dont get race conditions. 
    
    runtime.GOMAXPROCS(2)  
    var wg sync.WaitGroup

	
    // // TODO: Spawn both functions as goroutines

    requestChannel :=make(chan string)
    responseChannel := make(chan int)

    go server(requestChannel, responseChannel)

    wg.Add(2) //Two go routines
    go decrementing(requestChannel, &wg)
    go incrementing(requestChannel, &wg)

    wg.Wait()
    
	


    requestChannel <- "get"
    finalValue:= <- responseChannel 

    Println("The magic number is:", finalValue)

    requestChannel  <- "done"
    close(requestChannel)

    

}
