// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() {
		val := 10 * 5001
		fmt.Println(val)
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
	fmt.Println("end")
}

// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

/*
qna
in this code, can it happen that the end prints first and then the function and channel value

No, that cannot happen. The output order of your program is strictly guaranteed to always be the same.

Here is exactly why "end" will never print first, step-by-step:

1. Unbuffered Channels Block Execution
When you create the channel with messages := make(chan string), you are creating an unbuffered channel. By design, sending and receiving on an unbuffered channel are synchronous, blocking operations.

2. The Main Goroutine is Forced to Wait
When the main function reaches the line msg := <-messages, it completely pauses (blocks). It cannot move on to print "ping" or "end" until it actually receives a value from that channel.

3. The Guaranteed Sequence of Events
Because the main function is paused waiting for a message, the following sequence is guaranteed:

The Goroutine Executes First: The spawned goroutine runs, calculates val := 10 * 5001, and prints it (50010).

The Handshake: The goroutine sends "ping" into the channel (messages <- "ping").

The Main Goroutine Wakes Up: The main function receives "ping", unblocks, and resumes execution.

Sequential Printing: The main function then sequentially executes its remaining lines: it prints the message ("ping"), and finally prints "end".

Because the channel forces the main function to wait for the background goroutine to finish its work and send the message, the output will always reliably be:

50010
ping
end
*/
