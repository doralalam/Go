package main
import (
	"fmt"
	"math"
)
const inflation = 2.5	 // const variable outside the functions to get global scope
func main(){
	var investment, interestRate, tenure float64
	fmt.Println("Investment: ")
	fmt.Scan(&investment)
	fmt.Println("Interest Rate: ")
	fmt.Scan(&interestRate)
	fmt.Println("Tenure: ")
	fmt.Scan(&tenure)
	// calling function and assigning the 2 return values to 2 different local variables
	futureValue, futureRealValue := cal_fut_values(investment, interestRate, tenure)
	formattedString := fmt.Sprintf("Future Value is %.1f \nFuture Real Value (adjusted to inflation) is %.2f", futureValue, futureRealValue)
	fmt.Println(formattedString)
}
	// func, function_name, parameters, return type. Since we are returning 2 values, we must mention 2 return type in parenthesis
	func cal_fut_values(investment, interestRate, tenure float64) (float64, float64){
	fv := investment * math.Pow(1+interestRate/100, tenure)
	// : creates the variable and = assigns values to the variables
	frv := fv / math.Pow(1+inflation/100, tenure)
	return fv, frv
}

// the function can also be written like this 
/*
	func cal_fut_values(investment, interestRate, tenure float64) (fv float64, frv float64){
	fv := investment * math.Pow(1+interestRate/100, tenure)
	frv := fv / math.Pow(1+inflation/100, tenure)
	return
	// instead of mentioning the return values explicitly, we mentioned them along with their return types on top
}	
*/