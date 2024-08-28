package arg

import (
	"flag"
	"arg-project/car"
)

var Name string
var Shhh bool
var Age int
var Car car.Car

func init() {
	flag.StringVar(&Name, "name", "User", "specify the name for greet")
	flag.IntVar(&Age, "age", int(1), "specify the age")
	flag.BoolVar(&Shhh, "shhh", false, "specify whether it is necessary to output")
	flag.Var(&Car, "car", "specify the car in the format Brand:Speed")
	flag.Parse()
}




// func FlagsDeclaration() {
// 	flag.StringVar(&Name, "name", "User", "specify the name for greet")
// 	flag.IntVar(&Age, "age", int(1), "specify the age")
// 	flag.BoolVar(&Shhh, "shhh", false, "specify whether it is necessary to output")
// }
