package main

import "fmt"

type PersonA struct {
	FirstName, MiddleName, LastName string
}

// func (p *PersonA) Names() [3]string{
// 	return [3]string{p.FirstName, p.MiddleName, p.LastName}
// }

func(p *PersonA) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.FirstName
		if len(p.MiddleName) > 0 {
			out <- p.MiddleName
		}
		out <- p.LastName
	}()
	return out
}

func main() {
	p := PersonA{"Alex", "Graham", "Bell"}
	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}
}
