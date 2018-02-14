// +build unit

package validator

import "testing"

type testValidA58UnitInput struct {
	address []byte
	toPass  bool
	toErr   bool
}

func TestValidA58Unit(t *testing.T) {
	tests := []testValidA58UnitInput{
		testValidA58UnitInput{
			address: []byte("1NrxWfPM3w8GVXBkY2LpGy983wRxX45LMc"),
			toPass:  true,
			toErr:   false,
		},
		testValidA58UnitInput{
			address: []byte("0xE272B9E2e25eF2DDde34aEEF6aF1736e88c1D2c9"),
			toPass:  false,
			toErr:   true,
		},
		testValidA58UnitInput{
			address: []byte("5NrxWfPM3w8GVXBkY2LpGy983wRxX45LMc"),
			toPass:  false,
			toErr:   true,
		},
		testValidA58UnitInput{
			address: nil,
			toPass:  false,
			toErr:   false,
		},
	}

	for idx, tt := range tests {
		ok, err := ValidA58(tt.address)

		if ok != tt.toPass || tt.toErr != (err != nil) {
			t.Errorf("test #%v false; expected: toPass: %v, toErr: %v; received: ok: %v; err: %v", idx+1, tt.toPass, tt.toErr, ok, err)
		}
	}
}
