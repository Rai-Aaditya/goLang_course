## 🎴 Gemini Go Assignment: Task Tracker

Welcome to your **Task Tracker Assignment**! 🚀
This fun challenge is inspired by the Cards Chapter from the Udemy course. Let's put your Go skills to the test! 💪

---

### 📝 **Problem Statement**
**Design and implement a simple task management utility in Go.**

---

### 🎯 **Objective**
Show off your understanding of:
- **Custom types**
- **Functions with receivers (methods)**
- **Functions without receivers**

---

### 📚 **Background**
In Go:
- Functions **without receivers** are for creating/initializing data structures.
- Functions **with receivers** act on instances of custom types.

You'll use both in this assignment!

---

### 🛠️ **Requirements**
Implement these components in your Go program:

#### 1️⃣ **Custom Type Definition**
Define a custom type named **`taskList`** whose underlying type is a slice of strings (`[]string`).

#### 2️⃣ **Initialization Function**
Create a function called **`newTaskList`**:
- **Signature:** `func newTaskList() taskList`
- **Behavior:** Returns a new `taskList` with at least three default tasks (e.g., "Setup Go environment", "Read documentation", "Write code")
- **Note:** This function should **not** have a receiver.

#### 3️⃣ **Display Method**
Add a method called **`printTasks`**:
- **Receiver:** Variable of type `taskList`
- **Behavior:** Iterate over the `taskList` and print each task with its index number.

#### 4️⃣ **Search Method**
Add a method called **`hasTask`**:

**Receiver:** Variable of type `taskList`
**Arguments:** A string representing the task description to search for
**Returns:** Boolean value
**Behavior:** Returns `true` if the exact string exists in the `taskList`, `false` otherwise

#### 5️⃣ **Main Execution**
In your `main()` function:
1. Instantiate a new task list using `newTaskList()`
2. Call `printTasks()` on the list
3. Use `hasTask()` to check for a specific task and print the result

---

### 🎁 **Bonus Challenge**
Add a method called **`toString`**:
- **Receiver:** Variable of type `taskList`
- **Returns:** A single string
- **Behavior:** Converts the entire `taskList` into a comma-separated string (e.g., "Task 1, Task 2, Task 3")
- 💡 *Hint:* Check out the `strings` package in Go!

---

Have fun coding! 😄✨