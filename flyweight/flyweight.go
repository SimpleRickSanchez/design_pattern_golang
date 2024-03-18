package flyweight

import (
	"fmt"
	"sync"
)

// Flyweight 接口定义了享元对象的公共行为
type Flyweight interface {
	Operation(extrinsicState string) string
}

// ConcreteFlyweight 实现了 Flyweight 接口
type ConcreteFlyweight struct {
	// 内部状态，对于所有实例是共享的
	intrinsicState string
}

// Operation 方法实现了 Flyweight 接口，并接收一个外部状态作为参数
func (fw *ConcreteFlyweight) Operation(extrinsicState string) string {
	return fmt.Sprintf("Intrinsic state: %s, Extrinsic state: %s", fw.intrinsicState, extrinsicState)
}

// FlyweightFactory 是享元工厂，负责创建和管理 Flyweight 对象
type FlyweightFactory struct {
	// 使用 sync.Map 作为缓存来存储 Flyweight 对象
	flyweights sync.Map
}

// GetFlyweight 方法根据内部状态获取或创建 Flyweight 对象
func (f *FlyweightFactory) GetFlyweight(key string) Flyweight {
	// 尝试从缓存中获取 Flyweight 对象
	if flyweight, ok := f.flyweights.Load(key); ok {
		return flyweight.(Flyweight)
	}

	// 如果缓存中没有，则创建一个新的 Flyweight 对象并放入缓存
	flyweight := &ConcreteFlyweight{intrinsicState: key}
	f.flyweights.Store(key, flyweight)
	return flyweight
}
