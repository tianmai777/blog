package util

import "testing"

func TestEncodeMD5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{
			str: "luo",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeMD5(tt.args.str); got != tt.want {
				t.Errorf("EncodeMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
