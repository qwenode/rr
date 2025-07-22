package rr

import (
	"testing"
	"time"
)

func TestTimeIsZero(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected bool
	}{
		{
			name:     "零值时间",
			input:    time.Time{},
			expected: true,
		},
		{
			name:     "Unix时间戳0",
			input:    time.Unix(0, 0),
			expected: true,
		},
		{
			name:     "当前时间",
			input:    time.Now(),
			expected: false,
		},
		{
			name:     "特定时间",
			input:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "Unix时间戳1",
			input:    time.Unix(1, 0),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeIsZero(tt.input)
			if result != tt.expected {
				t.Errorf("TimeIsZero(%v) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTimeIsZeroPtr(t *testing.T) {
	zeroTime := time.Time{}
	unixZeroTime := time.Unix(0, 0)
	currentTime := time.Now()
	specificTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    *time.Time
		expected bool
	}{
		{
			name:     "nil指针",
			input:    nil,
			expected: true,
		},
		{
			name:     "零值时间指针",
			input:    &zeroTime,
			expected: true,
		},
		{
			name:     "Unix时间戳0指针",
			input:    &unixZeroTime,
			expected: true,
		},
		{
			name:     "当前时间指针",
			input:    &currentTime,
			expected: false,
		},
		{
			name:     "特定时间指针",
			input:    &specificTime,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TimeIsZeroPtr(tt.input)
			if result != tt.expected {
				t.Errorf("TimeIsZeroPtr(%v) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}