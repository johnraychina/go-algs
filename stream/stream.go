package stream

import "fmt"

type Consumer[V any] interface {
	Accept(f func(V))
}

type Sink[OUT any] interface {
	Consumer[OUT]
	End()
	Begin(int)
}

type BasePipline[IN any, OUT any] struct {
	Stream[OUT]
	source        *BasePipline[IN, OUT]
	previousStage *BasePipline[IN, OUT]
	nextStage     *BasePipline[IN, OUT]
	depth         int
}

type OpStateless[IN any, OUT any] struct {
	BasePipline[IN, OUT]
}
type OpStateful[IN any, OUT any] struct {
	BasePipline[IN, OUT]
}

// todo
type Stream[V any] interface {
	Map(f func(V) V) Stream[V]
	Filter(f func(V) bool) Stream[V]
	Sum() V
	Count() int
	//Collect() []V
}

type IntStream struct {
	list []int
}

func (i IntStream) Map(f func(int) int) Stream[int] {
	Stream[int]
}

func (i IntStream) Filter(f func(v int) bool) Stream[int] {
	//TODO implement me
	panic("implement me")
}

func (i IntStream) Sum() int {
	//TODO implement me
	panic("implement me")
}

func (i IntStream) Count() int {
	//TODO implement me
	panic("implement me")
}

// operation type:
// - intermediate: map, filter
// - terminal: collect
const (
	OpTypeInter = iota
	OpTypeTerminal
)

func (receiver Sink) name() {
	ints := []int{1, 2, 3, 4}
	s := StreamOfInt(ints)
	sum := s.Filter(func(v int) bool { return v > 2 }).
		Map(func(v int) int { return v * v }).Sum()

	fmt.Println(sum)
}

func StreamOfInt(ints []int) Stream[int] {
	return &IntStream{list: ints}
}
