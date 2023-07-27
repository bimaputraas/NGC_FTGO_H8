package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	job  string
}

func (p Person) GetInfo() {
	if p.age >= 50 {
		p.job = "Retired"
	}

	fmt.Printf("Name: %s\nAge: %d\nJob: %s\n", p.name, p.age, p.job)

}

func (p *Person) AddYear() {
	p.age = p.age + 1
}

func main() {
	var person1 Person
	person1 = Person{
		name: "Bambang",
		age:  48,
		job:  "Gambler",
	}

	var person2 Person
	person2 = Person{
		name: "Wahyu",
		age:  30,
		job:  "Teacher",
	}

	var person3 Person
	person3 = Person{
		name: "Ucup",
		age:  34,
		job:  "Pilot",
	}

	var person4 Person
	person4 = Person{
		name: "Udin",
		age:  18,
		job:  "Atlet",
	}

	// soal 1
	fmt.Println("Soal 1 :")
	person1.GetInfo()

	fmt.Println("Setelah add year :")
	person1.AddYear()
	person1.GetInfo()
	person1.AddYear()
	person1.GetInfo()

	// soal 2
	fmt.Println("\nSoal 2 :")

	slice_of_person := []Person{
		person1, person2, person3, person4,
	}

	for _, person := range slice_of_person {
		person.GetInfo()
		fmt.Println("")
	}

}
