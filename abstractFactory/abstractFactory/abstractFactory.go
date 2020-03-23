package abstractfactory

import "fmt"

//AbstractFactory 抽象工厂
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

//AbstractProductA 抽象产品A
type AbstractProductA interface {
	UserA()
}

//AbstractProductB 抽象产品B
type AbstractProductB interface {
	UserB()
}

//ProductA1 具体产品A1
type ProductA1 struct {
}

func (p ProductA1) UserA() {
	fmt.Println("Use A")
}

//ProductB1 具体产品B2
type ProductB1 struct {
}

func (p ProductB1) UserB() {
	fmt.Println("Use B")
}

//Factory1 具体工厂1
type Factory1 struct {
	A AbstractProductA
	B AbstractProductB
}

func (f Factory1) CreateProductA() AbstractProductA {
	return ProductA1{}
}

func (f Factory1) CreateProductB() AbstractProductB {
	return ProductB1{}
}
