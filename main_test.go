package main

import "testing"

func TestFix(t *testing.T) {
	if Fix("abcddd") != "abcd" {
		t.Error("expected abcd")
	}
}

func TestTableFix(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"aaaabc", "zabc"},
		{"wwwbb", "wb"},
		{"zzzqwed", "zqwed"},
	}

	for _, test := range tests {
		if output := Fix(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
			t.Error("Test Failed:: ", "input", test.input, " - Expected: ", test.expected, " -Received:  ", output)
		}
	}

}
