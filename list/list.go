package list

type Head struct {
	Next *Head
	Prev *Head
}

func (h *Head) Init() {
	h.Next = h
	h.Prev = h
}

func (h *Head) Add(new *Head) {
	add(new, h, h.Next)
}

func (h *Head) AddTail(new *Head) {
	add(new, h.Prev, h)
}

func (h *Head) Del(entry *Head) {
	del(entry.Prev, entry.Next)
}

func (h *Head) Replace() {
}

func (h *Head) ReplaceInit() {
}

func (h *Head) DelInit() {
}

func (h *Head) Move() {
}

func (h *Head) MoveTail() {
}

func (h *Head) IsLast() {
}

func (h *Head) Empty() bool {
	return true
}

func del(prev *Head, next *Head) {
	next.Prev = prev
	prev.Next = next
}

func add(new, prev, next *Head) {
	next.Prev = new
	new.Next = next
	new.Prev = prev
	prev.Next = new
}
