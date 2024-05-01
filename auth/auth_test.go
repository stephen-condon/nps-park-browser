package auth

import (
	"io/fs"
	"os"
	"testing"
)

func TestApiKey(t *testing.T) {
	type test struct {
		createFile    bool
		predefinedKey string
		expectedData  string
	}
	tests := []test{
		{true, "", "cba"},
		{false, "", ""},
		{false, "xyz", "xyz"},
	}
	for _, tc := range tests {
		apiKey = ""
		if tc.createFile {
			writeTestFile("./api.txt", []byte(tc.expectedData))
		}
		if len(tc.predefinedKey) > 0 {
			apiKey = tc.predefinedKey
		}
		result := ApiKey()

		if result != tc.expectedData {
			t.Errorf(`Result: %v; Expected Result: %v`, result, tc.expectedData)
		}

		if tc.createFile {
			deleteTestFile("./api.txt")
		}
	}

}

func writeTestFile(filename string, data []byte) {
	os.WriteFile(filename, data, fs.FileMode(0777))
}

func deleteTestFile(filename string) {
	os.Remove(filename)
}
