package code

import "testing"

// We try to testing 3 kind of case on Palindrome code
func palindromeTest(t *testing.T) {
	test1 := Utama("1 10") // First test with parameter 1 to 10
	if test1 != 9 {
		t.Error(("Result should be 9"))
	}
	test2 := Utama("11 20") // Second test with parameter 11 to 20
	if test2 != 1 {
		t.Error(("Result should be 1"))
	}
	test3 := Utama("21 30") // Third test with parameter 21 to 30
	if test3 != 1 {
		t.Error(("Result should be 1"))
	}
}
