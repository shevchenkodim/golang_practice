package basics

import "fmt"

/*
	var name_variable type = value
	const name_variable type = value

	var (
		name_variable type = value
		name_variable_2 type_2 = value_2
	)
	const (
		name_variable type = value
		name_variable_2 type_2 = value_2
	)

	Short variable definition
	name_variable := value
*/

func printVariables() {
	var hello string = "Hello world"
	var (
		name string = "Tom"
		age  int    = 27
	)

	name2 := "Tom"

	fmt.Println(hello, name, age, name2)
}
