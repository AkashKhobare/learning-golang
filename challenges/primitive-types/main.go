package main

import "fmt"

// package level vars
var val = "global var"

func main() {
	// local var
	val := 250
	fmt.Printf("Local val: %v\n", val)

	// get global var
	fmt.Printf("Global val: %v\n", getGlobalVal())

	// update global var and print it
	updateGlobalVal("updated global var")
	fmt.Printf("Updated Global val: %v\n", getGlobalVal())
	
	// create pointer to local var
	var pVal *int = &val

	// print address of local var
	fmt.Printf("Local val address: %v\n", pVal)

	// update local var value using pointer
	*pVal = 350

	// print updated local var
	fmt.Printf("Updated local val: %v\n", val)
}

func getGlobalVal() string {
	return val
}

func updateGlobalVal(newVal string) {
	val = newVal
}