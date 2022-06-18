package util

import (
	"testing"
)

type testversionn struct {
	version1 string
	version2 string
	expected int
}

func TestCompareVersions(t *testing.T) {
	tests := []testversionn{
		{version1: "1.2", version2: "1.1", expected: -1},
		{version1: "1.1", version2: "1.2", expected: 1},
		{version1: "1.1", version2: "1.1", expected: 0},
		{version1: "1.1.1.11", version2: "1.1.1.11", expected: 0},
		{version1: "1.1.1.10", version2: "1.1.1.11", expected: 1},
		{version1: "1.1.0", version2: "1.1.1.11", expected: 1},
	}

	for _, tc := range tests {
		actual := CompareVersions(tc.version1, tc.version2)
		if actual != tc.expected {
			t.Errorf("expected : %v ,actual : %v", tc.expected, actual)
		}
	}
}
