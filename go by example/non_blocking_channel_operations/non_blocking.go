// Basic sends and receives on channels are blocking. However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here’s a non-blocking receive. If a value is available on messages then select will take the <-messages case with that value. If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// A non-blocking send works similarly. Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases above the default clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both messages and signals.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

/*
I have a doubt on this, in the second select, it is mentioned that   msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver. Therefore the default case is selected

but we have a receiver for messages channel in the next i.e., 3rd select so ideally it should take the msg value

also if it has taken the msg value then in 3rd select, the receive for msg shall work


This is a very common point of confusion when learning Go channels! You have bumped into the difference between how code executes **sequentially** and how unbuffered channels operate **synchronously**.

To answer your question directly: **Block 2 fails to send because Block 3 hasn't happened yet.**

Here is the breakdown of exactly why this happens, step-by-step.

### 1. Unbuffered Channels Need Concurrent Partners

You created an unbuffered channel: `messages := make(chan string)`.
For a send to an unbuffered channel to succeed, a receiver must be actively waiting on the other end **at that exact moment**. It’s like passing a baton in a relay race; both runners must have their hands on the baton at the same time. You cannot drop the baton on the ground and expect the next runner to pick it up later.

### 2. Code Executes Sequentially

Your code runs inside a single goroutine (the `main` function) from top to bottom. It cannot execute Block 2 and Block 3 at the same time.

Here is the timeline of execution:

* **Time 1 (Block 1):** You try to receive. Nothing is there. It hits `default`.
* **Time 2 (Block 2):** You try to send `"hi"`. The `select` statement asks, "Is anyone waiting to receive this *right now*?" Because Block 3 hasn't started executing yet, the answer is no. Because it is unbuffered, the send would normally block, but since you have a `default` case, it abandons the send entirely and prints `"no message sent"`.
* **Time 3 (Block 3):** Now you try to receive. But because the send in Block 2 was abandoned, the channel is empty. It hits `default` and prints `"no activity"`.

### How to make your logic work

If you want Block 3 to receive the message sent by Block 2, you have two options:

#### Option A: Use a Buffered Channel

If you give the channel a buffer, you are essentially providing a "mailbox" where you can drop the baton without the receiver being there yet.

```go
// Change the channel creation to include a buffer size of 1
messages := make(chan string, 1)

```

If you make this single change to your code, Block 2 will successfully drop `"hi"` into the buffer, and Block 3 will successfully pull `"hi"` out of the buffer!

#### Option B: Use a Goroutine (Concurrency)

If you want to keep the unbuffered channel, the receiver must be waiting concurrently. You would need to put the receiver in a separate goroutine so it is running *at the same time* Block 2 tries to send.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	// Fire off a goroutine that waits for a message
	go func() {
		// This will wait here until someone sends a message
		msg := <-messages
		fmt.Println("received message in goroutine:", msg)
	}()

	// Give the goroutine a tiny moment to spin up and wait
	time.Sleep(time.Millisecond * 10)

	msg := "hi"
	select {
	case messages <- msg: // This succeeds now because the goroutine is waiting!
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
}

```

**Summary:** In a single goroutine with unbuffered channels, `select` statements cannot look ahead to see what code is coming next. They only know what is happening in that exact millisecond.

*/
