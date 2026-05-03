package main

import (
	"errors"
	"fmt"
)

func closures() func() int {
	i := 0 // This variable is "captured"

	// We return the function itself, not the result
	return func() int {
		i++ // It remembers 'i' from the outside
		return i
	}
}

// func NewSuffixer(suffix string) func(string) string {
//     return func(name string) string {
//         return name + "." + suffix // captures 'suffix'
//     }
// }



func functionsDemo() {

	// 4. Now we call the wrapped version.
	// wrappedVersion()

	//     dotCom := NewSuffixer("com")
	// dotOrg := NewSuffixer("org")

	// fmt.Println(dotCom("google")) // google.com
	// fmt.Println(dotOrg("kubernetes")) // kubernetes.org

	//anonaious functionsDemo
	//mostly used in goroutine usage
	func(name string) {
		fmt.Printf("hi this is %v\n", name)

	}("Mohd Kamaal")

	nextInt := closures() // 'nextInt' is now a function that carries 'i' with it

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	a, b := 3, 6

	add(a, b) //passing arguements
	m, s := multiply()
	result := sub(a, b)
	newresult := division(a, b)
	fmt.Println(m, s, result, newresult)

	interchange(&a, &b)
	fmt.Printf("value of a is %v, and b is %v\n", a, b)
	//& used to give the address of var //ADDRESS OF
	// * used to update the address value // VALUE AT
	printNumbers(1, 2, 3, 4, 5, 6, 6)
	_, err := catchingError(a)
	if err != nil {
		fmt.Println("error is ", err)

	}

}

func add(a int, b int) { //receving is parameters

	fmt.Println(a + b)

}
func printNumbers(a ...int) { //variadic function
	fmt.Println(a)

}
func sub(a, b int) int {
	return b - a
}

func division(a, b int) (i int) { //named return means u give name to var already in reutnr so need to initialize first
	i = a / b
	return
}

func multiply() (int, string) {
	return 4 * 3, "multiply"

}
func interchange(a, b *int) {
	temp := 0
	temp = *a
	*a = *b
	*b = temp
	fmt.Println(*a, *b)

}

type catchErr string

// ideomatic representaion of function in go
// to use method of interface then it must be linked with the type
func catchingError(a int) (bool, error) {
	if a > 10 {
		return false, errors.New("failing here") //create own error from screath
	}
	return true, catchErr("has no eror") //showing usage of interface method by implementing with the new type and linking them up

}

func (c catchErr) Error() string {
	return fmt.Sprintf("always something breaking %v", string(c))

}
