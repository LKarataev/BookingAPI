package utils

import (
	"sync"
)

type StringKeyLockerInterface interface {
	init()
	getLockBy(key string) *sync.Mutex
	Lock(key string)
	Unlock(key string)
}

type stringKeyLocker struct {
	locks   map[string]*sync.Mutex
	mapLock sync.Mutex
}

func (l *stringKeyLocker) init() {
	l.locks = map[string]*sync.Mutex{}
}

func (l *stringKeyLocker) getLockBy(key string) *sync.Mutex {
	l.mapLock.Lock()
	defer l.mapLock.Unlock()

	mutex, found := l.locks[key]
	if found {
		return mutex
	}

	mutex = &sync.Mutex{}
	l.locks[key] = mutex
	return mutex
}

func (l *stringKeyLocker) Lock(key string) {
	l.getLockBy(key).Lock()
}

func (l *stringKeyLocker) Unlock(key string) {
	l.getLockBy(key).Unlock()
}

var lockerImpl StringKeyLockerInterface
var lock sync.Mutex

func GetStringKeyLocker() StringKeyLockerInterface {
	lock.Lock()
	defer lock.Unlock()

	if lockerImpl == nil {
		lockerImpl = new(stringKeyLocker)
		lockerImpl.init()
	}
	return lockerImpl
}
