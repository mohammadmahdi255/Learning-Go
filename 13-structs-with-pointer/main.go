package main

import "fmt"

type Older interface {
	BeOlderPointer(int)
}

type Person struct {
	Name string
	Age  int
}

func New(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) BeOlder(increase int) Person {
	p.Age += increase
	return p
}

func (p *Person) BeOlderPointer(increase int) {
	p.Age += increase
}

// enable this and secction [1] to see changes
func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Mohammad Mahdi Nemati",
		Age:  21,
	}

	fmt.Printf("before (p) Person %p : %v\n", &p, p)

	z := p.BeOlder(1)
	p.BeOlderPointer(1)

	fmt.Printf("after (p) Person %p : %v\n", &p, p)
	fmt.Printf("after (z) Person %p : %v\n", &z, z)

	var bo Older

	// compile error: cannot use p (variable of type Person) as type Older in assignment:
	// Person does not implement Older (BeOlderPointer method has pointer receiver)
	// bo = p

	bo = &p

	bo.BeOlderPointer(10)

	fmt.Printf("after (p) Person %p : %v\n", &p, p)

	fmt.Printf("ُString on Person %p : %v\n", &p, p)
	fmt.Printf("ُString on *Person %p : %v\n", &p, &p)

	// secction [1]
	var _ fmt.Stringer = p
	var _ fmt.Stringer = &p

}
