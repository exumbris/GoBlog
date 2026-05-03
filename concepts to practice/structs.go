package main

type Person struct {
	Name string
	Age int
	City string
}

type AddressBook struct {
	people []*Person
}