package main

import "fmt"

type IGreeter interface {
	Greet()
}

type Greeter struct {
	greetingFormat   GreetingFormat
	greetingStrategy GreetingStrategy
}

func NewGreeter(greetingFormat GreetingFormat) IGreeter {
	g := &Greeter{greetingFormat: greetingFormat}
	switch greetingFormat {
	case Formal:
		g.greetingStrategy = &FormalStrategy{}
	case Informal:
		g.greetingStrategy = &InformalStrategy{}
	}

	return g
}

func (g *Greeter) Greet() {

	db := &DB{}
	name := db.GetName()

	s := g.greetingStrategy
	s.PrintGreeting(name)
}

type GreetingFormat int

const (
	Formal GreetingFormat = iota
	Informal
)

type GreetingStrategy interface {
	PrintGreeting(name string)
}

type FormalStrategy struct{}

func (f *FormalStrategy) PrintGreeting(name string) {
	fmt.Printf("Dear %s\n", name)
}

type InformalStrategy struct{}

func (f *InformalStrategy) PrintGreeting(name string) {
	fmt.Printf("Hi %s\n", name)
}

type DB struct {
}

func (db *DB) GetName() string {
	return "John"
}

func main() {

	// greeting

	// get Name from DB

	// say greeting in formal way // Dear + Name
	// say greeting in informal way // Hi + Name

	g := NewGreeter(Informal)
	g.Greet()
}
