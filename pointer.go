package rr

func AsPointer[T any](v T) *T {
    return &v
}

func AsNonPointer[T any](v *T) T {
    if v == nil {
        var zero T
        return zero
    }
    return *v
}
