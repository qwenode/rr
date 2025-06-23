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