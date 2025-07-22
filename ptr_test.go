package rr

import (
	"testing"
)

func TestToPtr(t *testing.T) {
	// Test with int
	intVal := 42
	intPtr := ToPtr(intVal)
	if intPtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if *intPtr != intVal {
		t.Errorf("Expected %d, got %d", intVal, *intPtr)
	}

	// Test with string
	strVal := "hello"
	strPtr := ToPtr(strVal)
	if strPtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if *strPtr != strVal {
		t.Errorf("Expected %s, got %s", strVal, *strPtr)
	}

	// Test with bool
	boolVal := true
	boolPtr := ToPtr(boolVal)
	if boolPtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if *boolPtr != boolVal {
		t.Errorf("Expected %t, got %t", boolVal, *boolPtr)
	}

	// Test with float64
	floatVal := 3.14
	floatPtr := ToPtr(floatVal)
	if floatPtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if *floatPtr != floatVal {
		t.Errorf("Expected %f, got %f", floatVal, *floatPtr)
	}

	// Test with struct
	type TestStruct struct {
		Name string
		Age  int
	}
	structVal := TestStruct{Name: "Alice", Age: 30}
	structPtr := ToPtr(structVal)
	if structPtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if *structPtr != structVal {
		t.Errorf("Expected %+v, got %+v", structVal, *structPtr)
	}

	// Test with slice
	sliceVal := []int{1, 2, 3}
	slicePtr := ToPtr(sliceVal)
	if slicePtr == nil {
		t.Error("ToPtr should not return nil")
	}
	if len(*slicePtr) != len(sliceVal) {
		t.Errorf("Expected slice length %d, got %d", len(sliceVal), len(*slicePtr))
	}
	for i, v := range sliceVal {
		if (*slicePtr)[i] != v {
			t.Errorf("Expected slice[%d] = %d, got %d", i, v, (*slicePtr)[i])
		}
	}
}

func TestFromPtr(t *testing.T) {
	// Test with non-nil int pointer
	intVal := 42
	intPtr := &intVal
	result := FromPtr(intPtr)
	if result != intVal {
		t.Errorf("Expected %d, got %d", intVal, result)
	}

	// Test with nil int pointer
	var nilIntPtr *int
	result = FromPtr(nilIntPtr)
	if result != 0 {
		t.Errorf("Expected 0 for nil int pointer, got %d", result)
	}

	// Test with non-nil string pointer
	strVal := "hello"
	strPtr := &strVal
	strResult := FromPtr(strPtr)
	if strResult != strVal {
		t.Errorf("Expected %s, got %s", strVal, strResult)
	}

	// Test with nil string pointer
	var nilStrPtr *string
	strResult = FromPtr(nilStrPtr)
	if strResult != "" {
		t.Errorf("Expected empty string for nil string pointer, got %s", strResult)
	}

	// Test with non-nil bool pointer
	boolVal := true
	boolPtr := &boolVal
	boolResult := FromPtr(boolPtr)
	if boolResult != boolVal {
		t.Errorf("Expected %t, got %t", boolVal, boolResult)
	}

	// Test with nil bool pointer
	var nilBoolPtr *bool
	boolResult = FromPtr(nilBoolPtr)
	if boolResult != false {
		t.Errorf("Expected false for nil bool pointer, got %t", boolResult)
	}

	// Test with struct
	type TestStruct struct {
		Name string
		Age  int
	}
	structVal := TestStruct{Name: "Bob", Age: 25}
	structPtr := &structVal
	structResult := FromPtr(structPtr)
	if structResult != structVal {
		t.Errorf("Expected %+v, got %+v", structVal, structResult)
	}

	// Test with nil struct pointer
	var nilStructPtr *TestStruct
	structResult = FromPtr(nilStructPtr)
	expectedZero := TestStruct{}
	if structResult != expectedZero {
		t.Errorf("Expected zero struct %+v, got %+v", expectedZero, structResult)
	}
}

// Test round-trip conversion
func TestToPtrFromPtrRoundTrip(t *testing.T) {
	// Test int round-trip
	intVal := 123
	roundTripInt := FromPtr(ToPtr(intVal))
	if roundTripInt != intVal {
		t.Errorf("Round-trip failed for int: expected %d, got %d", intVal, roundTripInt)
	}

	// Test string round-trip
	strVal := "test"
	roundTripStr := FromPtr(ToPtr(strVal))
	if roundTripStr != strVal {
		t.Errorf("Round-trip failed for string: expected %s, got %s", strVal, roundTripStr)
	}

	// Test bool round-trip
	boolVal := false
	roundTripBool := FromPtr(ToPtr(boolVal))
	if roundTripBool != boolVal {
		t.Errorf("Round-trip failed for bool: expected %t, got %t", boolVal, roundTripBool)
	}
}