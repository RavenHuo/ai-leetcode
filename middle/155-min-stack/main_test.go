package main

import "testing"

func TestMinStack(t *testing.T) {
	tests := []struct {
		name  string
		want  []int
		setup func(*MinStack)
		check func(*MinStack) []int
	}{
		{
			name: "负数最小值",
			want: []int{-3, 0, -2},
			setup: func(s *MinStack) {
				s.Push(-2)
				s.Push(0)
				s.Push(-3)
			},
			check: func(s *MinStack) []int {
				got := []int{s.GetMin()}
				s.Pop()
				got = append(got, s.Top(), s.GetMin())
				return got
			},
		},
		{
			name: "重复最小值",
			want: []int{1, 1, 2},
			setup: func(s *MinStack) {
				s.Push(2)
				s.Push(1)
				s.Push(1)
			},
			check: func(s *MinStack) []int {
				got := []int{s.GetMin()}
				s.Pop()
				got = append(got, s.GetMin())
				s.Pop()
				got = append(got, s.GetMin())
				return got
			},
		},
		{
			name: "单个元素",
			want: []int{5, 5},
			setup: func(s *MinStack) {
				s.Push(5)
			},
			check: func(s *MinStack) []int { return []int{s.Top(), s.GetMin()} },
		},
		{
			name: "递减序列",
			want: []int{-1, 0},
			setup: func(s *MinStack) {
				s.Push(3)
				s.Push(2)
				s.Push(1)
			},
			check: func(s *MinStack) []int {
				got := []int{s.GetMin()}
				s.Pop()
				return append(got, s.GetMin())
			},
		},
		{
			name: "正负数混合",
			want: []int{-10, -10},
			setup: func(s *MinStack) {
				s.Push(10)
				s.Push(-10)
				s.Push(20)
			},
			check: func(s *MinStack) []int {
				got := []int{s.GetMin()}
				s.Pop()
				s.Pop()
				return append(got, s.GetMin())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Constructor()
			tt.setup(&s)
			if got := tt.check(&s); len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			} else {
				for i := range got {
					if got[i] != tt.want[i] {
						t.Fatalf("got %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
