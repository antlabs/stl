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

	if s.Len() != 0 {
		t.Error("s.Len() != 0")
	}

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

	if s.Len() != 5 {
		t.Error("s.Len() != 5")
	}

	offset := unsafe.Offsetof(s.Head)

	need := []int{5, 4, 3, 2, 1}
	i := 0
	s.ForEachPrev(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		if s.ID != need[i] {
			t.Error("s.ID != need[i]\n")
		}

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

	if s.Len() != 5 {
		t.Error("s.Len() != 5")
	}

	offset := unsafe.Offsetof(s.Head)

	need := []int{5, 4, 3, 2, 1}
	i := 0
	s.ForEachPrevSafe(func(pos *Head) {

		posEntry := (*student)(pos.Entry(offset))

		s.Del(pos)
		if posEntry.ID != need[i] {
			t.Error("posEntry.ID != need[i]\n")
		}

		i++
	})

	if s.Len() != 0 {
		t.Error("s.Len() != 0\n")
	}
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

	if s.Len() != 5 {
		t.Error("s.Len() != 5")
	}

	offset := unsafe.Offsetof(s.Head)

	need := []int{1, 2, 3, 4, 5}
	i := 0
	s.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		if s.ID != need[i] {
			t.Error("s.ID != need[i]\n")
		}

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

		if s.ID != need[i] {
			t.Error("s.ID != need[i]\n")
		}

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

	if firstStudent.ID != 1 {
		t.Error("firstStudent != 1\n")
	}

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
	if lastStudent.ID != 2 {
		t.Error("lastStudent != 2\n")
	}

}

func Test_FirstEntryOrNil(t *testing.T) {
	// 返回nil
	s := student{}
	s.Init()

	offset := unsafe.Offsetof(s.Head)
	p := s.FirstEntryOrNil(offset)
	if p != nil {
		t.Error("p != nil\n")
	}

	// 返回第一个元素
	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)

	first := (*student)(s.FirstEntryOrNil(offset))
	if first.ID != 1 {
		t.Error("first.ID != 1\n")
	}

}

func Test_Replace(t *testing.T) {
	old := student{}
	old.Init()
	offset := unsafe.Offsetof(old.Head)

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}
	s4 := student{ID: 4}
	s5 := student{ID: 5}

	old.AddTail(&s1.Head)
	old.AddTail(&s2.Head)
	old.AddTail(&s3.Head)
	old.AddTail(&s4.Head)
	old.AddTail(&s5.Head)

	if old.Len() != 5 {
		t.Error("old.Len() != 5\n")
	}

	new := student{}
	new.Init()
	old.Replace(&new.Head)
	if old.Len() != 5 {
		t.Error("old.Len() != 5\n")
	}

	if new.Len() != 5 {
		t.Error("new.Len() != 5\n")
	}

	need := []int{1, 2, 3, 4, 5}
	i := 0
	new.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		if s.ID != need[i] {
			t.Error("s.ID != need[i]\n")
		}

		i++
	})

}

func Test_ReplaceInit(t *testing.T) {
	old := student{}
	old.Init()
	offset := unsafe.Offsetof(old.Head)

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}
	s4 := student{ID: 4}
	s5 := student{ID: 5}

	old.AddTail(&s1.Head)
	old.AddTail(&s2.Head)
	old.AddTail(&s3.Head)
	old.AddTail(&s4.Head)
	old.AddTail(&s5.Head)

	if old.Len() != 5 {
		t.Error("old.Len() != 5\n")
	}

	new := student{}
	new.Init()
	old.ReplaceInit(&new.Head)
	if old.Len() != 0 {
		t.Errorf("old.Len():%d != 5\n", old.Len())
	}

	if new.Len() != 5 {
		t.Errorf("new.Len():%d != 0\n", new.Len())
	}

	need := []int{1, 2, 3, 4, 5}
	i := 0
	new.ForEach(func(pos *Head) {

		s := (*student)(pos.Entry(offset))

		if s.ID != need[i] {
			t.Error("s.ID != need[i]\n")
		}

		i++
	})
}

func Test_DelInit(t *testing.T) {
	s := student{}
	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}

	s.Add(&s1.Head)
	s.Add(&s2.Head)
	if s.Len() != 2 {
		t.Error("s.Len() != 2\n")
	}

	s.DelInit(&s1.Head)

	if s1.IsLast() == false {
		t.Error("s1.IsLast() == false\n")
	}

}

func Test_Move(t *testing.T) {
	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)
	s.AddTail(&s3.Head)

	s.Move(&s2.Head)

	need := []int{2, 1, 3}
	i := 0
	offset := unsafe.Offsetof(s.Head)
	s.ForEach(func(pos *Head) {
		posEntry := (*student)(pos.Entry(offset))
		if posEntry.ID != need[i] {
			t.Error("posEntry.ID != need[i]\n")
		}

		//fmt.Printf("%d\n", posEntry.ID)
		i++
	})
}

func Test_MoveTail(t *testing.T) {
	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)
	s.AddTail(&s3.Head)

	s.MoveTail(&s2.Head)

	need := []int{1, 3, 2}
	offset := unsafe.Offsetof(s.Head)
	i := 0
	s.ForEach(func(pos *Head) {
		posEntry := (*student)(pos.Entry(offset))
		if posEntry.ID != need[i] {
			t.Error("posEntry.ID != need[i]\n")
		}
		i++
	})
}

func Test_RotateLeft(t *testing.T) {
	s := &student{}

	s.Init()

	s1 := student{ID: 1}
	s2 := student{ID: 2}
	s3 := student{ID: 3}

	s.AddTail(&s1.Head)
	s.AddTail(&s2.Head)
	s.AddTail(&s3.Head)

	s.RotateLeft()

	need := []int{2, 3, 1}
	offset := unsafe.Offsetof(s.Head)
	i := 0
	s.ForEach(func(pos *Head) {
		posEntry := (*student)(pos.Entry(offset))

		if posEntry.ID != need[i] {
			t.Error("posEntry.ID != need[i]\n")
		}
		//fmt.Printf("%d\n", posEntry.ID)
		i++
	})
}

func Test_Empty(t *testing.T) {
	s := student{}
	s.Init()

	s1 := student{ID: 1}
	s.Add(&s1.Head)
	s.Del(&s1.Head)

	if s.Empty() == false {
		t.Error("s.Empty() == false\n")
	}

}
