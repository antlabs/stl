package list

import (
	"fmt"
	"testing"
	"unsafe"
)

type student struct {
	ID int
	Head
}

func Test_ListDel(t *testing.T) {
}

func Test_ListAddTail(t *testing.T) {
	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}
	s4 := student{ID: 4}
	s5 := student{ID: 5}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)
	s.AddTail(&s3.Head)
	s.AddTail(&s4.Head)
	s.AddTail(&s5.Head)

	fmt.Printf(":%d\n", s.Len())

	offset := unsafe.Offsetof(s.Head)

	s.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		fmt.Printf("hello world::%d\n", s.ID)
	})
}

func Test_ListAdd(t *testing.T) {

	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}
	s4 := student{ID: 4}
	s5 := student{ID: 5}

	s.Add(&s1.Head)
	s.Add(&s2.Head)
	s.Add(&s3.Head)
	s.Add(&s4.Head)
	s.Add(&s5.Head)

	fmt.Printf(":%d\n", s.Len())

	offset := unsafe.Offsetof(s.Head)

	s.ForEachSafe(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		s.Del(pos)

		fmt.Printf("hello world::%d\n", s.ID)
	})

	fmt.Printf("delete after\n")

	s.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		fmt.Printf("hello world::%d\n", s.ID)
	})

}
