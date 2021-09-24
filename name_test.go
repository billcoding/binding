package binding

import "testing"

func TestName(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID string `binding:"name(ID2)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"XYZ"}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID2": "XYZ"})
	}
}
