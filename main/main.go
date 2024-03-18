package main

import (
	"abstract_factory"
	"adapter"
	bdr "builder"
	"fmt"
	ptt "prototype"

	"factory_method"
	"singleton"
)

func main() {
	/* ---------------Creational_Patterns------------------ */

	/* ---------------singleton------------------ */
	singleton1 := singleton.GetInstance1()
	singleton2 := singleton.GetInstance2()
	singleton3 := singleton.SingletonFactory()

	fmt.Println(singleton1, singleton2, singleton3)

	/* -----------------factory_method---------------- */
	factoryA := &factory_method.ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println(productA.Use()) // 输出：Product A is being used

	// 创建产品B的工厂
	factoryB := &factory_method.ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println(productB.Use()) // 输出：Product B is being used

	/* -----------------abstract_factory---------------- */
	factory1 := &abstract_factory.ConcreteFactory1{}
	factory2 := &abstract_factory.ConcreteFactory2{}

	// 使用抽象工厂创建产品
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()
	fmt.Println(productA1.Use())
	fmt.Println(productB1.Use())

	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()
	fmt.Println(productA2.Use())
	fmt.Println(productB2.Use())

	/* -----------------prototype---------------- */
	// 创建原型对象
	prototype := ptt.NewConcretePrototype("Hello, Prototype!")

	// 复制原型对象
	clonedPrototype := prototype.Clone()

	// 显示原型对象和复制对象的状态
	prototype.Display()
	clonedPrototype.Display()

	// 修改原型对象的状态
	prototype.Value = "Modified Prototype"

	// 再次显示原型对象和复制对象的状态，验证它们是否独立
	fmt.Println("Modified ConcretePrototype:", prototype.Value)
	clonedPrototype.Display()

	/* -----------------builder---------------- */
	// 创建 ConcreteBuilder 实例
	builder := bdr.NewConcreteBuilder()

	// 创建 Director 实例，并将 ConcreteBuilder 传入
	director := bdr.NewDirector(builder)

	// 指挥构建 Car 对象
	car := director.ConstructCar()

	// 输出构建好的 Car 对象
	fmt.Printf("Brand: %s, Engine: %s, Color: %s\n", car.Brand, car.Engine, car.Color)

	/* ---------------Creational_Patterns------------------ */

	/* ---------------Strutural_Patterns------------------ */

	/* ---------------adapter------------------ */
	// 创建一个Adaptee实例
	adaptee := &adapter.Adaptee{Data: "Hello, Adaptee!"}

	// 创建一个Adapter实例
	target := adapter.NewAdapter(adaptee)

	// 客户端调用Target接口的Request方法
	fmt.Println(target.Request())

	/* ---------------Strutural_Patterns------------------ */

	/* ---------------Behavioral Patterns------------------ */

	/* ---------------Behavioral Patterns------------------ */

}
