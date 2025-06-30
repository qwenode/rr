package rr

func SlicesDelete[T comparable](sources []T, deleteElement T) []T {
    if sources == nil {
        return sources
    }
    res := make([]T, 0, len(sources))
    for _, source := range sources {
        if source != deleteElement {
            res = append(res, source)
        }
    }
    return res
}

func SlicesDeleteArray[T comparable](sources []T, deleteElement []T) []T {
    if sources == nil || deleteElement == nil {
        return sources
    }
    res := make([]T, 0, len(sources))

    // Create a map to store elements to delete
    deleteMap := make(map[T]struct{})
    for _, del := range deleteElement {
        deleteMap[del] = struct{}{}
    }

    // Add elements not in deleteMap to result
    for _, source := range sources {
        if _, exists := deleteMap[source]; !exists {
            res = append(res, source)
        }
    }

    return res
}
