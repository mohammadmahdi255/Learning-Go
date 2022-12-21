package main

import "fmt"

type Printer interface {
	Print() string
}

type increaser interface {
	increaser(int)
}

type Person struct {
	Name string
}

type Student struct {
	Name string
}

func (p Person) Print() string {
	return fmt.Sprintf("Person %s", p.Name)
}

func (s Student) Print() string {
	return s.Name
}

func (s Student) Hello() string {
	return "Hello"
}

type Any interface{}

func main() {
	var p Printer
	s := Student{
		Name: "Mahdi Nemati",
	}

	s.Hello()

	p = s

	fmt.Println(p.Print())

	s = p.(Student)

	_, ok := p.(Person)
	if !ok {
		fmt.Println("p is not Person")
	}
}
