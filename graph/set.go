package graph

func NewSet() *Set {
	return &Set{
		m: make(map[*Node]struct{}),
	}
}

// The set is not safe for concurrent use.
type Set struct {
	m map[*Node]struct{}
}

// Add adds a node to the set, returns true if the node doesn't exist yet and false otherwise.
func (s *Set) Add(n *Node) bool {
	_, found := s.m[n]
	if found {
		return false
	}
	s.m[n] = struct{}{}
	return true
}

// Len returns the length/amount of the set.
func (s *Set) Len() int {
	return len(s.m)
}

func (s *Set) Range() <-chan *Node {
	ch := make(chan *Node, len(s.m))
	go func() {
		for n := range s.m {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

// Contains check if the given node exists inside the set.
func (s *Set) Contains(n *Node) bool {
	_, found := s.m[n]
	if found {
		return true
	}
	return false
}

// Diff returns a new set of elements which doesn't exists in the given set.
func (s *Set) Diff(other *Set) *Set {
	diff := NewSet()
	for n := range s.m {
		if !other.Contains(n) {
			diff.Add(n)
		}
	}
	return diff
}
