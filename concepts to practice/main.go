package main

import "fmt"






func main() {
	// friends := &AddressBook{}
	// friends.AddPerson("Ariel", 27, "Minneapolis")
	// friends.AddPerson("Tom", 28, "Minneapolis")
	// friends.AddPerson("Alex", 27, "Minneapolis")

	// friends.ListPeople()

	// friends.UpdateCity("Ariel", "Chicago")
	// friends.UpdateCity("Tom", "New York")

	// friends.ListPeople()

	// var first string = "Name"
	// var number int = 26

	// fmt.Printf("Name is: %s | Age is: %d\n", first, number)

	// FillData(&first, &number)

	// fmt.Printf("Name is: %s | Age is: %d\n", first, number)

	// student1 := &Student{
	// 	Name: "Ariel",
	// 	Grade: 95,
	// }

	// fmt.Printf("Name of Student: %s\nGrade of Student:%d\n", student1.Name, student1.Grade)
	// student1.UpdateGrade(100)
	// fmt.Printf("Name of Student: %s\nGrade of Student:%d\n", student1.Name, student1.Grade)
	// UpdateGradeNoPointer(&student1.Grade, 90)
	// fmt.Printf("Name of Student: %s\nGrade of Student:%d\n", student1.Name, student1.Grade)


	// var original int = 42
	// fmt.Println(original)
	// tryToChange(original)
	// fmt.Println(original)
	// actuallyChange(&original)
	// fmt.Println(original)

	// var post1 BlogPost

	// FakeScan(&post1.ID, &post1.Title, &post1.Content)

	// fmt.Println(post1)

	// post2 := &BlogPost{

	// }

	// FakeScan(&post2.ID, &post2.Title, &post2.Content)
	// fmt.Println(post2)


	// counterTest := &Counter{
	// 	count: 0,
	// }

	// fmt.Printf("%d\n", counterTest.count)
	// counterTest.Increment()
	// fmt.Printf("%d\n", counterTest.count)

	// valueOne := 2
	// fmt.Println(valueOne)
	// DoubleScored(&valueOne)
	// fmt.Println(valueOne)

	myPlayer := &Player{
		name: "Ariel",
		health: 100,
	}

	fmt.Println(myPlayer.health)
	reduceHealth(*myPlayer)
	fmt.Println(myPlayer.health)


}


