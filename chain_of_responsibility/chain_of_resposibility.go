package chain_of_responsibility

import (
	"fmt"
)

// Handler 接口定义处理请求的方法
type Handler interface {
	HandleRequest(request interface{})
}

// ConcreteHandlerA 是处理请求的具体实现类 A
type ConcreteHandlerA struct {
	successor Handler
}

// SetSuccessor 方法用于设置下一个处理者
func (h *ConcreteHandlerA) SetSuccessor(successor Handler) {
	h.successor = successor
}

// HandleRequest 方法处理请求，如果满足条件则处理，否则传递给下一个处理者
func (h *ConcreteHandlerA) HandleRequest(request interface{}) {
	if request == "requestA" {
		fmt.Println("ConcreteHandlerA handles requestA")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

// ConcreteHandlerB 是处理请求的具体实现类 B
type ConcreteHandlerB struct {
	successor Handler
}

// SetSuccessor 方法用于设置下一个处理者
func (h *ConcreteHandlerB) SetSuccessor(successor Handler) {
	h.successor = successor
}

// HandleRequest 方法处理请求，如果满足条件则处理，否则传递给下一个处理者
func (h *ConcreteHandlerB) HandleRequest(request interface{}) {
	if request == "requestB" {
		fmt.Println("ConcreteHandlerB handles requestB")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}
