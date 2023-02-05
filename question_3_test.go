package delos

import "testing"

func Test_question3(t *testing.T) {
	tests := []struct {
		name string
		req  []int32
		want string
	}{
		{
			name: "case 1",
			req:  []int32{1, 3, 5, 4},
			want: "YES",
		},
		{
			name: "case 2",
			req:  []int32{1, 5, 7, 2, 4},
			want: "YES",
		},
		{
			name: "case 3",
			req:  []int32{1, 3, 4, 9},
			want: "NO",
		},
		{
			name: "case 4",
			req:  []int32{1, 1, 4, 1, 1},
			want: "YES",
		},
		{
			name: "case 5",
			req:  []int32{2, 0, 0, 0},
			want: "YES",
		},
		{
			name: "case 6",
			req:  []int32{0, 0, 2, 0},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := question3(tt.req)
			if got != tt.want {
				t.Errorf("Test_question3 failed, got = %v, want %v", got, tt.want)
			}
		})
	}
}
