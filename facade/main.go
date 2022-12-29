package main

import "fmt"

type Greeter struct {
	greetingDb          *GreetingDb
	greetingConstructor *GreetingConstructor
}

func NewGreeter() *Greeter {
	return &Greeter{
		&GreetingDb{},
		&GreetingConstructor{},
	}
}

func (g *Greeter) greet() {
	//get the name of the person to greet //come from db
	//construct a greeting // Hello + name of the person
	fmt.Println("Hello")
}

func main() {
	g := NewGreeter()
	g.greet()
}

type GreetingDb struct {
}

func (g *GreetingDb) getGreeterName() string {
	return "John"
}

type GreetingConstructor struct {
}

func (g *GreetingConstructor) getGreeting(name string) string {
	return "Hello " + name
}
