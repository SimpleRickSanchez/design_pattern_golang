package singleton

import (
	"sync"
)

// lazy initialization
type Singleton1 struct {
	Data int
}

var instance1 *Singleton1
var once1 sync.Once

func GetInstance1() *Singleton1 {
	once1.Do(func() { // thread safe
		instance1 = &Singleton1{}
	})
	return instance1
}

// eager initialization
type Singleton2 struct {
	Data int
}

var instance *Singleton2 = &Singleton2{} // thread safe

func GetInstance2() *Singleton2 {
	return instance
}

// singleton factory
type Singleton3 struct {
	Data int
}

var once2 sync.Once
var instance3 *Singleton3

func GetInstance3() *Singleton3 {
	once2.Do(func() {
		instance3 = &Singleton3{}
	})
	return instance3
}

// SingletonFactory is a factory function that returns a Singleton.
func SingletonFactory() *Singleton3 {
	return GetInstance3()
}
