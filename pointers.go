package main

import "fmt"

func pointers() {

	var a *int //only be used to keep address of var that is not pointer
	x := 3
	a = &x //a has no address now it can't put x at nay address it can only use others once it has own addres then it can be manipulated by defrencing
	*a = 4

	// var n *int: You have a business card that is completely blank. If you try to call the number on it, you can't.

	// n := new(int): You have a business card with an address written on it. When you go to that address, you find a house. The house is currently empty (it has the "zero value" of 0), but the house exists.
	fmt.Printf("this is the type %T and value %v\n", a, *a)

	p := new(4)
	fmt.Println(p)//will print the address cuz its a pointer with memory but zero sized

	k := 4
	l := 5
	swap(&k, &l) //sharing address as aggruemnts
	fmt.Printf("this is k %v and this is l %v\n", k, l)
	k, l = swapValue(k, l) //sharing address as aggruemnts
	fmt.Printf("this is k %v and this is l %v\n", k, l)

}
func swap(a, b *int) { //pointer creation as parameters
	//changing the address here and no need to return
	temp := *a
	*a = *b
	*b = temp

}

func swapValue(a, b int) (int, int) {
	//need return cuz it copies things
	temp := a
	a = b
	b = temp
	return a, b
}
