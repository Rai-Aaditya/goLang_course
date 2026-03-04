package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type taskList []string

func newTaskList() taskList {
	tasks := taskList{}
	var count int
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter number of tasks you want to enter: ")
	fmt.Scan(&count)

	reader.ReadString('\n')

	for i := 0; i < count; i++ {
		fmt.Printf("Enter task %v: ", i)
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		tasks = append(tasks, task)
	}

	return tasks
}

func (t taskList) printTasks() {
	for _, i := range t {
		fmt.Println(i)
	}
}

func (t taskList) hasTask(val string) bool {
	return slices.Contains(t, val)
}

func (t taskList) toString() string {
	return strings.Join(t, ",")
}
