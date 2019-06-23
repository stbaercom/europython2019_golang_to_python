package main

import ml "com.stbaer/demo_go/main_lib_simple"

func main() {


	person := ml.NewPerson("Stefan","Baerisch",41)

	println(person.String())

}