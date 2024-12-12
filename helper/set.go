package helper

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := Set[T]{}
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Contains(wanted T) bool {
	_, ok := s[wanted]
	return ok
}

func (s Set[T]) Add(vals ...T) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Members() []T {
	members := make([]T, len(s))
	i := 0
	for v := range s {
		members[i] = v
		i++
	}
	return members
}
