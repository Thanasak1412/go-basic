// Go program to illustrate the
// concept of multiple interfaces
package main

import "fmt"

// AuthorDetails Interface 1
type AuthorDetails interface {
	details()
	printAuthorName()
}

// AuthorArticles Interface 2
type AuthorArticles interface {
	articles()
	printAuthorName()
}

// Structure
type author struct {
	aName     string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tArticles int
}

func (a author) printAuthorName() {
	//TODO implement me
	fmt.Println("Current Salary is:", a.salary)
}

// Implementing method
// of the interface 1
func (a author) details() {
	fmt.Printf("Author Name: %s\n", a.aName)
	fmt.Printf("Branch: %s and passing year: %d\n", a.branch, a.year)
	fmt.Printf("College Name: %s\n", a.college)
	fmt.Printf("Salary: %d\n", a.salary)
	fmt.Printf("Published articles: %d\n", a.particles)
}

// Implementing method
// of the interface 2
func (a author) articles() {
	pendingArticles := a.tArticles - a.particles
	fmt.Printf("Pending articles: %d:\n", pendingArticles)
}

func main() {
	// Assigning values
	// to the structure
	values := author{
		aName:     "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tArticles: 309,
	}

	// Accessing the method
	// of the interface 1
	var i1 AuthorDetails = values
	i1.details()
	i1.printAuthorName()

	// Accessing the method
	// of the interface 2
	var i2 AuthorArticles = values
	i2.articles()
}
