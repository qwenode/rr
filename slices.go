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

// 返回不在 reference 中的元素
func SlicesNotIn[T comparable](source []T, reference []T) []T {

    // 处理边界情况
    if source == nil || reference == nil || len(source) == 0 {
        return source
    }

    // 创建map存储 reference 切片的元素
    referenceMap := make(map[T]struct{})
    for _, v := range reference {
        referenceMap[v] = struct{}{}
    }

    // 创建结果切片存储不在 reference 中的元素
    result := make([]T, 0)

    // 遍历 source 切片，找出在 reference 中不存在的元素
    for _, v := range source {
        if _, exists := referenceMap[v]; !exists {
            result = append(result, v)
        }
    }

    return result
}

// 将 appendElement 中的元素添加到 sources 中，如果 sources 中已经存在，则不添加
func SlicesUniqueAppend[T comparable](sources []T, appendElement []T) []T {
    if appendElement == nil {
        return sources
    }
    if sources == nil {
        return appendElement
    }
    exists := make(map[T]struct{})
    for _, v := range sources {
        exists[v] = struct{}{}
    }
    for _, v := range appendElement {
        if _, found := exists[v]; !found {
            exists[v] = struct{}{}
            sources = append(sources, v)
        }
    }
    return sources
}
