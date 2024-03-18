package state

import "fmt"

// Context 上下文
type Context struct {
	state State
}

// SetState 设置状态
func (c *Context) SetState(s State) {
	c.state = s
}

// Request 请求操作
func (c *Context) Request() {
	c.state.Handle(c)
}

// State 抽象状态
type State interface {
	Handle(c *Context)
}

// ConcreteStateA 具体状态A
type ConcreteStateA struct{}

// Handle 处理请求
func (s *ConcreteStateA) Handle(c *Context) {
	fmt.Println("当前状态是A, 处理请求...")
	c.SetState(&ConcreteStateB{}) // 转换到下一个状态
}

// ConcreteStateB 具体状态B
type ConcreteStateB struct{}

// Handle 处理请求
func (s *ConcreteStateB) Handle(c *Context) {
	fmt.Println("当前状态是B, 处理请求...")
	// 可以在这里设置转换到其他状态，或者不转换
}
