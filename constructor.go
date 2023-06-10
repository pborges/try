package try

func Pass[V any](v V) Result[V] {
	return Ok[V]{Value: v}
}

func Fail[V any](e error) Result[V] {
	return Error{Value: e}
}

func PassSlice[V any](v []V) []Result[V] {
	out := make([]Result[V], 0, len(v))
	for _, r := range v {
		out = append(out, Pass[V](r))
	}
	return out
}
