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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("\nExpected:\t %v \nActual:\t %v", tc.expected, actual)
			}
			if err != nil {
				t.Errorf("Test - FAIL: Error %v", err)
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct { 
		name		string
		inputHTML	string
		expected	string
	} {
		// Tests go here
		{
			name: "Get <p>",
			inputHTML: "<html><body><p>Paragraph</p></body></html>",
			expected: "Paragraph",
		},
		{
			name: "Get <p> from within <main> when a prior <p> exists",
			inputHTML: "<html><body><p>Paragraph1</p><main><p>Paragraph2</p></main></body></html>",
			expected: "Paragraph2",
		},
		{
			name: "Default to first <p> if <main> is empty",
			inputHTML: "<html><body><p>Paragraph</p><main></main></body></html>",
			expected: "Paragraph",
		},
		{
			name: "Return empty string if no <p> found at all",
			inputHTML: "<html><body><main></body></html>",
			expected: "",
		},
		{
			name: "Get first <p> from within <main> where many exist",
			inputHTML: "<html><body><p>Paragraph1</p><main><p>Paragraph2</p><p>Paragraph3</p></body></html>",
			expected: "Paragraph2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getFirstParagraphFromHTML(tc.inputHTML)
			if actual != tc.expected {
				t.Errorf("\nExpected:\t %v \nActual:\t %v", tc.expected, actual)
			}
			if err != nil {
				t.Errorf("Test - FAIL: Error %v", err)
			}
		})
	}
}