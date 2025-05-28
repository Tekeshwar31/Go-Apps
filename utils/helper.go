package utils

import "fmt"

// Exported function (starts with capital letter)
func Greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Can have multiple functions in the same package
func Add(a, b int) int {
    return a + b
}