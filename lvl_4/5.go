package main

func main() {
	// & in this assignment points to the reference, not a value copy
	phil := gopher{name: "Phil", age: 35}
	validateAge()
}

type gopher struct {
	age     int
	name    string
	isAdult bool
}

// since we expect to work on a reference, * dereferences the pointer and
// allows us to work on the referenced object
func validateAge(g *gopher) {
	g.isAdult = g.age >= 21
}
