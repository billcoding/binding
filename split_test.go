package binding

import (
	"testing"
)

func TestSplit(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID  []string `binding:"split(T)"`
			ID2 []string `binding:"split(T) split_sp(0)"`
			ID3 []string `binding:"split(T) split_sp(1)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{[]string{"A", "B", "C"}, []string{"A", "B", "C"}, []string{"A", "B", "C"}}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID": "A,B,C", "ID2": "A0B0C", "ID3": "A1B1C"})
	}
}
