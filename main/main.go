package main

import (
	"abstract_factory"
	"adapter"
	"bridge"
	bdr "builder"
	"composite"
	"factory_method"
	"fmt"
	ptt "prototype"
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

	/* ----------------------End------------------------- */

	/* ---------------Strutural_Patterns------------------ */

	/* ---------------adapter------------------ */
	// 创建一个Adaptee实例
	adaptee := &adapter.Adaptee{Data: "Hello, Adaptee!"}

	// 创建一个Adapter实例
	target := adapter.NewAdapter(adaptee)

	// 客户端调用Target接口的Request方法
	fmt.Println(target.Request())

	/* ---------------bridge------------------ */
	// 创建抽象部分对象
	abstraction := &bridge.Abstraction{}

	// 设置实现部分对象 A
	abstraction.SetImplementor(&bridge.ConcreteImplementorA{})
	abstraction.Operation() // 输出 "ConcreteImplementorA operation"

	// 设置实现部分对象 B
	abstraction.SetImplementor(&bridge.ConcreteImplementorB{})
	abstraction.Operation() // 输出 "ConcreteImplementorB operation"

	/* ---------------composite------------------ */
	// 创建复合组件
	root := composite.NewComposite("Root")

	// 创建叶子组件并添加到复合组件中
	leafA := composite.NewLeaf("Leaf A")
	leafB := composite.NewLeaf("Leaf B")
	root.Add(leafA)
	root.Add(leafB)

	// 创建另一个复合组件并添加到根复合组件中
	compositeX := composite.NewComposite("Composite X")
	leafX1 := composite.NewLeaf("Leaf X1")
	leafX2 := composite.NewLeaf("Leaf X2")
	compositeX.Add(leafX1)
	compositeX.Add(leafX2)
	root.Add(compositeX)

	// 执行根复合组件的操作
	root.Operation()

	// 移除一个叶子组件
	root.Remove(leafB)

	// 再次执行根复合组件的操作
	root.Operation()

	// 获取并操作一个子组件
	child := root.GetChild(1)
	if child != nil {
		child.Operation()
	}
	/* ---------------decorator------------------ */
	/* ---------------facade------------------ */
	/* ---------------flyweight------------------ */
	/* ---------------proxy------------------ */

	/* ----------------------End------------------------- */

	/* ---------------Behavioral Patterns------------------ */

	/* ---------------template_method------------------ */
	/* ---------------strategy------------------ */
	/* ---------------command------------------ */
	/* ---------------chain_of_responsibility------------------ */
	/* ---------------state------------------ */
	/* ---------------observer------------------ */
	/* ---------------memento------------------ */
	/* ---------------interpreter------------------ */
	/* ---------------iterator------------------ */
	/* ---------------visitor------------------ */
	/* ---------------mediator------------------ */

	/* ----------------------End------------------------- */

}
