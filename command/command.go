package command

import "fmt"

// Receiver 接收者接口
type Receiver interface {
	Action()
}

// ConcreteReceiver 具体接收者
type ConcreteReceiver struct {
	Name string
}

// Action 实现Receiver接口的方法
func (r *ConcreteReceiver) Action() {
	fmt.Printf("Receiver %s action performed\n", r.Name)
}

// Command 命令接口
type Command interface {
	Execute(receiver Receiver)
}

// ConcreteCommand 具体命令
type ConcreteCommand struct {
	Receiver Receiver
}

// Execute 实现Command接口的方法
func (c *ConcreteCommand) Execute(receiver Receiver) {
	receiver.Action()
}

// Invoker 调用者
type Invoker struct {
	command Command
}

// SetCommand 设置命令
func (i *Invoker) SetCommand(cmd Command) {
	i.command = cmd
}

// ExecuteCommand 执行命令
func (i *Invoker) ExecuteCommand(receiver Receiver) {
	i.command.Execute(receiver)
}
