package go_advance

import (
	"testing"
)

func TestAlterPrint(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
			ret := AlterPrint()
			if ret != s {
				t.Errorf("AlterPrint() = %v, want %v", ret, s)
			}
		})
	}
}
