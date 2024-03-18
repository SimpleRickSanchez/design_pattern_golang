package adapter

// Target 是客户端期望的接口
type Target interface {
	Request() string
}

// Adaptee 是需要适配的接口
type Adaptee struct {
	Data string
}

func (a *Adaptee) SpecificRequest() string {
	return a.Data
}

// Adapter 是适配器结构体，它持有一个Adaptee类型的实例
type Adapter struct {
	adaptee *Adaptee
}

// NewAdapter 创建一个新的适配器实例
func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

// Request 实现了Target接口的方法
func (a *Adapter) Request() string {
	return a.adaptee.SpecificRequest()
}
