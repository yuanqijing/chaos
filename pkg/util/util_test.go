package util

import "testing"

func TestExecCmd(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{
			name:    "test",
			want:    []byte("test"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecCmd("echo test")
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecCmd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("ExecCmd() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJson2Map(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "test",
			want:    map[string]interface{}{"test": "test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJson2Map([]byte(`{"test":"test"}`))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJson2Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ParseJson2Map() got = %v, want %v", got, tt.want)
			}
		})
	}
}
