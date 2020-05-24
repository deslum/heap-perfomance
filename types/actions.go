package types

func (h Actions) Len() int           { return len(h) }
func (h Actions) Less(i, j int) bool { return h[i].EventTimestamp < h[j].EventTimestamp && h[i].ConditionA && h[i].ConditionB}
func (h Actions) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type Action struct {
	EventTimestamp int64
	ConditionA     bool
	ConditionB     bool
}

type Actions []*Action

func (h *Actions) Push(x *Action) {
	*h = append(*h, x)
}

func (h *Actions) Pop() *Action {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
