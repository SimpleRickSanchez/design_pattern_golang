package mediator

import "fmt"

// Colleague 定义了一个同事接口
type Colleague interface {
	Send(message string)   // 发送消息给媒介者
	Notify(message string) // 从媒介者接收消息
}

// ConcreteColleagueA 是同事A的具体实现
type ConcreteColleagueA struct {
	mediator Mediator
	name     string
}

func NewConcreteColleagueA(mediator Mediator, name string) *ConcreteColleagueA {
	return &ConcreteColleagueA{mediator: mediator, name: name}
}

func (c *ConcreteColleagueA) Send(message string) {
	c.mediator.Send(c, message)
}

func (c *ConcreteColleagueA) Notify(message string) {
	fmt.Printf("Colleague A received message: %s\n", message)
}

// ConcreteColleagueB 是同事B的具体实现
type ConcreteColleagueB struct {
	mediator Mediator
	name     string
}

func NewConcreteColleagueB(mediator Mediator, name string) *ConcreteColleagueB {
	return &ConcreteColleagueB{mediator: mediator, name: name}
}

func (c *ConcreteColleagueB) Send(message string) {
	c.mediator.Send(c, message)
}

func (c *ConcreteColleagueB) Notify(message string) {
	fmt.Printf("Colleague B received message: %s\n", message)
}

// Mediator 定义了媒介者接口
type Mediator interface {
	Send(colleague Colleague, message string)
}

// ConcreteMediator 是媒介者的具体实现
type ConcreteMediator struct {
	colleagues map[Colleague]string
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: make(map[Colleague]string),
	}
}

func (m *ConcreteMediator) Register(colleague Colleague, name string) {
	m.colleagues[colleague] = name
}

func (m *ConcreteMediator) Send(colleague Colleague, message string) {
	// 媒介者根据接收到的消息和发送者信息，决定通知哪些同事
	for c, name := range m.colleagues {
		if c != colleague {
			c.Notify(fmt.Sprintf("From %s: %s", name, message))
		}
	}
}
