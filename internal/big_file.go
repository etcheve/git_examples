package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Person struct represents a person with a name and age.
type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new Person instance.
func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Greet prints a greeting message from the person.
func (p *Person) Greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.Name, p.Age)
}

// RandomNumbers generates a slice of n random integers.
func RandomNumbers(n int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, n)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

// PrintSlice prints a slice of integers.
func PrintSlice(numbers []int) {
	fmt.Println("Numbers in the slice:", numbers)
}

// Average calculates the average of a slice of integers.
func Average(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func main() {
	// Create and greet a person
	alice := NewPerson("Alice", 25)
	alice.Greet()

	// Generate random numbers
	nums := RandomNumbers(10)
	PrintSlice(nums)

	// Calculate and print the square of a number
	num := 7
	fmt.Printf("The square of %d is %d.\n", num, Square(num))

	// Calculate and print the average of the numbers
	avg := Average(nums)
	fmt.Printf("The average of the random numbers is %.2f.\n", avg)
}

