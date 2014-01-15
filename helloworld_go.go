// Go 1.2
// go run helloworld_go.go

package main

import (
	    . "fmt" // Using '.' to avoid prefixing functions with their package names
	    . "runtime" // This is probably not a good idea for large projects...
	    . "time"
)

var i = 0

func adder(ic chan int) {
	for x := 0; x < 1000000; x++ {
	       	i =<- ic
		i++
		ic <- i
	}
}

func subtractor(ic chan int) {
	for x := 0; x < 500000; x++ {
		i =<- ic
		i--
		ic <- i		
	}
}
func main() {
	    
	    ic := make(chan int,1)
	    
	    ic <- i
	    GOMAXPROCS(NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	    go adder(ic) // This spawns adder() as a goroutine
	    go subtractor(ic)
	    
	    for x := 0; x < 50; x++ {
		Println(i)
	    }
	    // No way to wait for the completion of a goroutine (without additional syncronization)
	    // We'll come back to using channels in Exercise 2. For now: Sleep
	    Sleep(1000*Millisecond)
	    i:= <-ic
	    Println("Done:", i);
}
