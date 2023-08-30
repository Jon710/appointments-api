package mypackage

import "testing"

func TestPrintHello(t *testing.T) {
	// Test case 1: valid input
	expected := "Hello, John! This is mypackage speaking!"
	result, err := PrintHello("John")
	if err != nil {
		t.Errorf("PrintHello('John') returned an error: %v", err)
	}
	if result != expected {
		t.Errorf("PrintHello('John') = %v, expected %v", result, expected)
	}
}

func TestPrintHelloEmpty(t *testing.T) {
	// Test case 2: empty input
	expected := ""
	result, err := PrintHello("")
	if err == nil {
		t.Errorf("PrintHello('') did not return an error")
	}
	if result != expected {
		t.Errorf("PrintHello('') = %v, expected %v", result, expected)
	}
}
