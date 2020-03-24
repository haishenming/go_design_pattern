package builder

import "fmt"

type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	Result() string
}

type ConcreteBuilder struct {
	result string
}

func (c ConcreteBuilder) Result() string {
	return c.result
}

func (c *ConcreteBuilder) BuildPartA() {
	c.result += "A"
	fmt.Println(c.result)
}

func (c *ConcreteBuilder) BuildPartB() {
	c.result += "B"
	fmt.Println(c.result)
}

func (c *ConcreteBuilder) BuildPartC() {
	c.result += "C"
	fmt.Println(c.result)
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}
