package iterator

// Iterator 定义迭代器接口
type Iterator interface {
	Next() bool    // 移动到下一个元素
	HasNext() bool // 检查是否还有下一个元素
	Current() int  // 返回当前元素
	Reset()        // 重置迭代器到初始位置
}

// ConcreteIterator 实现具体迭代器
type ConcreteIterator struct {
	list    []int
	index   int
	hasNext bool
}

func NewConcreteIterator(list []int) *ConcreteIterator {
	return &ConcreteIterator{
		list:    list,
		index:   0,
		hasNext: true,
	}
}

func (it *ConcreteIterator) Next() bool {
	it.index++
	it.hasNext = it.index < len(it.list)
	return it.hasNext
}

func (it *ConcreteIterator) HasNext() bool {
	return it.hasNext
}

func (it *ConcreteIterator) Current() int {
	return it.list[it.index]
}

func (it *ConcreteIterator) Reset() {
	it.index = 0
	it.hasNext = true
}

// Aggregate 定义聚合接口
type Aggregate interface {
	CreateIterator() Iterator
}

// ConcreteAggregate 实现具体聚合
type ConcreteAggregate struct {
	list []int
}

func NewConcreteAggregate(list []int) *ConcreteAggregate {
	return &ConcreteAggregate{
		list: list,
	}
}

func (agg *ConcreteAggregate) CreateIterator() Iterator {
	return NewConcreteIterator(agg.list)
}
