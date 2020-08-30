package dating

import (
	"fmt"
)

func Mytest() {
	fmt.Printf("my test\n")

}

/*
Dating App:

type Person struct {
	animal
	pref map[]
}

map[string]Person{}


*/

type Person struct {
	Animal    string
	Pref      map[string]int
	Locattion string
}

func CreateDatingThingy() map[string]Person {

	znoodle := map[string]Person{}

	znoodle["Mouse"] = Person{Animal: "cat"}
	return znoodle
}

func Warmup() map[string]int {

	testermap := map[string]int{}

	testermap["mousey"] = 3
	testermap["noodle"] = 2
	return testermap
}
