package basics

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func repoMethods(ch chan bool) {
	// run query sql
	time.Sleep(time.Second * 3)
	fmt.Println("run query sql")
	ch <- false
}

func saveToDB(ctx context.Context, cancel context.CancelFunc) error {
	var err error
	ch := make(chan bool, 1)

	go repoMethods(ch)

	select {
	case <-ctx.Done():
		fmt.Println("cancel saveToDB")
	case res := <-ch:
		if !res {
			err = errors.New("error saveToDB")
			fmt.Println("error saveToDB")
			cancel()
		}
	}
	return err
}

func spaceMethods(ch chan bool) {
	// run spaceMethods
	time.Sleep(time.Second * 5)
	fmt.Println("run spaceMethods")
	ch <- false
}

func saveToSpace(ctx context.Context, cancel context.CancelFunc) error {
	var err error
	ch := make(chan bool, 1)

	go spaceMethods(ch)

	select {
	case <-ctx.Done():
		fmt.Println("cancel saveToSpace")
	case res := <-ch:
		if !res {
			err = errors.New("error saveToSpace")
			fmt.Println("error saveToSpace")
			cancel()
		}
	}
	return err
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var err error
	ctx, cancel := context.WithCancel(context.Background())

	go func(err *error) {
		defer wg.Done()
		*err = saveToDB(ctx, cancel)
	}(&err)

	go func(err *error) {
		defer wg.Done()
		*err = saveToSpace(ctx, cancel)
	}(&err)

	wg.Wait()

	if err != nil {
		fmt.Println("Error after finish: ", err.Error())
	}
}
