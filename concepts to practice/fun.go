package main

import "fmt"

func Birthday(p *Person) {
	p.Age += 1
}

func (person *Person) Move (newCity string) {
	person.City = newCity
}

func (peopleList *AddressBook) AddPerson(name string, age int, city string) {
	newPerson := &Person {
		Name: name,
		Age: age,
		City: city,
	}
	peopleList.people = append(peopleList.people, newPerson)
}

func (peopleList *AddressBook) UpdateCity(name string, newCity string) {
	for _, p := range peopleList.people {
		if p.Name == name {
			p.City = newCity
			return	
		}
	}
}

func (peopleList *AddressBook) ListPeople() {
	for i, p := range peopleList.people {
		fmt.Printf("Person #%d\n", i)
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Age: %d\n", p.Age)
		fmt.Printf("City: %s\n", p.City)
		fmt.Println("--------------------")
	}
}