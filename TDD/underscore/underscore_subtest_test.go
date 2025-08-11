package main

import "testing"

func TestCame(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Simple Test", args{"thisIsACamelCase"}, "this_is_a_camel_case"},
		{"Ends with A", args{"thisendswithA"}, "thisendswith_a"},
		{"Test with Space", args{"this is with space"}, "this is with space"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Errorf("Camel() = %v, want = %v", got, tt.want)
			}
		})
	}
}
