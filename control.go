package main

import "fmt"

func control() {

	//It only works inside a defer function: If you call it normally, it will always return nil and do nothing.

	//It must be in the same Goroutine: You cannot recover from a panic that happened in a different background thread.

	fmt.Println("now rintin")

	defer fmt.Println("getting out of function")

	for range 3 {
		fmt.Println("amma")
	}
	for i := 0; i <= 6; i++ {
		if i == 0 {
			fmt.Println("we have started")
		}
		switch i {
		case 3:
			fmt.Println("this is 3")
		case 5:
			tryPanic()
		case 1:
			tryPanic()
		default:
			fmt.Println("we are not in conditions")

		}
		fmt.Println(i)
	}

}
//panic can't be recover in main function and only be recovered in single routine oe main routine only
func tryPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering from panic")
		}

	}()
	panic("we are panic")
}
