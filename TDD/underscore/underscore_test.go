package main

import "testing"

// Goal - thisIsCamelCase to this_is_camel_case

// func TestCamel(t *testing.T) {
// 	arg := "thisIsACamelCaseString"
// 	want := "this_is_a_camel_case_string"
// 	got := Camel(arg)
// 	if got != want {
// 		t.Errorf("Camel(%q) = %q; want %q", arg, got, want)
// 	}
// }

func TestCamel(t *testing.T) {
	testCases := []struct {
		arg  string
		want string
	}{
		{"thisIsACamelCaseString", "this_is_a_camel_case_string"},
		{"thisendswithA", "thisendswith_a"},
		{"this is with space", "this is with space"},
	}
	for _, tc := range testCases {
		t.Logf("Testing: %q", tc.arg)
		got := Camel(tc.arg)
		if got != tc.want {
			t.Errorf("Camel(%q) = %q; want %q", tc.arg, got, tc.want)
		}
	}
}
