package rr

func SlicesIsNil[T comparable](arr []T) bool {
    return SlicesIsEmpty(arr)
}
func SlicesIsEmpty[T comparable](arr []T) bool {
    if arr == nil {
        return true
    }
    if len(arr) <= 0 {
        return true
    }
    return false
}
func SlicesUnique[T comparable](sources []T) []T {
    if sources == nil {
        return sources
    }

    // 使用map来存储唯一值
    seen := make(map[T]struct{})
    result := make([]T, 0, len(sources))

    // 遍历切片，将未见过的元素添加到结果中
    for _, v := range sources {
        if _, exists := seen[v]; !exists {
            seen[v] = struct{}{}
            result = append(result, v)
        }
    }

    return result
}

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
