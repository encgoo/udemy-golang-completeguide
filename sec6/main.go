package main

import (
	"fmt"
)

// an empty type with only receiver functions
type englishBot struct{}
type spanishBot struct{}

// These two have different logic
func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// different logic shall be summarized as interface
type bot interface {
	getGreeting() string
}

// similar logic shall programmed to interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// interface can be put into another struct
type machine struct {
	Name string
	Mbot bot // interface can be used as a type inside struct
}

func (m machine) printGreeting() {
	fmt.Println(m.Mbot.getGreeting())
}

// interface can be grouped into other interface
type userInterface interface {
	// Needs to take input
	setInput(string) bool
}

// machInterface is both
type machInterface interface {
	bot
	userInterface
}

// machineStrct satisfies machInterface
type machineStrct struct{}

func (m machineStrct) getGreeting() string {
	return "Hello World"
}

func (m machineStrct) setInput(string) bool {
	return true
}

func main() {
	var eb englishBot
	var sb spanishBot

	printGreeting(eb)
	printGreeting(sb)

	var mach machine
	// eb is an english bot that satisfies bot
	mach.Mbot = eb
	mach.printGreeting()

	var mach1 machineStrct
	// compiler checks that machineStruct satisfies machInterface at this point
	var machInf machInterface = mach1
	fmt.Println(machInf.getGreeting())
}
