package main

// help me import greetings package
import (
	"example/home/calculate"
	"example/home/greetings"
)

func main() {

	message := greetings.Hello()
	println(message)

	sum := calculate.Add(12, 32)
	println("sum is:", sum)
}
