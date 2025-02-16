# Functions

->  Function is a piece of code that can run on demand, which can take input as a parameter and can return multiple values if needed
->  Functions are 1st class values
        1. Functions as 1st class values in Go allows us to pass one function as a paramater to another function whenever there is a scenario to execute certain function multiple times over the dependency of another function.
        For example:
            Go through the functions_as_values.go file in functions_deep_dive folder


# Anonymous Functions

->  Anonymous functions are the functions without a function name and won't be stored in program like a regular function
->  Anonymous functions are used in the scenarios where a function is required to run only once
->  Anonymous functions are used whenever a function is accepting another function as parameter value or whenever a function is returning another function


# Recursion

->  A recursion is said when a function calls itself

# Variadic functions

->  Whenever a function has been passed with n number of parameters i.e., we mention ellipsis (...) inside the acceptable parameter list, then we call that function as a variadic function.
->  Variadic functions are useful whenever we don't know the number of arguments we need to accept from the user
    Example:    func sumup(numbers ...int) int{
                    sum :=0
                    for _,val := range numbers{
                        sum += numbers
                    }
                }
->  We can even split the slices into individual parameters using ellipses (...)