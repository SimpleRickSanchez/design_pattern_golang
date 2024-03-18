package abstract_factory

// AbstractProduct 定义抽象产品的接口
type AbstractProductA interface {
	Use() string
}

type AbstractProductB interface {
	Use() string
}

// ConcreteProductA 实现抽象产品A的接口
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

// ConcreteProductB 实现抽象产品B的接口
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

// AbstractFactory 定义抽象工厂的接口
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// ConcreteFactory1 实现抽象工厂接口，创建具体产品A1和具体产品B1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ConcreteProductA{}
}

func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ConcreteProductB{}
}

// ConcreteFactory2 实现抽象工厂接口，创建具体产品A2和具体产品B2（可能是不同的实现）
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	// 返回不同的具体产品A实现
	return &ConcreteProductA{}
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	// 返回不同的具体产品B实现
	return &ConcreteProductB{}
}
