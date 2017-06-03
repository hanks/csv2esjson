package models

import "testing"

func TestCSV_Load(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		c       *CSV
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Load(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("CSV.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
