package try

func Map[IN any, OUT any](fn ResultFunc[IN, OUT]) func([]Result[IN]) []Result[OUT] {
	return func(in []Result[IN]) []Result[OUT] {
		out := make([]Result[OUT], 0, len(in))
		for _, r := range in {
			out = append(out, IfOk(r, func(in Ok[IN]) Result[OUT] {
				return fn(in)
			}))
		}
		return out
	}
}
