package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Using wait groups
func boringWG(fn1, fn2 func(i int)) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fn1(1)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		fn2(2)
	}()
	wg.Wait()
}

// TODO: Done channel
// TODO: Buffered channels

//func main() {
//	boringWG(func(i int) { fmt.Println("fn1") }, func(i int) { fmt.Println("fn2") })
//}

// Generator:
// function that returns a channel
func boring(msg string) <-chan string {
	c := make(chan string)
	r := rand.New(rand.NewSource(99))
	go func() {
		// some iterator
		for iterator := 0; iterator < 5; iterator++ {
			c <- fmt.Sprintf("%s %d", msg, iterator)
			// random sleep
			time.Sleep(time.Duration(r.Int()) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

//func main() {
//	c := boring("a")
//	fmt.Printf("said - %s", <-c)
//}

// They basically function as a service
//func main() {
//	joe := boring("Joe")
//	ann := boring("Ann")
//	for i := 0; i < 5; i++ {
//		fmt.Println(<-joe)
//		fmt.Println(<-ann)
//	}
//}

// These program make the two count in lock-step. To get around this we can use a fan-in function

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// independent communication
//func main() {
//	c := fanIn(boring("Joe"), boring("Ann"))
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-c)
//	}
//}

//If you wanted everything to be lock-step and sync
// lockstep

type Message struct {
	str  string
	wait chan bool
}

func lockeStep(msg1, msg2 Message) {
	// example
	c := make(chan Message)

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

	waitForIt := make(chan bool) // shared between all channels
	msg := "hello"
	iterator := 0
	c <- Message{fmt.Sprintf("%s %d", msg, iterator), waitForIt}
	time.Sleep(1 * time.Second)
	<-waitForIt

	select {
	case <-waitForIt:
		// do something
	}
}

//Timeout using select
//	select {
//		case <-time.After(8 * time.Second):
//		return
//	}

// Quit channel
// 	quit := make(chan bool)
// 	do something
// 	quit <- true

//	select {
//		case <-quit:
//			return
//	}

// Return on quit channel
// quit := make(chan bool)
// c := boring("Joe", quit)
// do stuff
//	quit <- "bye"
//fmt.Printf("Joe says: %q\\n", <-quit)
//
//select {
//case c <- fmt.Sprintf("%s: %d", msg, i):
//// do something
//case <-quit:
//cleanup()
//quit <- "See you"
//return
//}
// receives on quit channel, then he sends to the quit channel so the "Joe says" statement
// executes and program exits

// Daisy chain

// Google search example:
var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		return Result(fmt.Sprintf(kind))
	}
}

func Google(query string) []Result {
	results := make([]Result, 0)
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return results
}

// Pattern: fan-in

func GoogleFanIn(query string) (results []Result) {
	// fan-in
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

// Pattern: timeout pattern
// don’t wait for slow servers!

func GoogleTimeout(query string) (results []Result) {
	// fan-in
	c := make(chan Result)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	// timeout pattern
	timeout := time.After(200 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}
	return results
}

// Pattern: replication
// Q: how do we avoid timeouts from slow events
// A: Replicate the servers.Send requests to multiple replicas, and use the first response

func First(query string, replicas ...Search) Result {
	// replicas of a single search
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		// launch the same search multiple times
		go searchReplica(i)
	}
	// return the first response
	return <-c
}

func GoogleReplication(query string) (results []Result) {
	// fan-in
	c := make(chan Result)
	go func() { c <- First(query, Web, Web) }()
	go func() { c <- First(query, Image, Image) }()
	go func() { c <- First(query, Video, Video) }()

	// timeout pattern
	timeout := time.After(200 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}
	return results
}

func main() {
	s := Google("golang")
	fmt.Println(s)
	fn := GoogleFanIn("golang")
	fmt.Println(fn)
	tim := GoogleTimeout("golang")
	fmt.Println(tim)
	rep := GoogleReplication("golang")
	fmt.Println(rep)
}
