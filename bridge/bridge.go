package bridge

import "fmt"

// 抽象部分：定义高级接口
type Implementor interface {
	Operation() string
}

// 实现部分：提供具体实现
type ConcreteImplementorA struct{}

func (c *ConcreteImplementorA) Operation() string {
	return "ConcreteImplementorA operation"
}

type ConcreteImplementorB struct{}

func (c *ConcreteImplementorB) Operation() string {
	return "ConcreteImplementorB operation"
}

// 抽象部分：使用高级接口
type Abstraction struct {
	implementor Implementor
}

func (a *Abstraction) SetImplementor(implementor Implementor) {
	a.implementor = implementor
}

func (a *Abstraction) Operation() {
	fmt.Println(a.implementor.Operation())
}
