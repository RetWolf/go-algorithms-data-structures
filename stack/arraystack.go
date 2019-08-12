package stack

type ArrayStack struct {
	data []int
}

// Clear removes all elements from the stack
func (s *ArrayStack) Clear() {
	s.data = nil
}

// Push puts a new element on the top of the stack
func (s *ArrayStack) Push(el int) {
	s.data = append(s.data, el)
}

// Peek returns the element on top of the stack but does not remove it
func (s *ArrayStack) Peek() int {
	return s.data[len(s.data)-1]
}

// Pop removes the top element from the stack and returns it
func (s *ArrayStack) Pop() int {
	length := len(s.data) - 1
	last := s.data[length]
	s.data = s.data[:length]
	return last
}

// Size returns the size of the stack
func (s *ArrayStack) Size() int {
	return len(s.data)
}

// Find returns true if the element exists in the stack, or false if it wasn't
func (s *ArrayStack) Find(el int) bool {
	for i := 0; i < len(s.data); i++ {
		if s.data[i] == el {
			return true
		}
	}

	return false
}

// IndexOf returns the index of the found element, or -1 if the element wasn't found
func (s *ArrayStack) IndexOf(el int) int {
	for i := 0; i < len(s.data); i++ {
		if s.data[i] == el {
			return i
		}
	}

	return -1
}
