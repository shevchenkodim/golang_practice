package basics

/*
	func (param_name receiver_type) method_name (params) (returns_types){
		body method
	}
*/

/*
	Example 1:
	type library []string

	func (l library) printAll() {
		for _, val := range l {
			fmt.Println(val)
		}
	}

	var lib library = library{"1", "2", "3"}
	lib2 := library{"11", "22", "33"}
	lib.printAll()
	lib2.printAll()
*/

/*
	Example 2:
	type person struct {
		name string
		age int
	}

	func (p person) printInfo() {
		fmt.Println(p.name)
		fmt.Println(p.age)
	}

	func (p *person) setNewAge(age int) {
		(*p).age = age
	}

	var dim = person{name: "Dmytro", age: 22}
	var dimPointer *person = &dim
	dim.printInfo()
	dimPointer.setNewAge(25)
	dim.printInfo()
*/
