package binding

import "testing"

func TestTrim(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID string `binding:"trim(T)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"XYZ"}})
	}
	{
		type model struct {
			ID string `binding:"name(ID2) trim(T) trim_sp(0)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"XYZ"}})
	}
	for _, tc := range testCases {
		test(t, tc, map[string]interface{}{"ID": " XYZ       ", "ID2": "0XYZ00"})
	}
}
