package Calculate

import (
	"testing"
)

func TestCredit(t *testing.T) {
	testcase := []struct {
		name    string
		balance int
		value   int
		result  int
	}{
		{
			name:    "test1",
			balance: 900,
			value:   100,
			result:  800,
		},
		{
			name:    "test2",
			balance: 600,
			value:   50,
			result:  550,
		},
		{
			name:    "test3",
			balance: 200,
			value:   300,
			result:  -100,
		},
	}
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			expect := Credit(tc.balance, tc.value)
			if expect != tc.result {
				t.Errorf("Expected %v instead got %v", tc.result, expect)
			}
		})
	}
}
