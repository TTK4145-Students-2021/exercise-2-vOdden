
package main

import "fmt"
import "time"

var done = make(chan bool)
var msg = make(chan int,5)

func producer(){

    for i := 0; i < 10; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("[producer]: pushing %d\n", i)
        // Push real value to buffer
        msg <- i
      }
        close(msg)
        done <- true
}

func consumer(){

    time.Sleep(1 * time.Second)
    for msg := range msg {
        fmt.Printf("[consumer]: %d\n", msg)
        time.Sleep(50 * time.Millisecond)
    }

}

func main(){

    // TODO: make a bounded buffer

    go consumer()
    go producer()
    <-done
    fmt.Println("Done")
    //select {}
}
