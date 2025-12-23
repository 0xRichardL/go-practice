package algorithms

type SortedSet struct {
	M   map[int]bool
	Arr []int
}

func NewSortedSet() *SortedSet {
	return &SortedSet{
		M:   make(map[int]bool),
		Arr: make([]int, 0),
	}
}

func (s *SortedSet) Add(value int) {
	if s.M[value] {
		return
	}
	s.M[value] = true
	l, r := 0, len(s.Arr)
	for l < r {
		mid := (l + r) / 2
		if s.Arr[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// insert at position l
	s.Arr = append(s.Arr, 0)
	copy(s.Arr[l+1:], s.Arr[l:])
	s.Arr[l] = value
}
