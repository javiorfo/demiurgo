package steams

import "slices"

type List[T any] []T

func ListOf[T any](args ...T) Steam[T] {
	return List[T](args)
}

func (list List[T]) Filter(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

func (list List[T]) MapToAny(mapper func(T) any) Steam[any] {
	results := make(List[any], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) MapToString(mapper func(T) string) Steam[string] {
	results := make(List[string], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) MapToInt(mapper func(T) int) Steam[int] {
	results := make(List[int], len(list))
	for i, v := range list {
		results[i] = mapper(v)
	}
	return results
}

func (list List[T]) FilterMapToAny(predicate func(T) bool, mapper func(T) any) Steam[any] {
	results := make(List[any], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, mapper(v))
		}
	}
	return results
}

func (list List[T]) FlatMap(mapper func(T) Steam[T]) Steam[T] {
	results := make(List[T], len(list))
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[T]))
	}
	return results
}

func (list List[T]) FlatMapToAny(mapper func(T) Steam[any]) Steam[any] {
	results := make(List[any], len(list))
	for _, v := range list {
		results = slices.Concat(results, mapper(v).(List[any]))
	}
	return results
}

func (list List[T]) Limit(limit int) Steam[T] {
	results := make(List[T], 0)
	for i := 0; i < len(list) && i < limit; i++ {
		results = append(results, list[i])
	}
	return results
}

func (list List[T]) Count() int {
	return len(list)
}

func (list List[T]) ForEach(consumer func(T)) {
	for _, v := range list {
		consumer(v)
	}
}

func (list List[T]) Peek(consumer func(T)) Steam[T] {
	for _, v := range list {
		consumer(v)
	}
	return list
}

func (list List[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (list List[T]) AnyMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (list List[T]) NoneMatch(predicate func(T) bool) bool {
	for _, v := range list {
		if predicate(v) {
			return false
		}
	}
	return true
}

func (list List[T]) FindFirst() (T, bool) {
	if len(list) > 0 {
		return list[0], true
	}
	return *new(T), false
}

func (list List[T]) TakeWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if predicate(v) {
			results = append(results, v)
		} else {
			break
		}
	}
	return results
}

func (list List[T]) DropWhile(predicate func(T) bool) Steam[T] {
	results := make(List[T], 0)
	for _, v := range list {
		if !predicate(v) {
			results = append(results, v)
		}
	}
	return results
}

func (list List[T]) Skip(n int) Steam[T] {
	length := len(list)
	if length > n {
		length = length - n
	} else {
		return *new(List[T])
	}

	results := make(List[T], length)
	for i := 0; i < length; i++ {
		results[i] = list[i+n]
	}
	return results
}

func (list List[T]) Collect() []T {
	return list
}
