package list

import (
	"container/list"
	"testing"
	"time"
)

type timeNodeStdlib struct {
	expire     uint64
	userExpire time.Duration
	callback   func()
	isSchedule bool
	close      uint32
	lock       uint32
}

type timeNode struct {
	expire     uint64
	userExpire time.Duration
	callback   func()
	isSchedule bool
	close      uint32
	lock       uint32

	Head
}

func Benchmark_ListAdd_Stdlib(b *testing.B) {
	head := list.New()
	for i := 0; i < b.N; i++ {
		node := timeNodeStdlib{}
		head.PushBack(node)
	}
}

func Benchmark_ListAdd(b *testing.B) {
	head := timeNode{}
	head.Init()

	for i := 0; i < b.N; i++ {
		node := timeNode{}
		head.AddTail(&node.Head)
	}
}
