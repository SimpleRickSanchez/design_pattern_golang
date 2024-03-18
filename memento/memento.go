package memento

// Originator 发起人
type Originator struct {
	state string
}

// SetState 设置发起人的状态
func (o *Originator) SetState(state string) {
	o.state = state
}

// GetState 获取发起人的当前状态
func (o *Originator) GetState() string {
	return o.state
}

// CreateMemento 创建并返回发起人当前状态的备忘录
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// RestoreFromMemento 使用备忘录恢复发起人的状态
func (o *Originator) RestoreFromMemento(memento *Memento) {
	o.state = memento.state
}

// Memento 备忘录
type Memento struct {
	state string
}

// GetState 返回备忘录中保存的状态
func (m *Memento) GetState() string {
	return m.state
}

// Caretaker 管理者
type Caretaker struct {
	mementos []*Memento
}

// AddMemento 添加备忘录到管理者
func (c *Caretaker) AddMemento(memento *Memento) {
	c.mementos = append(c.mementos, memento)
}

// GetMemento 获取管理者保存的最后一个备忘录
func (c *Caretaker) GetMemento() *Memento {
	if len(c.mementos) == 0 {
		return nil
	}
	return c.mementos[len(c.mementos)-1]
}
