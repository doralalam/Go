/*  Hardcode without shortcuts

//main package indicates go that this is the start of executable file
package main
// fmt and math are some of the available packages in go
import (
	"fmt"
	"math"
)
//main function indicates to start execution of code from here
// main function execution is invoked by go itself
func main() {
	var investment float64
	var interestRate float64
	var tenure float64
	fmt.Println("Investment: ")
	fmt.Scan(&investment)
	fmt.Println("Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Println("Tenure: ")
	fmt.Scan(&tenure)
	futureValue := investment*math.Pow(1+interestRate/100,tenure)
	fmt.Println("Future value of the investment is ",futureValue)
}


*/


//Short Code can be

/*


package main
import (
	"fmt"
	"math"
)
func main(){
	var investment, interestRate, tenure float64
	const inflation = 2.5
	fmt.Println("Investment: ")
	fmt.Scan(&investment)
	fmt.Println("Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Println("Tenure: ")
	fmt.Scan(&tenure)
	futureValue := investment * math.Pow(1+interestRate/100, tenure)
	//variable can be declared using colon instead of declaring first and then initiating
	futureRealValue := futureValue / math.Pow(1+inflation/100, tenure)
	fmt.Printf("Future Value is %v and the Future Real Value (adjusted to inflation) is %v", futureValue, futureRealValue)
	//%v is a place holder just as in c language
}

*/

// This code can also be written as 

package main
import (
	"fmt"
	"math"
)
func main(){
	var investment, interestRate, tenure float64
	const inflation = 2.5
	fmt.Println("Investment: ")
	fmt.Scan(&investment)
	fmt.Println("Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Println("Tenure: ")
	fmt.Scan(&tenure)
	futureValue := investment * math.Pow(1+interestRate/100, tenure)
	futureRealValue := futureValue / math.Pow(1+inflation/100, tenure)
	formattedString := fmt.Sprintf("Future Value is %.1f \nFuture Real Value (adjusted to inflation) is %.2f", futureValue, futureRealValue)
	// using place holders, we can limit the decimal places to output
	//Formatted string can be stored in other string variable
	fmt.Println(formattedString)
}
