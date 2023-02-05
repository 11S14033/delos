package delos

import "testing"

func Test_question2(t *testing.T) {
	type req struct {
		totalStudent, totalCandies, firstStudent int64
	}

	tests := []struct {
		name string
		req  req
		want int64
	}{
		{
			name: "case 1",
			req: req{
				totalStudent: 3,
				totalCandies: 5,
				firstStudent: 2,
			},
			want: 3,
		},
		{
			name: "case 2",
			req: req{
				totalStudent: 5,
				totalCandies: 3,
				firstStudent: 1,
			},
			want: 3,
		},
		{
			name: "case 3",
			req: req{
				totalStudent: 352926151,
				totalCandies: 380324688,
				firstStudent: 94730870,
			},
			want: 122129406,
		},
		{
			name: "case 4",
			req: req{
				totalStudent: 5,
				totalCandies: 5,
				firstStudent: 1,
			},
			want: 5,
		},
		{
			name: "case 5",
			req: req{
				totalStudent: 5,
				totalCandies: 5,
				firstStudent: 1,
			},
			want: 5,
		},
		{
			name: "case 6",
			req: req{
				totalStudent: 5,
				totalCandies: 1,
				firstStudent: 3,
			},
			want: 3,
		},
		{
			name: "case 7",
			req: req{
				totalStudent: 2,
				totalCandies: 10,
				firstStudent: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := question2(tt.req.totalStudent, tt.req.totalCandies, tt.req.firstStudent)
			if got != tt.want {
				t.Errorf("Test_question2 failed, got = %v, want %v", got, tt.want)
			}
		})
	}
}
