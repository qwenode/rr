package random

import (
    "testing"
)

func TestString(t *testing.T) {
    for i := 0; i < 1000; i++ {
        _ = String(100)
        
    }
}

func TestIntRange(t *testing.T) {
    // Basic bounds checks (inclusive max).
    for i := 0; i < 200; i++ {
        v := IntRange(1, 1000)
        if v < 1 || v > 1000 {
            t.Fatalf("IntRange(1,1000) out of range: %d", v)
        }
    }

    // Adjacent bounds should be able to produce both values.
    seen01 := map[int]bool{}
    seen12 := map[int]bool{}
    for i := 0; i < 2000; i++ {
        seen01[IntRange(0, 1)] = true
        seen12[IntRange(1, 2)] = true
        if seen01[0] && seen01[1] && seen12[1] && seen12[2] {
            break
        }
    }
    if !(seen01[0] && seen01[1]) {
        t.Fatalf("IntRange(0,1) did not produce both values; seen=%v", seen01)
    }
    if !(seen12[1] && seen12[2]) {
        t.Fatalf("IntRange(1,2) did not produce both values; seen=%v", seen12)
    }

    // Single-value and reversed range behavior.
    if v := IntRange(5, 5); v != 5 {
        t.Fatalf("IntRange(5,5) expected 5, got %d", v)
    }
    for i := 0; i < 200; i++ {
        v := IntRange(10, 0)
        if v < 0 || v > 10 {
            t.Fatalf("IntRange(10,0) out of range: %d", v)
        }
    }
}

func TestGetInt(t *testing.T) {
    for i := 1; i < 1000; i++ {
        _ = getInt(i * 1000)
        
    }
}

func TestRandom(t *testing.T) {
    for i := 0; i < 1000; i++ {
        _ = Random(i, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
        
    }
}

func TestChoice(t *testing.T) {
    s := []string{}
    for i := 0; i < 1000; i++ {
        random_string := Random(i, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
        
        s = append(s, random_string)
    }
    for i := 0; i < 1000; i++ {
        Choice(s)
    }
    
}

func TestStringRange(t *testing.T) {
    for i := 0; i < 200; i++ {
        _ = StringRange(20, 80)
        
    }
}

func TestBytes(t *testing.T) {
    size := 1048576 // 1 MB
    
    for i := 0; i < 100; i++ {
        _, err := Bytes(size)
        if err != nil {
            t.Error(err)
            return
        }
    }
}
