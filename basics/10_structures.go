package basics

/*
	type name_struct struct {
		attributes
	}
*/

/*
	Example 1:

	type person struct {
		name string
		age int
	}

	p1 := person{"Dmytro", 22}
	fmt.Println(p1.age, p1.name)
*/

/*
	Pointers to Structures[Example 2]:
	dmytro := person{"Dmytro", 22}
	var dmytroPointer *person = &dmytro

	dmytroPointer.age = 25
	fmt.Println(dmytro.age)
*/
