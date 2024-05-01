package file

import (
	"errors"
	"testing"
)

func TestRead(t *testing.T) {
	type test struct {
		fileName      string
		expectedData  string
		expectedError error
	}
	tests := []test{
		{"", "abc", nil},
		{"alt_api.txt", "cde", nil},
		{"not_found", "", errors.New("open not_found: no such file or directory")},
	}
	for _, tc := range tests {
		// reset
		setFileName("")

		if len(tc.fileName) > 0 {
			setFileName(tc.fileName)
		}

		result, err := Read()

		if result != tc.expectedData {
			t.Errorf(`Result: %v; Expected Result: %v`, result, tc.expectedData)
		}

		if tc.expectedError == nil {
			if err != nil {
				t.Errorf(`Received Error: %v; Expected nil`, err.Error())
			}
		} else {
			if err == nil {
				t.Errorf(`Received no error; Expected Error: %v`, tc.expectedError.Error())
			} else if err.Error() != tc.expectedError.Error() {
				t.Errorf(`Received Error: %v; Expected Error: %v`, err.Error(), tc.expectedError.Error())
			}
		}
	}

}
