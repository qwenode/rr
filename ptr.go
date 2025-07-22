package rr

// 任意转为指针 20250722
func ToPtr[T any](v T) *T {
    return &v
}

// 去掉指针 20250722
func FromPtr[T any](v *T) T {
    if v == nil {
        var zero T
        return zero
    }
    return *v
}
