package main

import "fmt"

func main() {
	intValue := 10
	typeSwitch(intValue) // same function has been called with different type of parameter
	floatValue := 10.0
	typeSwitch(floatValue) //  same function has been called with different type of parameter
	stringValue := "Hello"
	typeSwitch(stringValue) //  same function has been called with different type of parameter
}

// This function alone will accept any type of data as input and gives output respectively

func typeSwitch(value interface{}) { // This is a special interface type that works for any value
	// func typeSwitch(value any){} -- The above function can also be declared like this. Both means the same
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)

	}
}

/* the above functionality can also be achieved by

func typeSwitch(value interface{}) {
	var1, bool := value.(int) // value.(int) returns value and boolean value if it is int
	if bool { // this condition checks if the bool is true or not
		fmt.Println("Integer: ", var1)
		return
	}
	var2, bool := value.(float64)
	if bool {
		fmt.Println("Float64: ", var2)
	}
	var3, bool := value.(string)
	if bool {
		fmt.Println("String: ", var3)
	}
}

*/
