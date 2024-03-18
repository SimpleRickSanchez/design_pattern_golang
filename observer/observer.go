package observer

import "fmt"

// Observer 定义观察者接口
type Observer interface {
	Update(subject *ConcreteSubject)
}

// Subject 定义主题接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConcreteSubject 是主题的具体实现
type ConcreteSubject struct {
	observers []Observer
	state     string
}

// NewConcreteSubject 创建一个新的主题实例
func NewConcreteSubject() *ConcreteSubject {
	return &ConcreteSubject{}
}

// RegisterObserver 实现主题接口，用于注册观察者
func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

// RemoveObserver 实现主题接口，用于注销观察者
func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers 实现主题接口，通知所有观察者状态已改变
func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

// SetState 设置主题状态，并在状态改变时通知观察者
func (s *ConcreteSubject) SetState(state string) {
	s.state = state
	s.NotifyObservers()
}

// GetState 获取主题状态
func (s *ConcreteSubject) GetState() string {
	return s.state
}

// ConcreteObserver 是观察者的具体实现
type ConcreteObserver struct {
	name string
}

// NewConcreteObserver 创建一个新的观察者实例
func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

// Update 实现观察者接口，当主题状态改变时调用
func (o *ConcreteObserver) Update(subject *ConcreteSubject) {
	fmt.Printf("Observer %s received update: %s\n", o.name, subject.GetState())
}
