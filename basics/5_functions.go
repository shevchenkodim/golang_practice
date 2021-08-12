package basics

/*
	func name_func (params & type) (return_param & type){
		operator
	}


	func add(x int, y int){
		var z = x + y
		fmt.Println("x + y = ", z)
	}
	add(4, 5)   // x + y = 9


	func add(x, y int, a, b, c float32){
		var z = x + y
		var d = a + b + c
		fmt.Println("x + y = ", z)
		fmt.Println("a + b + c = ", d)
	}
	add(1, 2, 3.4, 5.6, 1.2)


	func add(numbers ...int){
		var sum = 0
		for _, number := range numbers{
			sum += number
		}
		fmt.Println("sum = ", sum)
	}
    add(1, 2, 3)        // sum = 6
    add(1, 2, 3, 4)     // sum = 10
    add(5, 6, 7, 2, 3)  // sum = 23
	add([]int{1, 2, 3}...)
	add([]int{1, 2, 3, 4}...)
	add([]int{5, 6, 7, 2, 3}...)


	func add(x, y int, firstName, lastName string) (int , string) {
		var z int = x + y
		var fullName = firstName + " " + lastName
		return z, fullName
	}
	var age, name = add(4, 5, "Tom", "Simpson")
*/

/*
	Anonymous functions

	f := func(x, y int) int{ return x + y}
    fmt.Println(f(3, 4))        // 7


	func action(n1 int, n2 int, operation func(int, int) int){
		result := operation(n1, n2)
		fmt.Println(result)
	}
    action(10, 25, func (x int, y int) int { return x + y })    // 35
*/
