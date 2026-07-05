package main

import "fmt"

func main()  {
	greet := greet()
	fmt.Println(greet)
}

func greet() string {
	return "Hello world"
}
