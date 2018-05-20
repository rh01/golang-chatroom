package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want int
		}{name: string("add"), args: struct {
			a int
			b int
		}{a: 1, b: 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {
	c := sub(2, 3)
	if c != -1 {
		t.Logf("sub(2,3) error expect %d, actual %d", -1, c)
	}
	t.Log("test success...")
}
