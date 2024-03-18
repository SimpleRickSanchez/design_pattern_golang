package strategy

// Strategy 接口定义了一个执行算法的方法
type Strategy interface {
	Execute() string
}

// ConcreteStrategyA 是实现了 Strategy 接口的具体策略 A
type ConcreteStrategyA struct{}

// Execute 是具体策略 A 的实现
func (s *ConcreteStrategyA) Execute() string {
	return "Executing strategy A"
}

// ConcreteStrategyB 是实现了 Strategy 接口的具体策略 B
type ConcreteStrategyB struct{}

// Execute 是具体策略 B 的实现
func (s *ConcreteStrategyB) Execute() string {
	return "Executing strategy B"
}

// Context 结构体维护对 Strategy 对象的引用，并定义了使用策略的方法
type Context struct {
	strategy Strategy
}

// SetStrategy 方法允许客户端设置一个新的策略
func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

// ExecuteStrategy 方法使用当前设置的策略来执行算法
func (c *Context) ExecuteStrategy() string {
	return c.strategy.Execute()
}
