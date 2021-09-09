package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("in a goroutine")
	fmt.Println("in main")
	// task 1. if we run just the abovem, in go routie will not print cause lol this may be wrong
	// but i believe it's cause the main func finished running before the go routine finished so it never happens

	// for i := 0; i < 5; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	// task 3. this will only print out the number 5 because each go routine has different scope

	for i := 0; i < 5; i++ {

		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	// task 4. this one will count down because we are passing in i as a parameter when we are calling the annonymous go func
	// when you use go keyword inside function you start a concurrently running thread

	time.Sleep(10 * time.Millisecond)
	// task 2. now it will print 'in go routine' because we have delayed main from finishing

	ch := make(chan int)
	go func() {
		ch <- 7 // send
	}()
	val := <-ch // receive
	fmt.Println("got:", val)

	queue := make(chan string)
	for i := 0; i < 3; i++ {
		go producer(i, queue)
	}

	for i := 0; i < 2; i++ {
		go consumer(i, queue)
	}

	time.Sleep(time.Second)
}

func producer(id int, ch chan<- string) {
	i := 0
	for {
		time.Sleep(100 * time.Millisecond) // simulate work
		i++
		msg := fmt.Sprintf("%d -> [%d]\n", id, i)
		ch <- msg
	}
}

func consumer(id int, ch <-chan string) {
	for {
		time.Sleep(100 * time.Millisecond) // simulate work
		msg := <-ch
		fmt.Printf("%s -> %d\n", msg, id)

	}

}
