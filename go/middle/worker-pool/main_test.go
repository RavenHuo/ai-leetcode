package main

import (
	"reflect"
	"testing"
)

func TestProcessTasks(t *testing.T) {
	tests := []struct {
		name        string
		tasks       []int
		workerCount int
		want        []int
	}{
		{name: "多个任务多个 worker", tasks: []int{1, 2, 3, 4}, workerCount: 2, want: []int{1, 4, 9, 16}},
		{name: "负数和零", tasks: []int{-2, 0, 5}, workerCount: 3, want: []int{4, 0, 25}},
		{name: "空任务", tasks: []int{}, workerCount: 2, want: []int{}},
		{name: "worker 多于任务", tasks: []int{6, 7}, workerCount: 5, want: []int{36, 49}},
		{name: "非法 worker 数量", tasks: []int{3, 4}, workerCount: 0, want: []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processTasks(tt.tasks, tt.workerCount); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("processTasks(%v, %d) = %v, want %v", tt.tasks, tt.workerCount, got, tt.want)
			}
		})
	}
}
