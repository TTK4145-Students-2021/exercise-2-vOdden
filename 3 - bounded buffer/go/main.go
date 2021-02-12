
package main

import "fmt"
import "time"

var done = make(chan bool)


func producer(msg chan<- int){

  for i := 0; i < 10; i++ {
      time.Sleep(100 * time.Millisecond)
      fmt.Printf("[producer]: pushing %d\n", i)
      // TODO: push real value to buffer
  msg <- i
  }

        done <- true
}

func consumer(msg <-chan int){

  time.Sleep(1 * time.Second)
  for {
      // Get real value from buffer
      fmt.Printf("[consumer]: %d\n", <-msg)
      time.Sleep(50 * time.Millisecond)
  }

}

func main(){

    done := make(chan bool) // Signals when done.

    // Bounded buffer
    buffer := make(chan int, 5)

    go consumer(buffer)
    go producer(buffer)
    <-done
    fmt.Println("Done")
    select {}
}
