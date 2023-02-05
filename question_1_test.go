package delos

import "testing"

func Test_question1(t *testing.T) {
	type req struct {
		expected schedule
		returned schedule
	}

	tests := []struct {
		name string
		req  req
		want int
	}{
		{
			name: "case 1",
			req: req{
				expected: schedule{15, 8, 2022},
				returned: schedule{22, 8, 2022},
			},
			want: 105,
		},
		{
			name: "case 2",
			req: req{
				expected: schedule{7, 6, 2022},
				returned: schedule{19, 8, 2022},
			},
			want: 1000,
		},
		{
			name: "case 3",
			req: req{
				expected: schedule{15, 8, 2022},
				returned: schedule{22, 9, 2024},
			},
			want: 24000,
		},
		{
			name: "case 4",
			req: req{
				expected: schedule{15, 8, 2022},
				returned: schedule{22, 8, 2020},
			},
			want: 0,
		},
		{
			name: "case 5",
			req: req{
				expected: schedule{15, 8, 2022},
				returned: schedule{22, 7, 2022},
			},
			want: 0,
		},
		{
			name: "case 6",
			req: req{
				expected: schedule{15, 8, 2022},
				returned: schedule{14, 8, 2022},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := question1(tt.req.expected, tt.req.returned)
			if got != tt.want {
				t.Errorf("Test_question1 failed, got = %v, want %v", got, tt.want)
			}
		})
	}
}
