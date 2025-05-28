// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// // ArrayAndJsonMethod demonstrates various array and JSON operations in Go
// func ArrayAndJsonMethod() {
// 	// 1. Array Operations
// 	fmt.Println("=== Array Operations ===")
	
// 	// Creating arrays
// 	numbers := [5]int{1, 2, 3, 4, 5}
// 	names := []string{"John", "Jane", "Bob"} // Slice (dynamic array)
	
// 	// Array iteration
// 	fmt.Println("\nIterating through array:")
// 	for i, num := range numbers {
// 		fmt.Printf("Index: %d, Value: %d\n", i, num)
// 	}
	
// 	// Array slicing
// 	fmt.Println("\nArray slicing:")
// 	slice := numbers[1:4]
// 	fmt.Println("Slice:", slice)
	
// 	// Array append
// 	fmt.Println("\nAppending to slice:")
// 	names = append(names, "Alice")
// 	fmt.Println("Updated names:", names)
	
// 	// Array length and capacity
// 	fmt.Println("\nArray length and capacity:")
// 	fmt.Printf("Length of numbers: %d\n", len(numbers))
// 	fmt.Printf("Length of names: %d\n", len(names))
	
// 	// 2. JSON Operations
// 	fmt.Println("\n=== JSON Operations ===")
	
// 	// Struct for JSON
// 	type Person struct {
// 		Name    string   `json:"name"`
// 		Age     int      `json:"age"`
// 		Hobbies []string `json:"hobbies"`
// 	}
	
// 	// Creating a struct instance
// 	person := Person{
// 		Name:    "John Doe",
// 		Age:     30,
// 		Hobbies: []string{"reading", "gaming", "coding"},
// 	}
	
// 	// Converting struct to JSON (marshaling)
// 	jsonData, err := json.Marshal(person)
// 	if err != nil {
// 		fmt.Println("Error marshaling:", err)
// 		return
// 	}
// 	fmt.Println("\nStruct to JSON:")
// 	fmt.Println(string(jsonData))
	
// 	// Converting JSON to struct (unmarshaling)
// 	jsonStr := `{"name":"Jane Doe","age":25,"hobbies":["painting","dancing"]}`
// 	var newPerson Person
// 	err = json.Unmarshal([]byte(jsonStr), &newPerson)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling:", err)
// 		return
// 	}
// 	fmt.Println("\nJSON to Struct:")
// 	fmt.Printf("Name: %s, Age: %d, Hobbies: %v\n", newPerson.Name, newPerson.Age, newPerson.Hobbies)
	
// 	// Working with JSON arrays
// 	jsonArray := `[{"name":"John","age":30},{"name":"Jane","age":25}]`
// 	var people []Person
// 	err = json.Unmarshal([]byte(jsonArray), &people)
// 	if err != nil {
// 		fmt.Println("Error unmarshaling array:", err)
// 		return
// 	}
// 	fmt.Println("\nJSON Array to Struct Slice:")
// 	for _, p := range people {
// 		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
// 	}
// }

// func main() {
// 	ArrayAndJsonMethod()
// } 