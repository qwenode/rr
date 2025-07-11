package rr

import (
    "reflect"
    "testing"
)

func TestSlicesDelete(t *testing.T) {
    tests := []struct {
        name          string
        sources       []int
        deleteElement int
        want          []int
    }{
        {
            "删除存在的元素",
            []int{1, 2, 3, 4, 5},
            3,
            []int{1, 2, 4, 5},
        },
        {
            "删除不存在的元素",
            []int{1, 2, 3, 4, 5},
            6,
            []int{1, 2, 3, 4, 5},
        },
        {
            "删除多个相同的元素",
            []int{1, 2, 3, 3, 4, 5, 3},
            3,
            []int{1, 2, 4, 5},
        },
        {
            "空切片",
            []int{},
            1,
            []int{},
        },
        {
            "只包含要删除元素的切片",
            []int{3, 3, 3},
            3,
            []int{},
        },
        {
            "nil切片",
            nil,
            1,
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDelete(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDelete() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesDeleteString(t *testing.T) {
    tests := []struct {
        name          string
        sources       []string
        deleteElement string
        want          []string
    }{
        {
            "删除存在的字符串",
            []string{"apple", "banana", "cherry", "date", "elderberry"},
            "cherry",
            []string{"apple", "banana", "date", "elderberry"},
        },
        {
            "删除不存在的字符串",
            []string{"apple", "banana", "cherry", "date"},
            "fig",
            []string{"apple", "banana", "cherry", "date"},
        },
        {
            "删除多个相同的字符串",
            []string{"apple", "banana", "cherry", "banana", "date"},
            "banana",
            []string{"apple", "cherry", "date"},
        },
        {
            "空字符串切片",
            []string{},
            "apple",
            []string{},
        },
        {
            "只包含要删除字符串的切片",
            []string{"apple", "apple", "apple"},
            "apple",
            []string{},
        },
        {
            "删除空字符串",
            []string{"apple", "", "banana", "", "cherry"},
            "",
            []string{"apple", "banana", "cherry"},
        },
        {
            "nil字符串切片",
            nil,
            "apple",
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDelete(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDelete() = %v, want %v", got, tt.want)
            }
        })
    }
}

// 自定义类型用于测试
type Person struct {
    Name string
    Age  int
}

func TestSlicesDeleteCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}
    person3 := Person{"Charlie", 35}
    person4 := Person{"Alice", 25} // 与person1值相同但是不同实例

    tests := []struct {
        name          string
        sources       []Person
        deleteElement Person
        want          []Person
    }{
        {
            "删除存在的自定义类型元素",
            []Person{person1, person2, person3},
            person2,
            []Person{person1, person3},
        },
        {
            "删除不存在的自定义类型元素",
            []Person{person1, person2, person3},
            Person{"David", 40},
            []Person{person1, person2, person3},
        },
        {
            "删除值相同的自定义类型元素",
            []Person{person1, person2, person3},
            person4, // person4与person1值相同
            []Person{person2, person3},
        },
        {
            "空切片",
            []Person{},
            person1,
            []Person{},
        },
        {
            "只包含要删除元素的切片",
            []Person{person1, person4}, // person1和person4值相同
            person1,
            []Person{},
        },
        {
            "nil自定义类型切片",
            nil,
            person1,
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDelete(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDelete() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesDeleteArray(t *testing.T) {
    tests := []struct {
        name          string
        sources       []int
        deleteElement []int
        want          []int
    }{
        {
            "删除多个存在的元素",
            []int{1, 2, 3, 4, 5, 6, 7},
            []int{2, 4, 6},
            []int{1, 3, 5, 7},
        },
        {
            "删除不存在的元素",
            []int{1, 2, 3, 4, 5},
            []int{6, 7, 8},
            []int{1, 2, 3, 4, 5},
        },
        {
            "删除部分存在的元素",
            []int{1, 2, 3, 4, 5},
            []int{3, 6, 7},
            []int{1, 2, 4, 5},
        },
        {
            "空源切片",
            []int{},
            []int{1, 2, 3},
            []int{},
        },
        {
            "nil源切片",
            nil,
            []int{1, 2, 3},
            nil,
        },
        {
            "nil删除元素切片",
            []int{1, 2, 3, 4, 5},
            nil,
            []int{1, 2, 3, 4, 5},
        },
        {
            "空删除元素切片",
            []int{1, 2, 3, 4, 5},
            []int{},
            []int{1, 2, 3, 4, 5},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDeleteArray(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDeleteArray() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesDeleteArrayString(t *testing.T) {
    tests := []struct {
        name          string
        sources       []string
        deleteElement []string
        want          []string
    }{
        {
            "删除多个存在的字符串",
            []string{"apple", "banana", "cherry", "date", "elderberry", "fig"},
            []string{"banana", "date", "fig"},
            []string{"apple", "cherry", "elderberry"},
        },
        {
            "删除不存在的字符串",
            []string{"apple", "banana", "cherry"},
            []string{"date", "elderberry", "fig"},
            []string{"apple", "banana", "cherry"},
        },
        {
            "删除部分存在的字符串",
            []string{"apple", "banana", "cherry", "date"},
            []string{"banana", "fig", "grape"},
            []string{"apple", "cherry", "date"},
        },
        {
            "空源切片",
            []string{},
            []string{"apple", "banana"},
            []string{},
        },
        {
            "nil源切片",
            nil,
            []string{"apple", "banana"},
            nil,
        },
        {
            "nil删除元素切片",
            []string{"apple", "banana", "cherry"},
            nil,
            []string{"apple", "banana", "cherry"},
        },
        {
            "空删除元素切片",
            []string{"apple", "banana", "cherry"},
            []string{},
            []string{"apple", "banana", "cherry"},
        },
        {
            "删除包含空字符串",
            []string{"apple", "", "banana", "cherry"},
            []string{"", "banana"},
            []string{"apple", "cherry"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDeleteArray(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDeleteArray() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesDeleteArrayCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}
    person3 := Person{"Charlie", 35}
    person4 := Person{"David", 40}
    person5 := Person{"Eve", 45}
    person6 := Person{"Alice", 25} // 与person1值相同但是不同实例

    tests := []struct {
        name          string
        sources       []Person
        deleteElement []Person
        want          []Person
    }{
        {
            "删除多个存在的自定义类型元素",
            []Person{person1, person2, person3, person4, person5},
            []Person{person2, person4},
            []Person{person1, person3, person5},
        },
        {
            "删除不存在的自定义类型元素",
            []Person{person1, person2, person3},
            []Person{Person{"Frank", 50}, Person{"Grace", 55}},
            []Person{person1, person2, person3},
        },
        {
            "删除值相同的自定义类型元素",
            []Person{person1, person2, person3, person6}, // person6与person1值相同
            []Person{person1},
            []Person{person2, person3},
        },
        {
            "空源切片",
            []Person{},
            []Person{person1, person2},
            []Person{},
        },
        {
            "nil源切片",
            nil,
            []Person{person1, person2},
            nil,
        },
        {
            "nil删除元素切片",
            []Person{person1, person2, person3},
            nil,
            []Person{person1, person2, person3},
        },
        {
            "空删除元素切片",
            []Person{person1, person2, person3},
            []Person{},
            []Person{person1, person2, person3},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesDeleteArray(tt.sources, tt.deleteElement); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesDeleteArray() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUnique(t *testing.T) {
    tests := []struct {
        name    string
        sources []int
        want    []int
    }{
        {
            "无重复元素",
            []int{1, 2, 3, 4, 5},
            []int{1, 2, 3, 4, 5},
        },
        {
            "有重复元素",
            []int{1, 2, 2, 3, 3, 3, 4, 5, 5},
            []int{1, 2, 3, 4, 5},
        },
        {
            "全部重复元素",
            []int{1, 1, 1, 1, 1},
            []int{1},
        },
        {
            "空切片",
            []int{},
            []int{},
        },
        {
            "nil切片",
            nil,
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesUnique(tt.sources); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUnique() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUniqueString(t *testing.T) {
    tests := []struct {
        name    string
        sources []string
        want    []string
    }{
        {
            "无重复字符串",
            []string{"apple", "banana", "cherry"},
            []string{"apple", "banana", "cherry"},
        },
        {
            "有重复字符串",
            []string{"apple", "banana", "apple", "cherry", "banana"},
            []string{"apple", "banana", "cherry"},
        },
        {
            "全部重复字符串",
            []string{"apple", "apple", "apple"},
            []string{"apple"},
        },
        {
            "包含空字符串",
            []string{"apple", "", "banana", "", "cherry"},
            []string{"apple", "", "banana", "cherry"},
        },
        {
            "空切片",
            []string{},
            []string{},
        },
        {
            "nil切片",
            nil,
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesUnique(tt.sources); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUnique() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUniqueCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}
    person3 := Person{"Charlie", 35}
    person4 := Person{"Alice", 25} // 与person1值相同但是不同实例

    tests := []struct {
        name    string
        sources []Person
        want    []Person
    }{
        {
            "无重复结构体",
            []Person{person1, person2, person3},
            []Person{person1, person2, person3},
        },
        {
            "有重复结构体",
            []Person{person1, person2, person4, person3, person1},
            []Person{person1, person2, person3},
        },
        {
            "全部重复结构体",
            []Person{person1, person4, person1},
            []Person{person1},
        },
        {
            "空切片",
            []Person{},
            []Person{},
        },
        {
            "nil切片",
            nil,
            nil,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesUnique(tt.sources); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUnique() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsNil(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want bool
    }{
        {
            "nil切片",
            nil,
            true,
        },
        {
            "空切片",
            []int{},
            true,
        },
        {
            "非空切片",
            []int{1, 2, 3},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsNil(tt.arr); got != tt.want {
                t.Errorf("SlicesIsNil() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsNilString(t *testing.T) {
    tests := []struct {
        name string
        arr  []string
        want bool
    }{
        {
            "nil字符串切片",
            nil,
            true,
        },
        {
            "空字符串切片",
            []string{},
            true,
        },
        {
            "非空字符串切片",
            []string{"apple", "banana", "cherry"},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsNil(tt.arr); got != tt.want {
                t.Errorf("SlicesIsNil() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsNilCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}

    tests := []struct {
        name string
        arr  []Person
        want bool
    }{
        {
            "nil自定义类型切片",
            nil,
            true,
        },
        {
            "空自定义类型切片",
            []Person{},
            true,
        },
        {
            "非空自定义类型切片",
            []Person{person1, person2},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsNil(tt.arr); got != tt.want {
                t.Errorf("SlicesIsNil() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsEmpty(t *testing.T) {
    tests := []struct {
        name string
        arr  []int
        want bool
    }{
        {
            "nil切片",
            nil,
            true,
        },
        {
            "空切片",
            []int{},
            true,
        },
        {
            "非空切片",
            []int{1, 2, 3},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsEmpty(tt.arr); got != tt.want {
                t.Errorf("SlicesIsEmpty() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsEmptyString(t *testing.T) {
    tests := []struct {
        name string
        arr  []string
        want bool
    }{
        {
            "nil字符串切片",
            nil,
            true,
        },
        {
            "空字符串切片",
            []string{},
            true,
        },
        {
            "非空字符串切片",
            []string{"apple", "banana", "cherry"},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsEmpty(tt.arr); got != tt.want {
                t.Errorf("SlicesIsEmpty() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesIsEmptyCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}

    tests := []struct {
        name string
        arr  []Person
        want bool
    }{
        {
            "nil自定义类型切片",
            nil,
            true,
        },
        {
            "空自定义类型切片",
            []Person{},
            true,
        },
        {
            "非空自定义类型切片",
            []Person{person1, person2},
            false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesIsEmpty(tt.arr); got != tt.want {
                t.Errorf("SlicesIsEmpty() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesNotIn(t *testing.T) {
    tests := []struct {
        name      string
        source    []int
        reference []int
        want      []int
    }{
        {
            "正常差集情况",
            []int{1, 2, 3, 4, 5},
            []int{3, 4, 6, 7},
            []int{1, 2, 5},
        },
        {
            "无差集情况",
            []int{1, 2, 3},
            []int{1, 2, 3, 4, 5},
            []int{},
        },
        {
            "部分差集情况",
            []int{1, 2, 3, 4},
            []int{2, 4, 6, 8},
            []int{1, 3},
        },
        {
            "source完全不在reference中",
            []int{1, 2, 3},
            []int{4, 5, 6},
            []int{1, 2, 3},
        },
        {
            "空source切片",
            []int{},
            []int{1, 2, 3},
            []int{},
        },
        {
            "空reference切片",
            []int{1, 2, 3},
            []int{},
            []int{1, 2, 3},
        },
        {
            "两个空切片",
            []int{},
            []int{},
            []int{},
        },
        {
            "nil source切片",
            nil,
            []int{1, 2, 3},
            nil,
        },
        {
            "nil reference切片",
            []int{1, 2, 3},
            nil,
            []int{1, 2, 3},
        },
        {
            "两个nil切片",
            nil,
            nil,
            nil,
        },
        {
            "source有重复元素",
            []int{1, 2, 2, 3, 3, 4},
            []int{2, 5, 6},
            []int{1, 3, 3, 4},
        },
        {
            "reference有重复元素",
            []int{1, 2, 3, 4},
            []int{2, 2, 3, 3, 5},
            []int{1, 4},
        },
        {
            "单个元素测试",
            []int{1},
            []int{2},
            []int{1},
        },
        {
            "单个元素匹配",
            []int{1},
            []int{1},
            []int{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesNotIn(tt.source, tt.reference); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesNotIn() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesNotInString(t *testing.T) {
    tests := []struct {
        name      string
        source    []string
        reference []string
        want      []string
    }{
        {
            "正常字符串差集",
            []string{"apple", "banana", "cherry", "date"},
            []string{"banana", "elderberry", "fig"},
            []string{"apple", "cherry", "date"},
        },
        {
            "无差集字符串",
            []string{"apple", "banana"},
            []string{"apple", "banana", "cherry"},
            []string{},
        },
        {
            "部分差集字符串",
            []string{"apple", "banana", "cherry"},
            []string{"banana", "date", "elderberry"},
            []string{"apple", "cherry"},
        },
        {
            "包含空字符串的source",
            []string{"apple", "", "banana"},
            []string{"apple", "cherry"},
            []string{"", "banana"},
        },
        {
            "包含空字符串的reference",
            []string{"apple", "", "banana"},
            []string{"", "cherry"},
            []string{"apple", "banana"},
        },
        {
            "两个都包含空字符串",
            []string{"apple", "", "banana"},
            []string{"", "apple", "cherry"},
            []string{"banana"},
        },
        {
            "空source字符串切片",
            []string{},
            []string{"apple", "banana"},
            []string{},
        },
        {
            "空reference字符串切片",
            []string{"apple", "banana"},
            []string{},
            []string{"apple", "banana"},
        },
        {
            "nil source字符串切片",
            nil,
            []string{"apple", "banana"},
            nil,
        },
        {
            "nil reference字符串切片",
            []string{"apple", "banana"},
            nil,
            []string{"apple", "banana"},
        },
        {
            "大小写敏感测试",
            []string{"Apple", "BANANA", "cherry"},
            []string{"apple", "banana", "CHERRY"},
            []string{"Apple", "BANANA", "cherry"},
        },
        {
            "重复字符串测试",
            []string{"apple", "apple", "banana", "cherry"},
            []string{"date", "cherry"},
            []string{"apple", "apple", "banana"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesNotIn(tt.source, tt.reference); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesNotIn() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesNotInCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}
    person3 := Person{"Charlie", 35}
    person4 := Person{"David", 40}
    person5 := Person{"Eve", 45}
    person6 := Person{"Alice", 25} // 与person1值相同但是不同实例

    tests := []struct {
        name      string
        source    []Person
        reference []Person
        want      []Person
    }{
        {
            "正常自定义类型差集",
            []Person{person1, person2, person3, person4},
            []Person{person2, person5},
            []Person{person1, person3, person4},
        },
        {
            "无差集自定义类型",
            []Person{person1, person2},
            []Person{person1, person2, person3, person4},
            []Person{},
        },
        {
            "部分差集自定义类型",
            []Person{person1, person2, person3},
            []Person{person2, person4, person5},
            []Person{person1, person3},
        },
        {
            "值相同的自定义类型",
            []Person{person1, person2, person3},
            []Person{person6, person4}, // person6与person1值相同
            []Person{person2, person3},
        },
        {
            "空source自定义类型切片",
            []Person{},
            []Person{person1, person2},
            []Person{},
        },
        {
            "空reference自定义类型切片",
            []Person{person1, person2},
            []Person{},
            []Person{person1, person2},
        },
        {
            "nil source自定义类型切片",
            nil,
            []Person{person1, person2},
            nil,
        },
        {
            "nil reference自定义类型切片",
            []Person{person1, person2},
            nil,
            []Person{person1, person2},
        },
        {
            "重复自定义类型元素",
            []Person{person1, person1, person2, person3},
            []Person{person1, person4},
            []Person{person2, person3},
        },
        {
            "单个自定义类型元素",
            []Person{person1},
            []Person{person2},
            []Person{person1},
        },
        {
            "单个自定义类型匹配",
            []Person{person1},
            []Person{person1},
            []Person{},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := SlicesNotIn(tt.source, tt.reference); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesNotIn() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUniqueAppend(t *testing.T) {
    tests := []struct {
        name          string
        sources       []int
        appendElement []int
        want          []int
    }{
        {
            name:          "添加新元素",
            sources:       []int{1, 2, 3},
            appendElement: []int{4, 5},
            want:          []int{1, 2, 3, 4, 5},
        },
        {
            name:          "添加重复元素",
            sources:       []int{1, 2, 3},
            appendElement: []int{2, 3, 4},
            want:          []int{1, 2, 3, 4},
        },
        {
            name:          "sources为空切片",
            sources:       []int{},
            appendElement: []int{1, 2, 3},
            want:          []int{1, 2, 3},
        },
        {
            name:          "appendElement为空切片",
            sources:       []int{1, 2, 3},
            appendElement: []int{},
            want:          []int{1, 2, 3},
        },
        {
            name:          "sources为nil",
            sources:       nil,
            appendElement: []int{1, 2, 3},
            want:          []int{1, 2, 3},
        },
        {
            name:          "appendElement为nil",
            sources:       []int{1, 2, 3},
            appendElement: nil,
            want:          []int{1, 2, 3},
        },
        {
            name:          "全部重复元素",
            sources:       []int{1, 2, 3},
            appendElement: []int{1, 2, 3},
            want:          []int{1, 2, 3},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SlicesUniqueAppend(tt.sources, tt.appendElement)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUniqueAppend() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUniqueAppendString(t *testing.T) {
    tests := []struct {
        name          string
        sources       []string
        appendElement []string
        want          []string
    }{
        {
            name:          "添加新字符串",
            sources:       []string{"apple", "banana", "cherry"},
            appendElement: []string{"date", "elderberry"},
            want:          []string{"apple", "banana", "cherry", "date", "elderberry"},
        },
        {
            name:          "添加重复字符串",
            sources:       []string{"apple", "banana", "cherry"},
            appendElement: []string{"banana", "cherry", "date"},
            want:          []string{"apple", "banana", "cherry", "date"},
        },
        {
            name:          "sources为空切片",
            sources:       []string{},
            appendElement: []string{"apple", "banana"},
            want:          []string{"apple", "banana"},
        },
        {
            name:          "appendElement为空切片",
            sources:       []string{"apple", "banana"},
            appendElement: []string{},
            want:          []string{"apple", "banana"},
        },
        {
            name:          "sources为nil",
            sources:       nil,
            appendElement: []string{"apple", "banana"},
            want:          []string{"apple", "banana"},
        },
        {
            name:          "appendElement为nil",
            sources:       []string{"apple", "banana"},
            appendElement: nil,
            want:          []string{"apple", "banana"},
        },
        {
            name:          "全部重复字符串",
            sources:       []string{"apple", "banana"},
            appendElement: []string{"apple", "banana"},
            want:          []string{"apple", "banana"},
        },
        {
            name:          "添加空字符串",
            sources:       []string{"apple", "banana"},
            appendElement: []string{""},
            want:          []string{"apple", "banana", ""},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SlicesUniqueAppend(tt.sources, tt.appendElement)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUniqueAppend() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestSlicesUniqueAppendCustomType(t *testing.T) {
    person1 := Person{"Alice", 25}
    person2 := Person{"Bob", 30}
    person3 := Person{"Charlie", 35}
    person4 := Person{"David", 40}
    person5 := Person{"Alice", 25} // 与person1值相同但是不同实例

    tests := []struct {
        name          string
        sources       []Person
        appendElement []Person
        want          []Person
    }{
        {
            name:          "添加新自定义类型元素",
            sources:       []Person{person1, person2},
            appendElement: []Person{person3, person4},
            want:          []Person{person1, person2, person3, person4},
        },
        {
            name:          "添加重复自定义类型元素",
            sources:       []Person{person1, person2},
            appendElement: []Person{person5, person3}, // person5与person1值相同
            want:          []Person{person1, person2, person3},
        },
        {
            name:          "sources为空切片",
            sources:       []Person{},
            appendElement: []Person{person1, person2},
            want:          []Person{person1, person2},
        },
        {
            name:          "appendElement为空切片",
            sources:       []Person{person1, person2},
            appendElement: []Person{},
            want:          []Person{person1, person2},
        },
        {
            name:          "sources为nil",
            sources:       nil,
            appendElement: []Person{person1, person2},
            want:          []Person{person1, person2},
        },
        {
            name:          "appendElement为nil",
            sources:       []Person{person1, person2},
            appendElement: nil,
            want:          []Person{person1, person2},
        },
        {
            name:          "全部重复自定义类型元素",
            sources:       []Person{person1, person2},
            appendElement: []Person{person1, person2},
            want:          []Person{person1, person2},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SlicesUniqueAppend(tt.sources, tt.appendElement)
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("SlicesUniqueAppend() = %v, want %v", got, tt.want)
            }
        })
    }
}
