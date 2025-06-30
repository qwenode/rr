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