// This code doesn't improve any performance.
// But it is very developer friendly to use aliases for better readability

package main

import "fmt"

// here we can make an alias name to the map
type floatMap map[string]string

// created a method just for fun
// we can even display the output directly without this method
func (m floatMap) output() {
	fmt.Println("Websites from method:", m)
}

func main() {
	// now we can use an alias name to map
	websites := make(floatMap, 2)
	websites["google"] = "www.google.com"
	websites["youtube"] = "www.youtube.com"
	// printing output directly
	fmt.Println("Websites:", websites)
	websites.output()

}
