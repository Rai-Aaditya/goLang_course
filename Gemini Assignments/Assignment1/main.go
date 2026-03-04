package main

import "fmt"

func main() {
	tasks := newTaskList()

	tasks.printTasks()
	fmt.Println("Result for hastask:", tasks.hasTask("Write code"))
	fmt.Println("Concatenated string: ", tasks.toString())
}
