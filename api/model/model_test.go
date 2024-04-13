package model

import "testing"

func TestBoard_CheckWinner(t *testing.T) {
	tests := []struct {
		name string
		b    *Board
		want string
	}{
		// TODO: Add test cases.
		{
			b: &Board{
				[3]string{"a", "", ""},
				[3]string{"a", "", ""},
				[3]string{"a", "", ""},
			},
			want: "a",
		},
		{
			b: &Board{
				[3]string{"", "a", ""},
				[3]string{"", "a", ""},
				[3]string{"", "a", ""},
			},
			want: "a",
		},
		{
			b: &Board{
				[3]string{"", "", "a"},
				[3]string{"", "", "a"},
				[3]string{"", "", "a"},
			},
			want: "a",
		},
		{
			b: &Board{
				[3]string{"", "", "a"},
				[3]string{"", "a", ""},
				[3]string{"a", "", ""},
			},
			want: "a",
		},
		{
			b: &Board{
				[3]string{"a", "", ""},
				[3]string{"", "a", ""},
				[3]string{"", "", "a"},
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.CheckWinner(); got != tt.want {
				t.Errorf("Board.CheckWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
