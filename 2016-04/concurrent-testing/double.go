package main

import "sync"

// Simple test double example

type MyType int

func (mt MyType) Value() int {
	return int(mt)
}

type MyDouble struct {
	sync.Mutex
	MyType
}

func (md *MyDouble) Value() int {
	md.Lock()
	defer md.Unlock()
	return md.MyType.Value()
}

func (md *MyDouble) Set(i int) {
	md.Lock()
	md.MyType = MyType(i)
	md.Unlock()
}

// END OMIT
