package basics

/*
type name_interface interface{
	some_function
}
*/

/*
	Example 1:

	type Vehicle interface {
		move()
	}

	func drive(vehicle Vehicle) {
		vehicle.move()
	}

	type Car struct {}

	func (c Car) move() {
		fmt.Println("car move")
	}

	tesla := Car{}
	drive(tesla)
*/

/*
	Example 2:

	type Vehicle interface {
		move()
	}

	func drive(vehicle Vehicle) {
		vehicle.move()
	}

	type Car struct {model string}
	type Aircraft struct {model string}


	func (c Car) move() {
		fmt.Println("car move")
	}

	func (a Aircraft) move() {
		fmt.Println("aircraft fly")
	}


	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	air := Aircraft{"365"}
	vehicles := [...]Vehicle{tesla, volvo, air}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
*/
