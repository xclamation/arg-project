package main

import (
	"flag"
	"fmt"
	"arg-project/arg"
)

func main() {
	
	if !arg.Shhh {
		fmt.Print("Non-flag commands: ")
		fmt.Print(flag.Args(), "\n")
		fmt.Printf("Hello, %s! %d years old\n", arg.Name, arg.Age)
		fmt.Printf("Your car is %s with speed %d", arg.Car.Brand, arg.Car.Speed)
	}
}
