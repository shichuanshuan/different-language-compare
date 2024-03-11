package main

import "fmt"

type Person struct {
	name string
	age  int
}

func test200(p1 *Person) {
	p1.age = 33
}

func test300(p1 *Person) {
	p1 = nil
}

func test400(p1 *Person) {
	p1 = new(Person)
}

func main() {
	var p *Person
	p = new(Person)
	p.age = 3

	test200(p)
	fmt.Println("1. age= ", p.age)

	test300(p)
	fmt.Println("2. age= ", p.age)

	test400(p)
	fmt.Println("3. age= ", p.age)
}
