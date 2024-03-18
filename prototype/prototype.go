package prototype

import (
	"fmt"
)

// Prototype 接口定义了原型对象应该具有的方法
type Prototype interface {
	Clone() Prototype
}

// ConcretePrototype 是具体的原型对象，实现了Prototype接口
type ConcretePrototype struct {
	Value string
}

// NewConcretePrototype 创建并初始化一个ConcretePrototype对象
func NewConcretePrototype(value string) *ConcretePrototype {
	return &ConcretePrototype{Value: value}
}

// Clone 实现了Prototype接口，用于复制ConcretePrototype对象
func (cp *ConcretePrototype) Clone() *ConcretePrototype {
	// 在Go中，可以使用浅拷贝来复制结构体，因为结构体是值类型
	return &ConcretePrototype{Value: cp.Value}
}

// Display 用于显示原型对象的状态
func (cp *ConcretePrototype) Display() {
	fmt.Println("ConcretePrototype:", cp.Value)
}
