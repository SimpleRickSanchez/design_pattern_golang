package facade

import "fmt"

// SubsystemA 是子系统 A 的接口
type SubsystemA interface {
	OperationA() string
}

// SubsystemB 是子系统 B 的接口
type SubsystemB interface {
	OperationB() string
}

// RealSubsystemA 是子系统 A 的具体实现
type RealSubsystemA struct{}

// OperationA 实现 SubsystemA 接口
func (s *RealSubsystemA) OperationA() string {
	return "Subsystem A operation"
}

// RealSubsystemB 是子系统 B 的具体实现
type RealSubsystemB struct{}

// OperationB 实现 SubsystemB 接口
func (s *RealSubsystemB) OperationB() string {
	return "Subsystem B operation"
}

// Facade 是外观结构体，封装了子系统 A 和子系统 B
type Facade struct {
	subsystemA SubsystemA
	subsystemB SubsystemB
}

// NewFacade 创建并初始化一个 Facade 实例
func NewFacade() *Facade {
	return &Facade{
		subsystemA: &RealSubsystemA{},
		subsystemB: &RealSubsystemB{},
	}
}

// Operation 提供了统一的接口给客户端使用
func (f *Facade) Operation() string {
	resultA := f.subsystemA.OperationA()
	resultB := f.subsystemB.OperationB()
	return fmt.Sprintf("%s\n%s", resultA, resultB)
}
