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
		_, err := Choice(s)
		if err != nil {
			t.Fatalf("Choice returned error unexpectedly: %v", err)
		}
    }
    
}

func TestRandomEdges(t *testing.T) {
	// Should not panic; n<=0 defaults to 10.
	if got := Random(0, "abc"); len(got) != 10 {
		t.Fatalf("Random(0,...) expected len=10, got len=%d (%q)", len(got), got)
	}
	if got := Random(-1, "abc"); len(got) != 10 {
		t.Fatalf("Random(-1,...) expected len=10, got len=%d (%q)", len(got), got)
	}
	if got := Random(10, ""); got != "" {
		t.Fatalf("Random(10,\"\") expected empty string, got %q", got)
	}
	if got := String(0); len(got) != 10 {
		t.Fatalf("String(0) expected len=10, got len=%d (%q)", len(got), got)
	}
}

func TestChoiceEmpty(t *testing.T) {
	if _, err := Choice(nil); err == nil {
		t.Fatalf("Choice(nil) expected error")
	}
	if _, err := Choice([]string{}); err == nil {
		t.Fatalf("Choice(empty) expected error")
	}
}

func TestBytesEdges(t *testing.T) {
	b, err := Bytes(0)
	if err == nil {
		t.Fatalf("Bytes(0) expected err")
	}
	if len(b) != 0 {
		t.Fatalf("Bytes(0) expected empty slice, got %d", len(b))
	}
	b, err = Bytes(-1)
	if err == nil {
		t.Fatalf("Bytes(-1) expected err")
	}
	if len(b) != 0 {
		t.Fatalf("Bytes(-1) expected empty slice, got %d", len(b))
	}
}

func TestGetIntInsecureEdges(t *testing.T) {
	assertPanic := func(fn func(), name string) {
		t.Helper()
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s expected panic", name)
			}
		}()
		fn()
	}
	assertPanic(func() { _ = GetIntInsecure(0) }, "GetIntInsecure(0)")
	assertPanic(func() { _ = GetIntInsecure(-1) }, "GetIntInsecure(-1)")
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
