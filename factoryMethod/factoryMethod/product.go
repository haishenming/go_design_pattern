package factorymethod

import "go_design_pattern/factoryMethod/factoryMethod/factory"

//Product1 生产
func Product1() {
	f := factory.Factory1{}
	c := f.Product()
	c.SetMaterial("123")
	c.Result()
}

//Product2 生产
func Product2() {
	f := factory.Factory2{}
	c := f.Product()
	c.SetMaterial("456")
	c.Result()
}
