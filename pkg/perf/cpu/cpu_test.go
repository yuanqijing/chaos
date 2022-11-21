package cpu

import (
	"github.com/yuanqijing/chaos/pkg/perf/cpu/apis"
	"testing"
)

func TestGetCores(t *testing.T) {
	tests := []struct {
		name    string
		want    []int
		wantErr bool
	}{
		{
			name:    "test",
			want:    []int{0, 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCores()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCores() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("GetCores() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmdUtilOnCore(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			want: func() []byte {
				t := apis.CpuUtilization{
					Core: 0,
				}
				return []byte(t.String())
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cmd{}
			_ = c.UtilizationPerCore(0)
		})
	}
}

func TestCmdUtilSystemWide(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			want: func() []byte {
				t := apis.CpuUtilization{
					Core: -1,
				}
				return []byte(t.String())
			}(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cmd{}
			_ = c.UtilizationSystemWide()
		})
	}
}
