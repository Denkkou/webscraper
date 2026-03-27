package main

import "testing"

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct { 
		name		string
		inputHTML	string
		expected	string
	} {
		// Tests go here
		{
			name: "Get <h1>",
			inputHTML: "<html><body><h1>Title</h1></body></html>",
			expected: "Title",
		},
		{
			name: "Get <h2>",
			inputHTML: "<html><body><h2>Secondary Title</h1></body></html>",
			expected: "Secondary Title",
		},
		{
			name: "Get <h1> even if <h2> exists",
			inputHTML: "<html><body><h1>Title</h1><h2>Secondary Title</h2></body></html>",
			expected: "Title",
		},
		{
			name: "Return empty string if no heading found",
			inputHTML: "<html><body><p>Paragraph</p></body></html>",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}