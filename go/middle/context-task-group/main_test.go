package main

import (
	"context"
	"errors"
	"testing"
)

func TestRunTaskGroup(t *testing.T) {
	tests := []struct {
		name  string
		tasks []func(context.Context) error
		want  error
	}{
		{
			name: "全部成功",
			tasks: []func(context.Context) error{
				func(context.Context) error { return nil },
				func(context.Context) error { return nil },
			},
			want: nil,
		},
		{
			name: "任务失败",
			tasks: []func(context.Context) error{
				func(context.Context) error { return errTaskFailed },
			},
			want: errTaskFailed,
		},
		{
			name: "多个任务中有一个失败",
			tasks: []func(context.Context) error{
				func(context.Context) error { return nil },
				func(context.Context) error { return errTaskFailed },
				func(context.Context) error { return nil },
			},
			want: errTaskFailed,
		},
		{
			name:  "空任务列表",
			tasks: nil,
			want:  nil,
		},
		{
			name: "父 context 已取消",
			tasks: []func(context.Context) error{
				func(ctx context.Context) error {
					<-ctx.Done()
					return ctx.Err()
				},
			},
			want: context.Canceled,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if tt.name == "父 context 已取消" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				cancel()
			}

			if got := runTaskGroup(ctx, tt.tasks); !errors.Is(got, tt.want) {
				t.Fatalf("runTaskGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
