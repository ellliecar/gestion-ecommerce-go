package utils

func FilterBy[T any](items []T, predicate func(T) bool) []T {
    result := make([]T, 0)
    for _, item := range items {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

func Reduce[T, U any](items []T, initial U, accumulator func(U, T) U) U {
    result := initial
    for _, item := range items {
        result = accumulator(result, item)
    }
    return result
}
