package binding

import (
	"testing"
)

func TestJoin(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID  string `binding:"join(T)"`
			ID2 string `binding:"name(ID) join(T) join_sp(0)"`
			ID3 string `binding:"name(ID) join(T) join_sp(1)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"A,B,C", "A0B0C", "A1B1C"}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID": []string{"A", "B", "C"}})
	}
}
