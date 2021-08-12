package basics

/*
	for [init counter]; [condition]; [update count]{
		// action
	}


   	for i := 1; i < 10; i++{
        fmt.Println(i * i)
    }


	var i = 1
	for ; i < 10; i++{
		fmt.Println(i * i)
	}


	var i = 1
	for ; i < 10;{
		fmt.Println(i * i)
		i++
	}


	var i = 1
	for i < 10{
		fmt.Println(i * i)
		i++
	}
*/

/*
	Iterating over arrays

	for index, value := range array{
		// action
	}


	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users{
		fmt.Println(index, value)
	}


	for _, value := range users{
		fmt.Println(value)
	}


	var users = [3]string{"Tom", "Alice", "Kate"}
	for i:= 0; i < len(users); i++{
		fmt.Println(users[i])
	}
*/
