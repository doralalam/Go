package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	hobbies := [3]string{"hobby1", "hobby2", "hobby3"}
	fmt.Println("Hobbies array:", hobbies) // task 1 completed

	fmt.Println("First element in hobbies:", hobbies[0])
	fmt.Println("2nd and 3rd elemenets in hobbies:", hobbies[1:3]) // task 2 completed

	slicedHobbies := hobbies[0:2]
	//slicedHobbies = hobbies[:2]
	fmt.Println("1st and 2nd elements:", slicedHobbies) // task 3 completed

	updatedSlicedHobbies := append(slicedHobbies[1:], hobbies[2])
	fmt.Println("Resliced for 2nd and last element from original array:", updatedSlicedHobbies) // task 4 completed

	courseGoals := []string{"goal1", "goal2"}
	fmt.Println("Dynamic array of course goals:", courseGoals) // task 5 completed

	courseGoals[1] = "new_goal"
	newCourseGoals := append(courseGoals, "goal2")
	fmt.Println("Appended course goals: ", newCourseGoals) // task 6 completed

	productsList := []Product{{title: "Book", id: "101", price: 150.50}, {title: "Pen", id: "105", price: 55.0}}
	// dynamic array of Product type
	fmt.Println("Dynamic list of producst:", productsList)
	newProductsList := append(productsList, Product{title: "Bag", id: "201", price: 500.0})
	fmt.Println("Appended dynamic list of products:", newProductsList) // task 7 completed

}
