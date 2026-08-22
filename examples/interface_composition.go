package main

import (
	"fmt"
	"math"
)

// =============================================================================
// INTERFACES AND COMPOSITION - Learning Go OOP
// =============================================================================
//
// Go doesn't have classes or inheritance. Instead it uses:
// - INTERFACES: Define behavior (what something can do)
// - COMPOSITION: Build complex types by embedding simpler ones
//
// This is often summarized as: "Favor composition over inheritance"
// =============================================================================

// =============================================================================
// PART 1: INTERFACES
// =============================================================================

// Shape is an interface - it defines BEHAVIOR, not data
// Any type that has these methods automatically implements Shape
// No "implements" keyword needed - this is called "implicit implementation"
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Printer is another interface - single method interfaces are common in Go
type Printer interface {
	Print()
}

// Rectangle is a struct (data)
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle implements Shape interface by having Area() and Perimeter() methods
// No explicit declaration needed - if it has the methods, it implements the interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Print() {
	fmt.Printf("Rectangle: %.2f x %.2f\n", r.Width, r.Height)
}

// Circle is another struct that implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Print() {
	fmt.Printf("Circle: radius %.2f\n", c.Radius)
}

// Triangle also implements Shape
type Triangle struct {
	A, B, C float64 // sides
}

func (t Triangle) Area() float64 {
	// Heron's formula
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) Print() {
	fmt.Printf("Triangle: sides %.2f, %.2f, %.2f\n", t.A, t.B, t.C)
}

// PrintShapeInfo accepts ANY type that implements Shape interface
// This is polymorphism - one function works with multiple types
func PrintShapeInfo(s Shape) {
	fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// =============================================================================
// PART 2: COMPOSITION (Embedding)
// =============================================================================

// Address is a simple struct
type Address struct {
	Street  string
	City    string
	Country string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.Country)
}

// Person has an embedded Address
// This is COMPOSITION - Person "has a" Address
type Person struct {
	Name string
	Age  int
	Address // Embedded struct (no field name = composition)
}

// Employee embeds Person - composition can be nested
type Employee struct {
	Person   // Embedded - gets all Person fields and methods
	JobTitle string
	Salary   float64
}

// Company shows composition with multiple embedded types
type Company struct {
	Name      string
	Employees []Employee
	Address   // Embedded
}

func (c Company) TotalSalary() float64 {
	total := 0.0
	for _, e := range c.Employees {
		total += e.Salary
	}
	return total
}

// =============================================================================
// PART 3: INTERFACE COMPOSITION
// =============================================================================

// Interfaces can also be composed from other interfaces

// Reader interface - can read data
type Reader interface {
	Read() string
}

// Writer interface - can write data
type Writer interface {
	Write(data string)
}

// ReadWriter combines Reader and Writer interfaces
// Any type implementing ReadWriter must have BOTH Read() and Write()
type ReadWriter interface {
	Reader
	Writer
}

// File implements ReadWriter interface
type File struct {
	Name    string
	Content string
}

func (f *File) Read() string {
	return f.Content
}

func (f *File) Write(data string) {
	f.Content = data
}

// =============================================================================
// PART 4: EMPTY INTERFACE
// =============================================================================

// interface{} or "any" (Go 1.18+) accepts ANY type
// Useful for generic containers, but lose type safety
func PrintAnything(value any) {
	// Type switch to handle different types
	switch v := value.(type) {
	case int:
		fmt.Printf("  Integer: %d\n", v)
	case string:
		fmt.Printf("  String: %s\n", v)
	case Shape:
		fmt.Printf("  Shape with area: %.2f\n", v.Area())
	default:
		fmt.Printf("  Unknown type: %v\n", v)
	}
}

// =============================================================================
// MAIN - Demonstrations
// =============================================================================

func main() {
	fmt.Println("=== PART 1: INTERFACES ===")
	fmt.Println()

	// Create different shapes
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	triangle := Triangle{A: 3, B: 4, C: 5}

	// All shapes can be stored in a slice of Shape interface
	shapes := []Shape{rect, circle, triangle}

	for _, shape := range shapes {
		// Type assertion to get the concrete type for Print()
		if p, ok := shape.(Printer); ok {
			p.Print()
		}
		PrintShapeInfo(shape)
		fmt.Println()
	}

	fmt.Println("=== PART 2: COMPOSITION ===")
	fmt.Println()

	// Create an employee using composition
	emp := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
			Address: Address{
				Street:  "123 Main St",
				City:    "New York",
				Country: "USA",
			},
		},
		JobTitle: "Software Engineer",
		Salary:   100000,
	}

	// Access embedded fields directly (no need for emp.Person.Name)
	fmt.Printf("Name: %s\n", emp.Name)                // Direct access
	fmt.Printf("City: %s\n", emp.City)                // From embedded Address
	fmt.Printf("Full Address: %s\n", emp.FullAddress()) // Method from embedded Address
	fmt.Printf("Job: %s\n", emp.JobTitle)
	fmt.Println()

	// Company with multiple employees
	company := Company{
		Name: "Tech Corp",
		Address: Address{
			Street:  "456 Tech Ave",
			City:    "San Francisco",
			Country: "USA",
		},
		Employees: []Employee{
			emp,
			{
				Person:   Person{Name: "Jane Smith", Age: 28},
				JobTitle: "Product Manager",
				Salary:   120000,
			},
		},
	}

	fmt.Printf("Company: %s\n", company.Name)
	fmt.Printf("Location: %s\n", company.FullAddress())
	fmt.Printf("Total Salaries: $%.2f\n", company.TotalSalary())
	fmt.Println()

	fmt.Println("=== PART 3: INTERFACE COMPOSITION ===")
	fmt.Println()

	file := &File{Name: "example.txt"}
	file.Write("Hello, World!")
	fmt.Printf("File: %s\n", file.Name)
	fmt.Printf("Content: %s\n", file.Read())
	fmt.Println()

	fmt.Println("=== PART 4: EMPTY INTERFACE ===")
	fmt.Println()

	PrintAnything(42)
	PrintAnything("Hello")
	PrintAnything(circle)
	PrintAnything([]int{1, 2, 3})
}
