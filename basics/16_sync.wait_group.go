package basics

/*
	var wg sync.WaitGroup
	wg.Add(2)
	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Goroutines %d start \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Goroutines %d finish \n", id)
	}

	go work(1)
	go work(2)

	wg.Wait()
*/
