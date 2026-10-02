package vo_test_test

import (
	"github.com/GoPersonalCluster/go_rabbitMq_filter/app/internal/tests/vo_test"
	"testing"
)

func TestTestNewIPAddress(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		t *testing.T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vo_test.TestNewIPAddress(tt.t)
		})
	}
}
