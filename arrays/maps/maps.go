// to demonstrate maps usage

package main

import "fmt"

func main() {

	// To create a map, we must enter type of key in [] and type of value outside []
	// here map[string] indicates, keys (known as labels when associated with maps) are of string type
	// map[string]string indicates, keys are strings followed by values are also strings
	websites := map[string]string{"Google": "www.google.com", "AWS": "www.aws.com"}
	fmt.Println("Websites:", websites)
	// to access values using keys
	fmt.Println(websites["Google"])
	// to assign a new key-value to map
	websites["Linkedin"] = "www.linkedin.com"
	fmt.Println(websites)
	// to delete a key value pair, mention map name followed by key to delete
	delete(websites, "Google")
	fmt.Println(websites)

}

/* 	the advantage of using maps over arrays in this scenarios is that,
it will be easy to remember key names instead of index values or positions in arrays
*/
