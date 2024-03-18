package template_method

import "fmt"

// 定义了一个模板方法的接口
type AbstractClass interface {
	PrimitiveOperation1()
	PrimitiveOperation2()
	TemplateMethod()
}

// 具体实现类，实现 AbstractClass 接口
type ConcreteClass struct {
	Data string
}

// 实现 PrimitiveOperation1 方法
func (c *ConcreteClass) PrimitiveOperation1() {
	fmt.Println("ConcreteClass implements PrimitiveOperation1: ", c.Data)
}

// 实现 PrimitiveOperation2 方法
func (c *ConcreteClass) PrimitiveOperation2() {
	fmt.Println("ConcreteClass implements PrimitiveOperation2: ", c.Data)
}

// 实现 TemplateMethod 方法，这是一个模板方法，它定义了算法的骨架
func (c *ConcreteClass) TemplateMethod() {
	fmt.Println("TemplateMethod starts executing...")
	c.PrimitiveOperation1()
	c.PrimitiveOperation2()
	fmt.Println("TemplateMethod ends executing...")
}
