package binding

import "testing"

func TestPrefix(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID string `binding:"prefix(PREFIX-)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"PREFIX-XYZ"}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID": "XYZ"})
	}
}
