package ch

import "context"

func Stream[T any](ctx context.Context, in []T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for _, element := range in {
			select {
			case <-ctx.Done():
				return
			case out <- element:
			}
		}
	}()
	return out
}

func Filter[T any](ctx context.Context, predicate func(T) bool, in <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for element := range in {
			if predicate(element) {
				select {
				case <-ctx.Done():
					return
				case out <- element:
				}
			}
		}
	}()
	return out
}

func Map[I any, O any](ctx context.Context, transformer func(I) O, in <-chan I) <-chan O {
	out := make(chan O)
	go func() {
		defer close(out)
		for element := range in {
			select {
			case <-ctx.Done():
				return
			case out <- transformer(element):
			}
		}
	}()
	return out
}
func ForEach[T any](ctx context.Context, apply func(T), in <-chan T) {
	for element := range in {
		select {
		case <-ctx.Done():
			return
		default:
			apply(element)
		}
	}
}

func Collect[T any](ctx context.Context, in <-chan T) []T {
	out := make([]T, 0)
	for element := range in {
		select {
		case <-ctx.Done():
			return out
		default:
			out = append(out, element)
		}
	}
	return out
}
