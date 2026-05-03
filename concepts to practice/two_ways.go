package main

// exercise two
type Student struct {
	Name string
	Grade int
}

func (s1 *Student) UpdateGrade(newGrade int) {
	s1.Grade = newGrade
}

func UpdateGradeNoPointer(oldGrade *int, newGrade int) {
	*oldGrade = newGrade
}

