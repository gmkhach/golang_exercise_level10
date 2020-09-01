package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		        Exercise 1
		        Get this code working...
			            c := make(chan int)
			            c <- 42
			            fmt.Println(<-c)
		            a. using func literal, aka, anonymous self-executing func
		            b. using a buffered channel
	*/
	// a.
	c1 := make(chan int)
	go func() {
		c1 <- 42
	}()
	fmt.Println(<-c1)

	// b.
	c2 := make(chan int, 1)
	c2 <- 42
	fmt.Println(<-c2)

	/*
		        Exercise 2
		        Get this code running:
			        cs := make(chan<- int)

			        go func() {
				        cs <- 42
			        }()
			        fmt.Println(<-cs)

			        fmt.Printf("------\n")
			        fmt.Printf("cs\t%T\n", cs)
	*/
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)

	/*
		        Exercise 3
		        Start with this code and pull the values off the channel using a for range loop:
		        func main() {
			        c := gen()
			        receive(c)

			        fmt.Println("about to exit")
		        }

		        func gen() <-chan int {
			        c := make(chan int)

			        for i := 0; i < 100; i++ {
				        c <- i
			        }

			        return c
		        }
	*/
	c3 := gen1()

	receive1(c3)

	fmt.Println("about to exit")

	/*
		        Exercise 4
		        Start with this code and pull the values off the channel using a for range loop:
		        func main() {
			        q := make(chan int)
			        c := gen(q)

			        receive(c, q)

			        fmt.Println("about to exit")
		        }

		        func gen(q <-chan int) <-chan int {
			        c := make(chan int)

			        for i := 0; i < 100; i++ {
				        c <- i
			        }

			        return c
		        }
	*/
	q := make(chan int)
	c4 := gen2(q)

	receive2(c4, q)

	fmt.Println("about to exit")

	/*
	        Exercise 5
	        Show the comma ok idiom starting with this code:
		        c := make(chan int)

		        v, ok :=
		        fmt.Println(v, ok)

		        close(c)

		        v, ok =
		        fmt.Println(v, ok)
	*/
	c5 := make(chan int)

	go func() {
		c5 <- 42
	}()

	v, ok := <-c5
	fmt.Println(v, ok)

	close(c5)

	v, ok = <-c5
	fmt.Println(v, ok)

	/*
	   Exercise 6
	   Write a program that...
	       a. puts 100 numbers to a channel
	       b. pulls the numbers off the channel and prints them
	*/
	c6 := make(chan int)

	go func() {
		for i := 1000; i < 1100; i++ {
			c6 <- i
		}
		close(c6)
	}()

	for v := range c6 {
		fmt.Println(v)
	}

	/*
	   Exercise 7
	   Write a program that ...
	       a. launches 10 goroutines, each one of which add 10 numbers to a channel
	       b. pulls the numbers off the channel and prints them
	*/
	c7 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(10)

	go func() {
		for i := 0; i < 10; i++ {
			go func(i int) {
				for j := 0; j < 10; j++ {
					c7 <- (i*10 + j)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c7)
	}()

	for v := range c7 {
		fmt.Println(v)
	}
}

func gen1() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive1(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}

func gen2(q chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive2(cr, q <-chan int) {
	for {
		select {
		case v := <-cr:
			fmt.Println(v)
		case v := <-q:
			fmt.Println("done and quitting:", v)
			return
		}
	}
}
