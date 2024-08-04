package rr

import (
    "reflect"
    "testing"
)

func TestPointer(t *testing.T) {
    type args[T any] struct {
        v T
    }
    type testCase[T any] struct {
        name string
        args args[T]
        want *T
    }
    a := 1
    tests := []testCase[int]{
        {
            name: "test1",
            args: args[int]{1},
            want: &a,
        },
    }
    for _, tt := range tests {
        t.Run(
            tt.name, func(t *testing.T) {
                if got := AsPointer(tt.args.v); !reflect.DeepEqual(got, tt.want) {
                    t.Errorf("Pointer() = %v, want %v", got, tt.want)
                }
            },
        )
    }
}
