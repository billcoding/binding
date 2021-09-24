package binding

import (
	"testing"
)

func TestDefault(t *testing.T) {
	testCases := make([]*testCase, 0)
	{
		type model struct {
			ID   int   `binding:"default(100)"`
			ID8  int8  `binding:"default(100)"`
			ID16 int16 `binding:"default(100)"`
			ID32 int32 `binding:"default(100)"`
			ID64 int64 `binding:"default(100)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{100, 100, 100, 100, 100}})
	}
	{
		type model struct {
			ID   uint   `binding:"default(100)"`
			ID8  uint8  `binding:"default(100)"`
			ID16 uint16 `binding:"default(100)"`
			ID32 uint32 `binding:"default(100)"`
			ID64 uint64 `binding:"default(100)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{100, 100, 100, 100, 100}})
	}
	{
		type model struct {
			ID   float32 `binding:"default(100)"`
			ID64 float64 `binding:"default(100)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{100, 100}})
	}
	{
		type model struct {
			ID string `binding:"default(100)"`
		}
		testCases = append(testCases, &testCase{&model{}, &model{"100"}})
	}
	for _, tc := range testCases {
		test(t, tc, nil)
	}
}
