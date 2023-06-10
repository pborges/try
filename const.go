package try

type Func[IN any, OUT any] func(in IN) (OUT, error)

type ResultFunc[IN any, OUT any] func(in Result[IN]) Result[OUT]

type ResultOkFunc[IN any, OUT any] func(in Ok[IN]) Result[OUT]