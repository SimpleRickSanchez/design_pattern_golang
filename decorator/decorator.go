package decorator

import "fmt"

// Shape 接口定义
type Shape interface {
	Draw() string
}

// Circle 结构体实现 Shape 接口
type Circle struct {
	Radius float64
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Drawing a circle with radius: %.2f", c.Radius)
}

// Decorator 结构体包装 Shape 接口
type Decorator struct {
	Shape Shape
}

func (d *Decorator) Draw() string {
	// 在这里可以添加额外的行为
	result := d.Shape.Draw()
	// 在这里可以添加额外的行为
	return result
}

// RedCircleDecorator 结构体扩展 Decorator
type RedCircleDecorator struct {
	Decorator Decorator
}

func (d *RedCircleDecorator) Draw() string {
	return fmt.Sprintf("Drawing a red %s", d.Decorator.Draw())
}
