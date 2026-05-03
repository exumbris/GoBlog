package main

type BlogPost struct {
	ID int
	Title string
	Content string
}

func FakeScan(idPtr *int, titlePtr *string, contentPtr *string) {
	*idPtr = 1
	*titlePtr = "My First Post"
	*contentPtr = "Hello, World!"
}

