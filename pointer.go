package rr

// Deprecated:  20250722 by Node
func AsPointer[T any](v T) *T {
    return &v
}

// Deprecated:  20250722 by Node
func AsNonPointer[T any](v *T) T {
    if v == nil {
        var zero T
        return zero
    }
    return *v
}
