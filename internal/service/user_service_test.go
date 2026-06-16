package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	// Set a fixed "current time" for the test
	mockNow := time.Date(2024, 5, 20, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "Birthday has already passed this year",
			dob:      time.Date(1990, 1, 15, 0, 0, 0, 0, time.UTC),
			expected: 34,
		},
		{
			name:     "Birthday is exactly today",
			dob:      time.Date(1990, 5, 20, 0, 0, 0, 0, time.UTC),
			expected: 34,
		},
		{
			name:     "Birthday has not happened yet this year",
			dob:      time.Date(1990, 11, 25, 0, 0, 0, 0, time.UTC),
			expected: 33, // Still 33 until Nov 25th, 2024
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := calculateAge(tc.dob, mockNow)
			if result != tc.expected {
				t.Errorf("Expected age %d, but got %d", tc.expected, result)
			}
		})
	}
}
