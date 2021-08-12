package basics

/*
	var numbers [count]type

	var numbers [5]int

	var numbers [5]int = [5]int{1,2,3,4,5}

	var numbers [5]int = [5]int{1,2} // [1 2 0 0 0]

	var numbers [1000000] int
	start := time.Now()
	for index := range numbers {
		numbers[index] = rand.Intn(10000)
	}
	fmt.Println(time.Since(start).Seconds())

*/
