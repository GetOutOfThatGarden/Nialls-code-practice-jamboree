package main

import "fmt"

// In this code I'm trying to see if it's possible to use pointers and addresses to access..
// ..variables from outside of the function that they're declared in.

// In sideVars() I declare two variables, x and y. x is an int and y is a pointer to x. GHC

func sideVars() {

	x := 5
	y := &x

}

// In main() I'm printing the value of variable y from sideVars(), which is the content of x (5).
// It would appear that even though I'm trying to print *y, it's still not visible outside of the function.
func main() {

	fmt.Println(*y)

}

// The only way I can see around this problem is by returning a variable from a function. NH
