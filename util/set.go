package util

// Empty Struct
type Empty struct {
}

var (
	empty Empty
)

// Set UnOrder set
type Set[T int64 | float64 | string] struct {
	set map[T]Empty
}

func NewSet[T int64 | float64 | string]() *Set[T] {
	return &Set[T]{
		set: map[T]Empty{},
	}
}

func NewSetWithParam[T int64 | float64 | string](param ...T) *Set[T] {
	set := NewSet[T]()
	for _, t := range param {
		set.Add(t)
	}
	return set
}

func (set *Set[T]) Add(x T) {
	set.set[x] = empty
}

func (set *Set[T]) Contains(x T) bool {
	_, ok := set.set[x]
	return ok
}

func (set *Set[T]) Remove(x T) {
	delete(set.set, x)
}

func (set *Set[T]) Export() []T {
	res := make([]T, len(set.set))[0:0]
	for k, _ := range set.set {
		res = append(res, k)
	}
	return res
}
