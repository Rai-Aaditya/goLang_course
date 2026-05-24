package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// we can omit the value of the receiver i.e, eb as we are not using it
// func (eb englishBot) getGreeting() string{
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hello there!"
}

// func (sb spanishBot) getGreeting string{
func (spanishBot) getGreeting() string {
	return "Hola!"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// rewriting printgreeting function using interface logic:
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
