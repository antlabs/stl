package list

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type student struct {
	ID int
	Head
}

func Test_ListDel(t *testing.T) {
	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}

	s.Add(&s1.Head)
	s.Add(&s2.Head)
	s.Add(&s3.Head)

	offset := unsafe.Offsetof(s.Head)
	s.ForEachSafe(func(pos *Head) {
		s.Del(pos)
		posEntry := (*student)(pos.Entry(offset))
		fmt.Printf("%d\n", posEntry.ID)
	})

	assert.Equal(t, s.Len(), 0)
}

func Test_ForEachPrev(t *testing.T) {
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

	assert.Equal(t, s.Len(), 5)

	offset := unsafe.Offsetof(s.Head)

	need := []int{5, 4, 3, 2, 1}
	i := 0
	s.ForEachPrev(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		assert.Equal(t, s.ID, need[i])
		i++
	})
}

func Test_ForEachPrevSafe(t *testing.T) {
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

	assert.Equal(t, s.Len(), 5)

	offset := unsafe.Offsetof(s.Head)

	need := []int{5, 4, 3, 2, 1}
	i := 0
	s.ForEachPrevSafe(func(pos *Head) {

		posEntry := (*student)(pos.Entry(offset))

		s.Del(pos)
		assert.Equal(t, posEntry.ID, need[i])
		i++
	})

	assert.Equal(t, s.Len(), 0)
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

	assert.Equal(t, s.Len(), 5)

	offset := unsafe.Offsetof(s.Head)

	need := []int{1, 2, 3, 4, 5}
	i := 0
	s.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		assert.Equal(t, s.ID, need[i])
		i++
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

	need := []int{5, 4, 3, 2, 1}
	fmt.Printf(":%d\n", s.Len())

	offset := unsafe.Offsetof(s.Head)

	i := 0
	s.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		//s.Del(pos)

		assert.Equal(t, s.ID, need[i])
		fmt.Printf("hello world::%d\n", s.ID)
		i++
	})

}

func Test_FirstEntry(t *testing.T) {
	s := student{}
	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)

	offset := unsafe.Offsetof(s.Head)

	firstStudent := (*student)(s.FirstEntry(offset))
	assert.Equal(t, firstStudent.ID, 1)
}

func Test_lastEntry(t *testing.T) {
	s := student{}
	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)

	offset := unsafe.Offsetof(s.Head)

	lastStudent := (*student)(s.LastEntry(offset))
	assert.Equal(t, lastStudent.ID, 2)
}

func Test_FirstEntryOrNil(t *testing.T) {
	// 返回nil
	s := student{}
	s.Init()

	offset := unsafe.Offsetof(s.Head)
	p := s.FirstEntryOrNil(offset)
	assert.Equal(t, p, unsafe.Pointer(uintptr(0)))

	// 返回第一个元素
	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)

	first := (*student)(s.FirstEntryOrNil(offset))
	assert.Equal(t, first.ID, 1)

}

func Test_ReplaceInit(t *testing.T) {
}
