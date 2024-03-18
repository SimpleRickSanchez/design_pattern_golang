package builder

// Car 是最终要构建的对象
type Car struct {
	Brand  string
	Engine string
	Color  string
}

// Builder 定义构建 Car 的接口
type Builder interface {
	SetBrand(brand string)
	SetEngine(engine string)
	SetColor(color string)
	GetCar() *Car
}

// ConcreteBuilder 实现了 Builder 接口
type ConcreteBuilder struct {
	car *Car
}

// NewConcreteBuilder 创建一个新的 ConcreteBuilder 实例
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{
		car: &Car{},
	}
}

// SetBrand 设置 Brand
func (b *ConcreteBuilder) SetBrand(brand string) {
	b.car.Brand = brand
}

// SetEngine 设置 Engine
func (b *ConcreteBuilder) SetEngine(engine string) {
	b.car.Engine = engine
}

// SetColor 设置 Color
func (b *ConcreteBuilder) SetColor(color string) {
	b.car.Color = color
}

// GetCar 返回构建好的 Car 对象
func (b *ConcreteBuilder) GetCar() *Car {
	return b.car
}

// Director 负责指挥构建过程
type Director struct {
	builder Builder
}

// NewDirector 创建一个新的 Director 实例
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// ConstructCar 指挥构建 Car 对象
func (d *Director) ConstructCar() *Car {
	d.builder.SetBrand("Toyota")
	d.builder.SetEngine("V6")
	d.builder.SetColor("Red")
	return d.builder.GetCar()
}
