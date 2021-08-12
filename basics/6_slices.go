package basics

/*
	Slice can be initialized:
	var users []string
	var users = []string{"Tom", "Alice", "Kate"}
	users2 := []string{"Tom", "Alice", "Kate"}


	Add to slice:
	//append(slice, value)

	users := []string{"Tom", "Alice", "Kate"}
	users = append(users, "Bob")


	Removing an item:
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]


	Slice operator:
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6]
	users2 := initialUsers[:4]
	users3 := initialUsers[3:]


	Example 1:

	var users []string = []string{"Tom", "Alice", "Kate"}
	fmt.Println(users[2])   // Kate
	users[2] = "Katherine"

	for _, value := range users{
		fmt.Println(value)
	}


	Example 2:

	var users []string = make([]string, 3)
	users[0] = "Tom"
	users[1] = "Alice"
	users[2] = "Bob"
*/
