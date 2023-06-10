package try

type Result[V any] interface {
	isResult()
}

type Ok[V any] struct {
	Value V
}

func (Ok[V]) isResult() {}

type Error struct {
	Value error
}

func (Error) isResult() {}
