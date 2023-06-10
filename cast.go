package try

func IfOk[IN any, OUT any](in Result[IN], fn ResultOkFunc[IN, OUT]) Result[OUT] {
	if val, ok := in.(Ok[IN]); ok {
		return fn(val)
	}
	return in
}

func To[IN any, OUT any](fn Func[IN, OUT]) ResultFunc[IN, OUT] {
	return func(in Result[IN]) Result[OUT] {
		if val, ok := in.(Ok[IN]); ok {
			res, err := fn(val.Value)
			if err != nil {
				return Fail[OUT](err)
			}
			return Pass[OUT](res)
		}
		return in
	}
}

func Collect[V any](res []Result[V]) ([]V, []error) {
	oks := make([]V, 0, len(res))
	errs := make([]error, 0, len(res))
	for _, r := range res {
		switch c := r.(type) {
		case Ok[V]:
			oks = append(oks, c.Value)
		case Error:
			errs = append(errs, c.Value)
		}
	}
	return oks, errs
}
