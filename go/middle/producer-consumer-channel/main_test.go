package main

import (
	"reflect"
	"testing"
)

func TestProducerConsumer(t *testing.T) {
	tests := []struct {
		name       string
		items      []int
		bufferSize int
		want       []int
	}{
		{name: "buffered channel", items: []int{1, 2, 3, 4, 5}, bufferSize: 2, want: []int{1, 2, 3, 4, 5}},
		{name: "empty input", items: []int{}, bufferSize: 1, want: []int{}},
		{name: "unbuffered channel", items: []int{0, -1, 0, 2}, bufferSize: 0, want: []int{0, -1, 0, 2}},
		{name: "single item", items: []int{42}, bufferSize: 0, want: []int{42}},
		{name: "large buffer", items: []int{3, 1, 4, 1, 5, 9}, bufferSize: 10, want: []int{3, 1, 4, 1, 5, 9}},
		{name: "repeated zero", items: []int{0, 0, 0}, bufferSize: 1, want: []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := producerConsumer(tt.items, tt.bufferSize)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("producerConsumer() = %v, want %v", got, tt.want)
			}
		})
	}
}
