package binding

import "testing"

func TestBinding(t *testing.T) {
	testItemCases := make([]*testItemCase, 0)
	// test for name
	{
		type model struct {
			ID int `binding:"name(ID2)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{}, map[string]interface{}{"ID2": 0}})
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{100}, map[string]interface{}{"ID2": "100"}})
	}
	// test for default
	{
		type model struct {
			ID int `binding:"default(100)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{100}, nil})
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{10}, map[string]interface{}{"ID": "10"}})
	}
	// test for split
	{
		type model struct {
			ID []string `binding:"default(A,B,C) split(T)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{[]string{"A", "B", "C"}}, nil})
	}
	// test for join
	{
		type model struct {
			ID string `binding:"join(T)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{"A,B,C"}, map[string]interface{}{"ID": []string{"A", "B", "C"}}})
	}
	// test for prefix
	{
		type model struct {
			ID string `binding:"prefix(PREFIX-)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{"PREFIX-ABC"}, map[string]interface{}{"ID": "ABC"}})
	}
	// test for suffix
	{
		type model struct {
			ID string `binding:"suffix(-SUFFIX)"`
		}
		testItemCases = append(testItemCases, &testItemCase{&model{}, &model{"ABC-SUFFIX"}, map[string]interface{}{"ID": "ABC"}})
	}
	for _, tic := range testItemCases {
		testItem(t, tic)
	}
}
