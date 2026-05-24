package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// this method works but if we want to check through the links multiple times, then it will come back to the link at particular index after checking through all the links in the slice and will then check the link again.

	// in this case, the links are checked in a serial order
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // we only use go keywords in front of function calls
	}
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// putting a function literal betwen the channel loop to pause the code for some time and then continue for next routing

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

/*
	Our main code is considered as a single go routine which compiles(takes) whole code at a time and executes each line of code one by one

	The code line
		_, err := http.Get(link)
	is considered as a blocking code which blocks the process for the time at which the response is fetched.

	To mitigate this bottleneck we add additional go routine with the syntax:
		go checkLink(link)
	by adding go before the function call.
	What this does is it creates an additional go routine which has the function body and when it encounters the blocking code which is the get request line, this go routines goes to sleep and waits for the response to come transferring the control to parent go routine(main code)

	in the parent routine, next iteration of for loop occurs and again for the function call a new additional go routine gets created(new routine for each encounter of go keyword)


	go keyword creates a new thread go routine and runs the function within it

	if we just use go keyword, the main app will run and exit without caring about other go routines. To enable communication between different routines, we use channels
	channels are type specific just like other data types like slice of strings. if we specify channel of type string, we can not pass integer value using it.
*/
