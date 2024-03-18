package composite

import "fmt"

// Component 是组件接口
type Component interface {
	Operation()
	Add(Component)
	Remove(Component)
	GetChild(int) Component
}

// Leaf 是叶子组件，它没有子组件
type Leaf struct {
	name string
}

// NewLeaf 创建一个新的 Leaf 实例
func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

// Operation 实现 Component 接口的方法
func (l *Leaf) Operation() {
	fmt.Printf("Leaf %s: Operation()\n", l.name)
}

// Add 实现 Component 接口的方法，叶子组件不能添加子组件
func (l *Leaf) Add(Component) {
	fmt.Println("Leaf cannot have children")
}

// Remove 实现 Component 接口的方法，叶子组件不能移除子组件
func (l *Leaf) Remove(Component) {
	fmt.Println("Leaf cannot have children")
}

// GetChild 实现 Component 接口的方法，叶子组件没有子组件
func (l *Leaf) GetChild(int) Component {
	return nil
}

// Composite 是复合组件，它可以包含子组件
type Composite struct {
	name     string
	children []Component
}

// NewComposite 创建一个新的 Composite 实例
func NewComposite(name string) *Composite {
	return &Composite{
		name:     name,
		children: make([]Component, 0),
	}
}

// Operation 实现 Component 接口的方法
func (c *Composite) Operation() {
	fmt.Printf("Composite %s: Operation()\n", c.name)
	for _, child := range c.children {
		child.Operation()
	}
}

// Add 实现 Component 接口的方法，用于向复合组件中添加子组件
func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

// Remove 实现 Component 接口的方法，用于从复合组件中移除子组件
func (c *Composite) Remove(component Component) {
	for i, child := range c.children {
		if child == component {
			c.children = append(c.children[:i], c.children[i+1:]...)
			return
		}
	}
	fmt.Println("Component not found")
}

// GetChild 实现 Component 接口的方法，用于获取指定索引的子组件
func (c *Composite) GetChild(index int) Component {
	if index < 0 || index >= len(c.children) {
		return nil
	}
	return c.children[index]
}
