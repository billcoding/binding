package binding

import "testing"

func TestSuffix(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID string `binding:"suffix(-SUFFIX)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"XYZ-SUFFIX"}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID": "XYZ"})
	}
}
