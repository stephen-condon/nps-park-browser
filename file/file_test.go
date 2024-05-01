package file

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	type test struct {
		createFile    bool
		fileName      string
		expectedData  string
		expectedError error
	}
	tests := []test{
		{true, "", "abc", nil},
		{true, "alt_api.txt", "cde", nil},
		{false, "not_found", "", errors.New("open not_found: no such file or directory")},
	}
	for _, tc := range tests {
		// reset
		setFileName("")

		workingFilename := tc.fileName

		if len(tc.fileName) > 0 {
			setFileName(tc.fileName)
		} else {
			workingFilename = "./api.txt"
		}
		if tc.createFile {
			writeTestFile(workingFilename, []byte(tc.expectedData))
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

		if tc.createFile {
			deleteTestFile(workingFilename)

		}
	}

}

func writeTestFile(filename string, data []byte) {
	os.WriteFile(filename, data, fs.FileMode(0777))
}

func deleteTestFile(filename string) {
	os.Remove(filename)
}
