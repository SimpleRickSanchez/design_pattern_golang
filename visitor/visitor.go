package visitor

import "fmt"

// Visitor 接口
type Visitor interface {
	VisitDog(dog *Dog)
	VisitCat(cat *Cat)
}

// ConcreteVisitor 实现 Visitor 接口
type SoundVisitor struct{}

func (v *SoundVisitor) VisitDog(dog *Dog) {
	fmt.Println(dog.Name, " Dog says: Woof!")
}

func (v *SoundVisitor) VisitCat(cat *Cat) {
	fmt.Println(cat.Name, " Cat says: Meow!")
}

type EatVisitor struct{}

func (v *EatVisitor) VisitDog(dog *Dog) {
	fmt.Println(dog.Name, " Dog eats: Dog food!")
}

func (v *EatVisitor) VisitCat(cat *Cat) {
	fmt.Println(cat.Name, " Cat eats: Fish!")
}

// Element 接口
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElement 实现 Element 接口
type Dog struct {
	Name string
}

func (d *Dog) Accept(visitor Visitor) {
	visitor.VisitDog(d)
}

type Cat struct {
	Name string
}

func (c *Cat) Accept(visitor Visitor) {
	visitor.VisitCat(c)
}

// ObjectStructure 管理 Element 对象的集合
type Zoo struct {
	Animals []Element
}

func (z *Zoo) Accept(visitor Visitor) {
	for _, animal := range z.Animals {
		animal.Accept(visitor)
	}
}
