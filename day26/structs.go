package main

import "fmt"

// Person represents a person
type Person struct {
	// Name is public
	Name string
	// Email is private
	email string
	// map initialize with nil
	Phones map[string]string
}

// visibility is about package
// if you import person from another package,
// email is not public.

// NewPerson is something like a constructor
// very useful to initialize some fields, like in our case, phones
func NewPerson(name, email string) *Person {
	//returns a pointer to struct Person
	return &Person{
		Name:   name,
		email:  email,
		Phones: make(map[string]string),
	}
}

func main() {
	// instiate a person
	p := NewPerson("cassiobotaro", "cassiobotaro@gmail.com")
	// added some phones(set)
	p.Phones["casa"] = "123456789"
	p.Phones["celular"] = "987654321"
	fmt.Printf("p = %+v\n", p)
	// get Name
	fmt.Printf("p.Name = %+v\n", p.Name)
	// get Phones
	for k, v := range p.Phones {
		fmt.Printf("%s: %s\n", k, v)
	}
}
