package basics

/*
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 1

	fmt.Println(cap(intCh))
	fmt.Println(len(intCh))

	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
*/

/*
	intChannelTest := make(chan map[string]int)

	go func(){
		fmt.Println("Go routine starts")
		intChannelTest <- map[string]int {"test": 6, "test2": 12}
	}()

	fmt.Println((<- intChannelTest)["test"])
*/

/*
	fmt.Println(<- createChan(5))

	func createChan(n int) chan int {
		ch := make(chan int)
		go func() {
			ch <- n
		}()
		return ch
	}
*/

/*
	intCh := make(chan int, 3)
    intCh <- 10
    intCh <- 3
    close(intCh)    // канал закрыт

    for i := 0; i < cap(intCh); i++ {
         if val, opened := <-intCh; opened {
            fmt.Println(val)
         } else {
            fmt.Println("Channel closed!")
         }
   }
*/
