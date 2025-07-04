package main

import (
	"fmt"

	goset "github.com/pydea-rs/goset"
)

func main() {
	s := goset.NewSet()
	s.Add("String")
	s.Add(10)
	s.Add(3.14)
	s.Add(struct{ x int }{x: 1})
	fmt.Println(s)
	x := make([]string, 0)
	s.Add(x)
	if !s.Add(struct{ x int }{x: 1}) {
		fmt.Println("struct{ x int }{x: 1} is repetitive! Set count:", s.Count())
	}

}
