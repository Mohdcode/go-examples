package main

import (
	"fmt"
)

type vehicle struct {
	wheel int
	name  string
	user  string
}

type car struct {
	vehicle
}
type bike struct {
	vehicle
}

func (c vehicle) honk() {
	fmt.Printf("%v is honking\n", c.name)
}
func (c *vehicle) changeVehicle(name string) {
	c.name = name

}
func (v vehicle) Description(name string) {
	fmt.Printf("this is %v engine\n", name)

}

// func Inspect(e Engine) {
// 	e.Description()
// }

func (c vehicle) speak() {
	fmt.Printf("hi i am the %v\n", c.name)
}

type Engine interface {
	Description(string)
	speak()
}

func structs() {

	car1 := car{vehicle{wheel: 4, name: "maruti", user: "driver"}} //this is how embedding work in golang ccalled inheritance
	bike1 := bike{vehicle{2, "splendor", "rider"}}
	bike1.honk()
	bike1.changeVehicle("bullete 350")
	fmt.Printf("new vehicle name is %v\n", bike1.name)
	car1.Description("v8")
	bike1.Description("350cc")
	car1.speak()
	Engine.Description(car1, "v8")//one way to identifying the interface statisfaction
	var i interface{}=car1

	if e,ok:=i.(Engine);ok{
		fmt.Printf("this is %v\n",e)
	}
	// Engine.(car)
	// Inspect(car1)

	car1.honk()
	car1.changeVehicle("XUV")
	fmt.Printf("new vehicle name is %v\n", car1.name)

	a := struct {
		number int
		name   string
	}{1, "amma"} //making var of struct custom structure to keep the type in one var

	var b interface{} = "string" //can keep any type in it and any kind
	fmt.Println(a, b)

	c := []struct { //array of structs
		number int
		name   string
	}{
		{1, "bilal"},
		{2, "aarzoo"},
	}
	fmt.Println(c)

	// This is a flat slice: 1 dimension
	n := []interface{}{"string", 1234}

	val, ok := n[0].(string) //reserved for the interface only and some more
	if !ok {
		panic("BOOM")
	}
	fmt.Println(val)
}

// type catchErr string

// // ideomatic representaion of function in go
// // to use method of interface then it must be linked with the type
// func catchingError(a int) (bool, error) {
// 	if a > 10 {
// 		return false, errors.New("failing here") //create own error from screath
// 	}
// 	return true, catchErr("has no eror") //showing usage of interface method by implementing with the new type and linking them up

// }

// func (c catchErr) Error() string {
// 	return fmt.Sprintf("always something breaking %v", string(c))

// }
