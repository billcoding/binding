package binding

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type testCase struct {
	original interface{}
	current  interface{}
}

type testItemCase struct {
	original interface{}
	current  interface{}
	m        map[string]interface{}
}

func test(t *testing.T, tc *testCase, m map[string]interface{}) {
	New(tc.original).BindMap(m)
	require.Equal(t, tc.original, tc.current)
}

func testItem(t *testing.T, tc *testItemCase) {
	New(tc.original).BindMap(tc.m)
	require.Equal(t, tc.original, tc.current)
}
