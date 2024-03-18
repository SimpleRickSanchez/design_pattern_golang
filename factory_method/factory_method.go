package factory_method

// Product 定义产品的接口
type Product interface {
	Use() string
}

// ConcreteProductA 是具体的产品A
type ConcreteProductA struct {
	Data int
}

func (p *ConcreteProductA) Use() string {
	return "Product A is being used"
}

// ConcreteProductB 是具体的产品B
type ConcreteProductB struct {
	Data int
}

func (p *ConcreteProductB) Use() string {
	return "Product B is being used"
}

// Factory 定义创建产品的工厂接口
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA 是创建产品A的工厂
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB 是创建产品B的工厂
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}
