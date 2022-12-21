package main

import "fmt"

type Student struct {
	Name   string
	Family string
	age    int
}

// New is a popular pattern to create types
func (s Student) New(name, family string, age int) Student {
	return Student{
		Name:   name,
		Family: family,
		age:    age,
	}
}

func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Family: %s, age: %d", s.Name, s.Family, s.age)
}

func (s Student) Hello(to Student) string {
	return fmt.Sprintf("Hello %s, I am %s %s (%d) ", to.Name, s.Name, s.Family, s.age)
}

func main() {

	s := Student{
		Name:   "Mahdi",
		Family: "Nemati",
		age:    21,
	}

	fmt.Println(s)
	fmt.Println(s.Hello(Student{
		Name: "Amin",
	}))
}
