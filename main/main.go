package main

import (
	"abstract_factory"
	"adapter"
	"bridge"
	bdr "builder"
	cr "chain_of_responsibility"
	cmd "command"
	"composite"
	"decorator"
	"facade"
	"factory_method"
	"flyweight"
	"fmt"
	ptt "prototype"
	"proxy"
	"singleton"
	"state"
	"strategy"
	tmplmethod "template_method"
)

func main() {
	/* ---------------Creational_Patterns------------------ */

	fmt.Println("/* ---------------singleton------------------ */")
	singleton1 := singleton.GetInstance1() // lazy
	singleton2 := singleton.GetInstance2() // eager
	singleton3 := singleton.SingletonFactory()

	fmt.Println(singleton1, singleton2, singleton3)

	fmt.Println("/* -----------------factory_method---------------- */")
	factoryA := &factory_method.ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println(productA.Use()) // 输出：Product A is being used

	// 创建产品B的工厂
	factoryB := &factory_method.ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println(productB.Use()) // 输出：Product B is being used

	fmt.Println("/* -----------------abstract_factory---------------- */")
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

	fmt.Println("/* -----------------prototype---------------- */")
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

	fmt.Println("/* ---------------adapter------------------ */")
	// 创建一个Adaptee实例
	adaptee := &adapter.Adaptee{Data: "Hello, Adaptee!"}

	// 创建一个Adapter实例
	target := adapter.NewAdapter(adaptee)

	// 客户端调用Target接口的Request方法
	fmt.Println(target.Request())

	fmt.Println("/* ---------------bridge------------------ */")
	// 创建抽象部分对象
	abstraction := &bridge.Abstraction{}

	// 设置实现部分对象 A
	abstraction.SetImplementor(&bridge.ConcreteImplementorA{})
	abstraction.Operation() // 输出 "ConcreteImplementorA operation"

	// 设置实现部分对象 B
	abstraction.SetImplementor(&bridge.ConcreteImplementorB{})
	abstraction.Operation() // 输出 "ConcreteImplementorB operation"

	fmt.Println("/* ---------------composite------------------ */")
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

	fmt.Println("/* ---------------decorator------------------ */")
	// 创建一个 Circle 实例
	circle := &decorator.Circle{Radius: 5.0}

	// 使用 Decorator 包装 Circle 实例
	decoratedCircle := &decorator.Decorator{Shape: circle}

	// 使用 RedCircleDecorator 扩展 Decorator
	redCircle := &decorator.RedCircleDecorator{Decorator: *decoratedCircle}

	// 调用 RedCircleDecorator 的 Draw 方法
	fmt.Println(redCircle.Draw())

	fmt.Println("/* ---------------facade------------------ */")
	// 创建 Facade 实例
	facade := facade.NewFacade()

	// 客户端使用 Facade 的 Operation 方法
	fmt.Println(facade.Operation())

	fmt.Println("/* ---------------flyweight------------------ */")
	// 创建享元工厂
	factory := flyweight.FlyweightFactory{}

	// 客户端请求享元对象
	extrinsicState1 := "Client State 1"
	extrinsicState2 := "Client State 2"
	flyweight1 := factory.GetFlyweight("Shared State 1")
	flyweight2 := factory.GetFlyweight("Shared State 2")

	// 使用享元对象
	result1 := flyweight1.Operation(extrinsicState1)
	result2 := flyweight2.Operation(extrinsicState2)

	// 输出结果
	fmt.Println(result1)
	fmt.Println(result2)

	// 再次请求相同的享元对象，应该得到之前缓存的对象
	flyweight1Again := factory.GetFlyweight("Shared State 1")
	result1Again := flyweight1Again.Operation(extrinsicState1)

	// 输出结果，应该是相同的对象
	fmt.Println(result1Again)

	fmt.Println("/* ---------------proxy------------------ */")

	// 使用代理对象
	proxy := &proxy.ProxyImage{}
	proxy.Display() // 第一次调用时，会加载真实图片
	proxy.Display() // 第二次调用时，直接使用已加载的图片

	/* ----------------------End------------------------- */

	/* ---------------Behavioral Patterns------------------ */

	fmt.Println("/* ---------------template_method------------------ */")

	// 创建具体实现类的对象
	obj := &tmplmethod.ConcreteClass{Data: "Hello, Template Method Pattern!"}

	// 调用模板方法
	obj.TemplateMethod()

	fmt.Println("/* ---------------strategy------------------ */")
	// 创建上下文对象
	context := &strategy.Context{}

	// 设置策略 A
	context.SetStrategy(&strategy.ConcreteStrategyA{})
	fmt.Println(context.ExecuteStrategy()) // 输出: Executing strategy A

	// 更改策略为 B
	context.SetStrategy(&strategy.ConcreteStrategyB{})

	fmt.Println(context.ExecuteStrategy()) // 输出: Executing strategy B

	fmt.Println("/* ---------------command------------------ */")
	// 创建接收者
	receiver := &cmd.ConcreteReceiver{Name: "Test"}

	// 创建具体命令，并设置接收者
	command := &cmd.ConcreteCommand{Receiver: receiver}

	// 创建调用者，并设置命令
	invoker := &cmd.Invoker{}
	invoker.SetCommand(command)

	// 执行命令
	invoker.ExecuteCommand(receiver)

	fmt.Println("/* ---------------chain_of_responsibility------------------ */")

	// 创建处理者对象，并设置它们的后继处理者
	handlerA := &cr.ConcreteHandlerA{}
	handlerB := &cr.ConcreteHandlerB{}
	handlerA.SetSuccessor(handlerB)

	// 模拟客户端请求
	clientRequests := []interface{}{"requestA", "requestB", "requestC"}

	// 遍历请求，逐个处理
	for _, request := range clientRequests {
		handlerA.HandleRequest(request)
	}

	fmt.Println("/* ---------------state------------------ */")

	st_context := &state.Context{}
	st_context.SetState(&state.ConcreteStateA{}) // 初始状态设置为A

	st_context.Request() // 执行A状态的处理
	st_context.Request() // 执行B状态的处理，因为A状态处理中更改了状态

	fmt.Println("/* ---------------observer------------------ */")

	fmt.Println("/* ---------------memento------------------ */")

	fmt.Println("/* ---------------interpreter------------------ */")

	fmt.Println("/* ---------------iterator------------------ */")

	fmt.Println("/* ---------------visitor------------------ */")

	fmt.Println("/* ---------------mediator------------------ */")

	/* ----------------------End------------------------- */

}
