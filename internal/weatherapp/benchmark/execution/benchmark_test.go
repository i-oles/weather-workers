package execution

import (
	"testing"
	"time"
)

func Test_calculateStandardDeviation(t *testing.T) {
	type args struct {
		durations []time.Duration
	}

	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "standard deviation 6 elements",
			args: args{
				durations: []time.Duration{
					32 * time.Second,
					34 * time.Second,
					32 * time.Second,
					34 * time.Second,
					32 * time.Second,
					34 * time.Second,
				},
			},
			want: 1 * time.Second,
		},
		{
			name: "standard deviation 2 elements",
			args: args{
				durations: []time.Duration{
					34 * time.Second,
					34 * time.Second,
				},
			},
			want: 0 * time.Second,
		},
		{
			name: "standard deviation 0 elements",
			args: args{
				durations: []time.Duration{},
			},
			want: 0 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateStandardDeviation(tt.args.durations); got != tt.want {
				t.Errorf("calculateStandardDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateAverageExecutionDuration(t *testing.T) {
	type args struct {
		durations []time.Duration
	}

	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "average 4 elements",
			args: args{
				durations: []time.Duration{
					3 * time.Second,
					10 * time.Second,
					13 * time.Second,
					23 * time.Second,
				},
			},
			want: time.Duration(12.25 * 1e9),
		},
		{
			name: "average 3 elements",
			args: args{
				durations: []time.Duration{
					10 * time.Second,
					20 * time.Second,
					30 * time.Second,
				},
			},
			want: 20 * time.Second,
		},
		{
			name: "average 0 elements",
			args: args{
				durations: []time.Duration{},
			},
			want: 0 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateAverageExecutionDuration(tt.args.durations); got != tt.want {
				t.Errorf("calculateAverageExecutionDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
