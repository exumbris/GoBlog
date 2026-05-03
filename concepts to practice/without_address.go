package main

func tryToChange(val int) {
	val = 999
	// fmt.Println(val)
}

func actuallyChange(val *int) {
	*val = 999
	// fmt.Println(*val)
}