package httputils

import "testing"

func TestHttpUtils(t *testing.T) {
	testHeaders := Headers{
		"foo": "bar",
	}
	if len(testHeaders) != 1 {
		t.Errorf("testHeaders did not initialize correctly")
	}

	if BaseUrl != "https://developer.nps.gov/api" {
		t.Errorf("BaseUrl does not match")
	}
}
