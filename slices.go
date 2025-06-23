package rr

func SlicesDelete[T comparable](sources []T, deleteElement T) []T {
    if sources == nil {
        return nil
    }
    res := make([]T, 0, len(sources))
    for _, source := range sources {
        if source != deleteElement {
            res = append(res, source)
        }
    }
    return res
}
